package matrixgates

import (
	"errors"

	"github.com/matrixbotio/exchange-gates-lib/workers"
)

//ExchangeInterface - universal exchange adapter interface
type ExchangeInterface interface {
	// Adapter
	GetName() string
	GetTag() string
	GetID() int

	// Methods
	Connect(credentials APICredentials) error
	GetAccountData() (*AccountData, error) // account data with balances
	VerifyAPIKeys(keyPublic, keySecret string) error

	// Order
	GetOrderData(pairSymbol string, orderID int64) (*OrderData, error)
	PlaceOrder(order BotOrder, pairLimits ExchangePairData) (*CreateOrderResponse, *BotOrderAdjusted, error)

	// Pair
	GetPairData(pairSymbol string) (*ExchangePairData, error)
	GetPairLastPrice(pairSymbol string) (float64, error)
	CancelPairOrder(pairSymbol string, orderID int64) error
	CancelPairOrders(pairSymbol string) error
	GetPairOpenOrders(pairSymbol string) ([]*Order, error)
	GetPairs() ([]*ExchangePairData, error)
	GetPairBalance(pair PairSymbolData) (*PairBalance, error)

	// Workers
	GetPriceWorker() workers.IPriceWorker
	GetCandleWorker() workers.ICandleWorker
	GetTradeEventsWorker() workers.ITradeEventWorker
}

const (
	exchangeIDbinanceSpot = 1
)

// GetExchangeAdapter - get supported exchange adapter with interface
func GetExchangeAdapter(exchangeID int) (ExchangeInterface, error) {
	switch exchangeID {
	default:
		return nil, errors.New("exchange not found")
	case exchangeIDbinanceSpot:
		return NewBinanceSpotAdapter(exchangeIDbinanceSpot), nil
	}
}
