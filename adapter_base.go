package matrixgates

import sharederrs "github.com/matrixbotio/shared-errors"

//ExchangeAdapter - abstract universal exchange adapter
type ExchangeAdapter struct {
	ExchangeID int
	Name       string
}

func newExchangeAdapter(name string, exchangeID int) *ExchangeAdapter {
	return &ExchangeAdapter{
		ExchangeID: exchangeID,
		Name:       name,
	}
}

//Connect to exchange
func (a *ExchangeAdapter) Connect(credentials APICredentials) *sharederrs.APIError {
	return nil
}

//PlaceOrder ..
func (a *ExchangeAdapter) PlaceOrder(order BotOrder) (*CreateOrderResponse, *sharederrs.APIError) {
	return nil, nil
}

//GetAccountData ..
func (a *ExchangeAdapter) GetAccountData() (*AccountData, *sharederrs.APIError) {
	return nil, nil
}

//GetPairLastPrice ..
func (a *ExchangeAdapter) GetPairLastPrice(pairSymbol string) (float64, *sharederrs.APIError) {
	return 0, nil
}

//CancelPairOrder ..
func (a *ExchangeAdapter) CancelPairOrder() *sharederrs.APIError {
	return nil
}

//CancelPairOrders ..
func (a *ExchangeAdapter) CancelPairOrders() *sharederrs.APIError {
	return nil
}

//GetOrderData ..
func (a *ExchangeAdapter) GetOrderData(pairSymbol string, orderID int64) (*TradeEventData, *sharederrs.APIError) {
	return nil, nil
}

//GetPairOpenOrders ..
func (a *ExchangeAdapter) GetPairOpenOrders() ([]*Order, *sharederrs.APIError) {
	//TODO
	return nil, nil
}

//GetPairs get all Binance pairs
func (a *ExchangeAdapter) GetPairs() ([]*ExchangePairData, *sharederrs.APIError) {
	return nil, nil
}

//VerifyAPIKeys ..
func (a *ExchangeAdapter) VerifyAPIKeys(keyPublic, keySecret string) *sharederrs.APIError {
	return nil
}
