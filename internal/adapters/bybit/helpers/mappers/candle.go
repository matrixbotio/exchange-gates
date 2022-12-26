package mappers

import (
	"errors"
	"fmt"
	"math"
	"strconv"
	"time"

	"github.com/hirokisan/bybit/v2"
	"github.com/matrixbotio/exchange-gates-lib/internal/consts"
	"github.com/matrixbotio/exchange-gates-lib/internal/workers"
)

type CandleData struct {
	Open   float64
	Close  float64
	High   float64
	Low    float64
	Volume float64
}

type CandleTimeData struct {
	StartTimeMs int64
	EndTimeMs   int64
}

func ParseCandle(open, close, high, low, volume string) (CandleData, error) {
	if open == "" {
		return CandleData{}, errors.New("'open' price is empty")
	}
	openPrice, err := strconv.ParseFloat(open, 64)
	if err != nil {
		return CandleData{}, fmt.Errorf("parse 'open' price: %w", err)
	}

	if close == "" {
		return CandleData{}, errors.New("'close' price is empty")
	}
	closePrice, err := strconv.ParseFloat(close, 64)
	if err != nil {
		return CandleData{}, fmt.Errorf("parse 'close' price: %w", err)
	}

	if high == "" {
		return CandleData{}, errors.New("'high' price is empty")
	}
	highPrice, err := strconv.ParseFloat(high, 64)
	if err != nil {
		return CandleData{}, fmt.Errorf("parse 'high' price: %w", err)
	}

	if low == "" {
		return CandleData{}, errors.New("'low' price is empty")
	}
	lowPrice, err := strconv.ParseFloat(low, 64)
	if err != nil {
		return CandleData{}, fmt.Errorf("parse 'low' price: %w", err)
	}

	if volume == "" {
		return CandleData{}, errors.New("'volume' is empty")
	}
	volumeParsed, err := strconv.ParseFloat(volume, 64)
	if err != nil {
		return CandleData{}, fmt.Errorf("parse 'volume': %w", err)
	}

	return CandleData{
		Open:   openPrice,
		Close:  closePrice,
		High:   highPrice,
		Low:    lowPrice,
		Volume: volumeParsed,
	}, nil
}

func ConvertWsCandle(
	pairSymbol string,
	eventData bybit.V5WebsocketPublicKlineData,
) (workers.CandleEvent, error) {
	data, err := ParseCandle(
		eventData.Open,
		eventData.Close,
		eventData.High,
		eventData.Low,
		eventData.Volume,
	)
	if err != nil {
		return workers.CandleEvent{}, fmt.Errorf("parse candle: %w", err)
	}

	return workers.CandleEvent{
		Symbol: pairSymbol,
		Candle: workers.CandleData{
			StartTime: int64(eventData.Start),
			EndTime:   int64(eventData.End),
			Interval:  consts.CandlesInterval,
			Open:      data.Open,
			Close:     data.Close,
			High:      data.High,
			Low:       data.Low,
			Volume:    data.Volume,
		},
		Time:       int64(eventData.Timestamp),
		IsFinished: eventData.Confirm,
	}, nil
}

func ConvertHistoricalCandle(
	pairSymbol string,
	eventData bybit.V5GetKlineItem,
	intervalDuration time.Duration,
	intervalCode string,
) (workers.CandleData, error) {
	data, err := ParseCandle(
		eventData.Open,
		eventData.Close,
		eventData.High,
		eventData.Low,
		eventData.Volume,
	)
	if err != nil {
		return workers.CandleData{}, fmt.Errorf("parse candle: %w", err)
	}

	timeData, err := getCandleEndTimeData(eventData, intervalDuration)
	if err != nil {
		return workers.CandleData{}, fmt.Errorf("get candle time: %w", err)
	}

	return workers.CandleData{
		StartTime: timeData.StartTimeMs,
		EndTime:   timeData.EndTimeMs,
		Interval:  intervalCode,
		Open:      data.Open,
		Close:     data.Close,
		High:      data.High,
		Low:       data.Low,
		Volume:    data.Volume,
	}, nil
}

func getCandleEndTimeData(
	eventData bybit.V5GetKlineItem,
	intervalDuration time.Duration,
) (CandleTimeData, error) {
	if eventData.StartTime == "" {
		return CandleTimeData{}, errors.New("start time is empty")
	}
	startTimeMs, err := strconv.ParseInt(eventData.StartTime, 10, 64)
	if err != nil {
		return CandleTimeData{}, fmt.Errorf("parse candle start time: %w", err)
	}

	startTimeSeconds := int64(math.Floor(float64(startTimeMs) / 1000))
	endTimeNanosecondsMod := int64(math.Mod(float64(startTimeMs), 1000) * math.Pow10(6))
	endTime := time.Unix(startTimeSeconds, endTimeNanosecondsMod).Add(intervalDuration)

	return CandleTimeData{
		StartTimeMs: startTimeMs,
		EndTimeMs:   endTime.UnixMilli(),
	}, nil
}
