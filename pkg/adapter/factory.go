package adapter

import (
	"errors"

	"github.com/matrixbotio/exchange-gates-lib/internal/adapters/binance"
	"github.com/matrixbotio/exchange-gates-lib/internal/adapters/binance/wrapper"
	"github.com/matrixbotio/exchange-gates-lib/internal/adapters/bybit"
	"github.com/matrixbotio/exchange-gates-lib/internal/consts"
)

func CreateAdapter(exchangeID int) (Adapter, error) {
	switch exchangeID {
	default:
		return nil, errors.New("exchange not found")
	case consts.ExchangeIDbinanceSpot:
		return binance.New(wrapper.NewWrapper()), nil
	case consts.ExchangeIDbybitSpot:
		return bybit.New(), nil
	case consts.ExchangeIDMexcSpot:
		// TODO
		return nil, errors.New("mexc adapter not ready yet")
	case consts.ExchangeIDKucoinSpot:
		// TODO
		return nil, errors.New("kucoin adapter not ready yet")
	case consts.ExchangeIDGateSpot:
		// TODO
		return nil, errors.New("gate adapter not ready yet")
	case consts.ExhcnageIDPoloniexSpot:
		// TODO
		return nil, errors.New("poloniex adapter not ready yet")
	}
}

func CreateAdapters() map[int]Adapter {
	return map[int]Adapter{
		consts.ExchangeIDbinanceSpot: binance.New(wrapper.NewWrapper()),
		consts.ExchangeIDbybitSpot:   bybit.New(),
	}
}
