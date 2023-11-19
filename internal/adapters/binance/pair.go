package binance

import (
	"context"
	"fmt"
	"strconv"

	"github.com/matrixbotio/exchange-gates-lib/internal/adapters/binance/helpers/mappers"
	"github.com/matrixbotio/exchange-gates-lib/internal/structs"
)

// GetPairLastPrice - get pair last price ^ↀᴥↀ^
func (a *adapter) GetPairLastPrice(pairSymbol string) (float64, error) {
	tickerService := a.binanceAPI.NewListPricesService()
	prices, err := tickerService.Symbol(pairSymbol).Do(context.Background())
	if err != nil {
		return 0, fmt.Errorf("last price: %w", err)
	}

	// until just brute force. need to be done faster
	var price float64 = 0
	var parseErr error
	for _, p := range prices {
		if p.Symbol == pairSymbol {
			price, parseErr = strconv.ParseFloat(p.Price, 64)
			if parseErr != nil {
				return 0, fmt.Errorf("parse price %q: %w", p.Price, err)
			}
			break
		}
	}
	return price, nil
}

// GetPairData - get pair data & limits
func (a *adapter) GetPairData(pairSymbol string) (structs.ExchangePairData, error) {
	exchangeInfo, err := a.binanceAPI.NewExchangeInfoService().
		Symbol(pairSymbol).Do(context.Background())
	if err != nil {
		return structs.ExchangePairData{}, err
	}

	// find pairSymbol
	for _, symbolData := range exchangeInfo.Symbols {
		return getExchangePairData(symbolData, a.ExchangeID)
	}

	return structs.ExchangePairData{},
		fmt.Errorf("data for %q pair not found", pairSymbol)
}

// GetPairOpenOrders - get open orders array
func (a *adapter) GetPairOpenOrders(pairSymbol string) ([]structs.OrderData, error) {
	ordersRaw, err := a.binanceAPI.NewListOpenOrdersService().
		Symbol(pairSymbol).Do(context.Background())
	if err != nil {
		return nil, err
	}

	return mappers.ConvertOrders(ordersRaw)
}

// GetPairs get all Binance pairs
func (a *adapter) GetPairs() ([]structs.ExchangePairData, error) {
	service := a.binanceAPI.NewExchangeInfoService()
	res, err := service.Do(context.Background())
	if err != nil {
		return nil, fmt.Errorf("connect: %w", err)
	}

	var lastError error
	pairs := []structs.ExchangePairData{}
	for _, symbolData := range res.Symbols {
		pairData, err := getExchangePairData(symbolData, a.ExchangeID)
		if err != nil {
			lastError = err
		} else {
			pairs = append(pairs, pairData)
		}
	}
	return pairs, lastError
}
