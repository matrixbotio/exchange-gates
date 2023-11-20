package adapters

import (
	"context"

	"github.com/matrixbotio/exchange-gates-lib/internal/structs"
	"github.com/matrixbotio/exchange-gates-lib/internal/workers"
	pkgStructs "github.com/matrixbotio/exchange-gates-lib/pkg/structs"
)

type Adapter interface {
	// ADAPTER
	GetName() string
	GetTag() string
	GetID() int

	// BASIC. TBD: https://github.com/matrixbotio/exchange-gates-lib/issues/156
	// Connect to exchange
	Connect(credentials pkgStructs.APICredentials) error
	CanTrade() (bool, error)
	VerifyAPIKeys(keyPublic, keySecret string) error
	GetAccountBalance() ([]structs.Balance, error)

	// ORDER
	// GetOrderData - get order data
	GetOrderData(pairSymbol string, orderID int64) (structs.OrderData, error)
	// GetClientOrderData - get order data by client order ID
	GetOrderByClientOrderID(pairSymbol, clientOrderID string) (structs.OrderData, error)
	// PlaceOrder - place order on exchange
	PlaceOrder(
		ctx context.Context,
		order structs.BotOrderAdjusted,
	) (structs.CreateOrderResponse, error)
	GetOrderExecFee(
		pairSymbol string,
		orderSide string,
		orderID int64,
	) (structs.OrderFees, error)

	// PAIR
	// GetPairData - get pair data & limits
	GetPairData(pairSymbol string) (structs.ExchangePairData, error)
	// GetPairLastPrice - get pair last price ^ↀᴥↀ^
	GetPairLastPrice(pairSymbol string) (float64, error)
	// CancelPairOrder - cancel one exchange pair order by ID
	CancelPairOrder(pairSymbol string, orderID int64, ctx context.Context) error
	// CancelPairOrder - cancel one exchange pair order by client order ID
	CancelPairOrderByClientOrderID(
		pairSymbol string,
		clientOrderID string,
		ctx context.Context,
	) error
	// GetPairOpenOrders - get open orders array
	GetPairOpenOrders(pairSymbol string) ([]structs.OrderData, error)
	// GetPairs get all Binance pairs
	GetPairs() ([]structs.ExchangePairData, error)
	GetPairBalance(pair structs.PairSymbolData) (structs.PairBalance, error)

	// WORKERS
	GetPriceWorker(callback workers.PriceEventCallback) workers.IPriceWorker
	GetCandleWorker() workers.ICandleWorker
	GetTradeEventsWorker() workers.ITradeEventWorker

	// CANDLE
	GetCandles(limit int, symbol string, interval string) ([]workers.CandleData, error)
}
