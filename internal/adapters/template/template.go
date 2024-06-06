package mexc

import (
	"context"

	adp "github.com/matrixbotio/exchange-gates-lib/internal/adapters"
	"github.com/matrixbotio/exchange-gates-lib/internal/structs"
	"github.com/matrixbotio/exchange-gates-lib/internal/workers"
	pkgStructs "github.com/matrixbotio/exchange-gates-lib/pkg/structs"
)

type adapter struct {
	ExchangeID int
	Name       string
	Tag        string

	//client *lib.client
}

func New() adp.Adapter {
	return &adapter{
		ExchangeID: 0,
		Name:       "Template Spot",
		Tag:        "template-spot",
	}
}

func (a *adapter) GetTag() string {
	return a.Tag
}

func (a *adapter) GetID() int {
	return a.ExchangeID
}

func (a *adapter) GetName() string {
	return a.Name
}

func (a *adapter) Connect(credentials pkgStructs.APICredentials) error {
	// TODO: create exchange client & connect
	// TODO: sync server time
	return nil
}

func (a *adapter) CanTrade() (bool, error) {
	// TODO
	return true, nil
}

func (a *adapter) VerifyAPIKeys(keyPublic, keySecret string) error {
	// TODO
	return nil
}

func (a *adapter) GetAccountBalance() ([]structs.Balance, error) {
	// TODO
	return nil, nil
}

func (a *adapter) GetOrderData(pairSymbol string, orderID int64) (structs.OrderData, error) {
	// TODO
	return structs.OrderData{}, nil
}

func (a *adapter) GetOrderByClientOrderID(
	pairSymbol string,
	clientOrderID string,
) (structs.OrderData, error) {
	// TODO
	return structs.OrderData{}, nil
}

func (a *adapter) PlaceOrder(
	ctx context.Context,
	order structs.BotOrderAdjusted,
) (structs.CreateOrderResponse, error) {
	// TODO
	return structs.CreateOrderResponse{}, nil
}

func (a *adapter) GetOrderExecFee(
	baseAssetTicker string,
	quoteAssetTicker string,
	orderSide string,
	orderID int64,
) (structs.OrderFees, error) {
	// TODO
	return structs.OrderFees{}, nil
}

func (a *adapter) GetPairData(pairSymbol string) (structs.ExchangePairData, error) {
	// TODO
	return structs.ExchangePairData{}, nil
}

func (a *adapter) GetPairLastPrice(pairSymbol string) (float64, error) {
	// TODO
	return 0, nil
}

func (a *adapter) CancelPairOrder(
	pairSymbol string,
	orderID int64,
	ctx context.Context,
) error {
	// TODO
	return nil
}

func (a *adapter) CancelPairOrderByClientOrderID(
	pairSymbol string,
	clientOrderID string,
	ctx context.Context,
) error {
	// TODO
	return nil
}

func (a *adapter) GetPairOpenOrders(pairSymbol string) ([]structs.OrderData, error) {
	// TODO
	return nil, nil
}

func (a *adapter) GetPairs() ([]structs.ExchangePairData, error) {
	// TODO
	return nil, nil
}

func (a *adapter) GetPairBalance(pair structs.PairSymbolData) (
	structs.PairBalance,
	error,
) {
	return structs.PairBalance{}, nil
}

func (a *adapter) GetPriceWorker(callback workers.PriceEventCallback) workers.IPriceWorker {
	// TODO
	return nil
}

func (a *adapter) GetCandleWorker() workers.ICandleWorker {
	// TODO
	return nil
}

func (a *adapter) GetTradeEventsWorker() workers.ITradeEventWorker {
	// TODO
	return nil
}

func (a *adapter) GetCandles(
	limit int,
	symbol string,
	interval string,
) ([]workers.CandleData, error) {
	// TODO
	return nil, nil
}
