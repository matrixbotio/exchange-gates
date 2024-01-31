// Code generated by mockery v2.14.1. DO NOT EDIT.

package adapters

import (
	context "context"

	internalstructs "github.com/matrixbotio/exchange-gates-lib/internal/structs"
	mock "github.com/stretchr/testify/mock"

	structs "github.com/matrixbotio/exchange-gates-lib/pkg/structs"

	workers "github.com/matrixbotio/exchange-gates-lib/internal/workers"
)

// MockAdapter is an autogenerated mock type for the Adapter type
type MockAdapter struct {
	mock.Mock
}

type MockAdapter_Expecter struct {
	mock *mock.Mock
}

func (_m *MockAdapter) EXPECT() *MockAdapter_Expecter {
	return &MockAdapter_Expecter{mock: &_m.Mock}
}

// CanTrade provides a mock function with given fields:
func (_m *MockAdapter) CanTrade() (bool, error) {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockAdapter_CanTrade_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CanTrade'
type MockAdapter_CanTrade_Call struct {
	*mock.Call
}

// CanTrade is a helper method to define mock.On call
func (_e *MockAdapter_Expecter) CanTrade() *MockAdapter_CanTrade_Call {
	return &MockAdapter_CanTrade_Call{Call: _e.mock.On("CanTrade")}
}

func (_c *MockAdapter_CanTrade_Call) Run(run func()) *MockAdapter_CanTrade_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockAdapter_CanTrade_Call) Return(_a0 bool, _a1 error) *MockAdapter_CanTrade_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// CancelPairOrder provides a mock function with given fields: pairSymbol, orderID, ctx
func (_m *MockAdapter) CancelPairOrder(pairSymbol string, orderID int64, ctx context.Context) error {
	ret := _m.Called(pairSymbol, orderID, ctx)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, int64, context.Context) error); ok {
		r0 = rf(pairSymbol, orderID, ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockAdapter_CancelPairOrder_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CancelPairOrder'
type MockAdapter_CancelPairOrder_Call struct {
	*mock.Call
}

// CancelPairOrder is a helper method to define mock.On call
//  - pairSymbol string
//  - orderID int64
//  - ctx context.Context
func (_e *MockAdapter_Expecter) CancelPairOrder(pairSymbol interface{}, orderID interface{}, ctx interface{}) *MockAdapter_CancelPairOrder_Call {
	return &MockAdapter_CancelPairOrder_Call{Call: _e.mock.On("CancelPairOrder", pairSymbol, orderID, ctx)}
}

func (_c *MockAdapter_CancelPairOrder_Call) Run(run func(pairSymbol string, orderID int64, ctx context.Context)) *MockAdapter_CancelPairOrder_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(int64), args[2].(context.Context))
	})
	return _c
}

func (_c *MockAdapter_CancelPairOrder_Call) Return(_a0 error) *MockAdapter_CancelPairOrder_Call {
	_c.Call.Return(_a0)
	return _c
}

// CancelPairOrderByClientOrderID provides a mock function with given fields: pairSymbol, clientOrderID, ctx
func (_m *MockAdapter) CancelPairOrderByClientOrderID(pairSymbol string, clientOrderID string, ctx context.Context) error {
	ret := _m.Called(pairSymbol, clientOrderID, ctx)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string, context.Context) error); ok {
		r0 = rf(pairSymbol, clientOrderID, ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockAdapter_CancelPairOrderByClientOrderID_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CancelPairOrderByClientOrderID'
type MockAdapter_CancelPairOrderByClientOrderID_Call struct {
	*mock.Call
}

// CancelPairOrderByClientOrderID is a helper method to define mock.On call
//  - pairSymbol string
//  - clientOrderID string
//  - ctx context.Context
func (_e *MockAdapter_Expecter) CancelPairOrderByClientOrderID(pairSymbol interface{}, clientOrderID interface{}, ctx interface{}) *MockAdapter_CancelPairOrderByClientOrderID_Call {
	return &MockAdapter_CancelPairOrderByClientOrderID_Call{Call: _e.mock.On("CancelPairOrderByClientOrderID", pairSymbol, clientOrderID, ctx)}
}

func (_c *MockAdapter_CancelPairOrderByClientOrderID_Call) Run(run func(pairSymbol string, clientOrderID string, ctx context.Context)) *MockAdapter_CancelPairOrderByClientOrderID_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(string), args[2].(context.Context))
	})
	return _c
}

func (_c *MockAdapter_CancelPairOrderByClientOrderID_Call) Return(_a0 error) *MockAdapter_CancelPairOrderByClientOrderID_Call {
	_c.Call.Return(_a0)
	return _c
}

// Connect provides a mock function with given fields: credentials
func (_m *MockAdapter) Connect(credentials structs.APICredentials) error {
	ret := _m.Called(credentials)

	var r0 error
	if rf, ok := ret.Get(0).(func(structs.APICredentials) error); ok {
		r0 = rf(credentials)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockAdapter_Connect_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Connect'
type MockAdapter_Connect_Call struct {
	*mock.Call
}

// Connect is a helper method to define mock.On call
//  - credentials structs.APICredentials
func (_e *MockAdapter_Expecter) Connect(credentials interface{}) *MockAdapter_Connect_Call {
	return &MockAdapter_Connect_Call{Call: _e.mock.On("Connect", credentials)}
}

func (_c *MockAdapter_Connect_Call) Run(run func(credentials structs.APICredentials)) *MockAdapter_Connect_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(structs.APICredentials))
	})
	return _c
}

func (_c *MockAdapter_Connect_Call) Return(_a0 error) *MockAdapter_Connect_Call {
	_c.Call.Return(_a0)
	return _c
}

// GetAccountBalance provides a mock function with given fields:
func (_m *MockAdapter) GetAccountBalance() ([]internalstructs.Balance, error) {
	ret := _m.Called()

	var r0 []internalstructs.Balance
	if rf, ok := ret.Get(0).(func() []internalstructs.Balance); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]internalstructs.Balance)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockAdapter_GetAccountBalance_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetAccountBalance'
type MockAdapter_GetAccountBalance_Call struct {
	*mock.Call
}

// GetAccountBalance is a helper method to define mock.On call
func (_e *MockAdapter_Expecter) GetAccountBalance() *MockAdapter_GetAccountBalance_Call {
	return &MockAdapter_GetAccountBalance_Call{Call: _e.mock.On("GetAccountBalance")}
}

func (_c *MockAdapter_GetAccountBalance_Call) Run(run func()) *MockAdapter_GetAccountBalance_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockAdapter_GetAccountBalance_Call) Return(_a0 []internalstructs.Balance, _a1 error) *MockAdapter_GetAccountBalance_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// GetCandleWorker provides a mock function with given fields:
func (_m *MockAdapter) GetCandleWorker() workers.ICandleWorker {
	ret := _m.Called()

	var r0 workers.ICandleWorker
	if rf, ok := ret.Get(0).(func() workers.ICandleWorker); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(workers.ICandleWorker)
		}
	}

	return r0
}

// MockAdapter_GetCandleWorker_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetCandleWorker'
type MockAdapter_GetCandleWorker_Call struct {
	*mock.Call
}

// GetCandleWorker is a helper method to define mock.On call
func (_e *MockAdapter_Expecter) GetCandleWorker() *MockAdapter_GetCandleWorker_Call {
	return &MockAdapter_GetCandleWorker_Call{Call: _e.mock.On("GetCandleWorker")}
}

func (_c *MockAdapter_GetCandleWorker_Call) Run(run func()) *MockAdapter_GetCandleWorker_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockAdapter_GetCandleWorker_Call) Return(_a0 workers.ICandleWorker) *MockAdapter_GetCandleWorker_Call {
	_c.Call.Return(_a0)
	return _c
}

// GetCandles provides a mock function with given fields: limit, symbol, interval
func (_m *MockAdapter) GetCandles(limit int, symbol string, interval string) ([]workers.CandleData, error) {
	ret := _m.Called(limit, symbol, interval)

	var r0 []workers.CandleData
	if rf, ok := ret.Get(0).(func(int, string, string) []workers.CandleData); ok {
		r0 = rf(limit, symbol, interval)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]workers.CandleData)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int, string, string) error); ok {
		r1 = rf(limit, symbol, interval)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockAdapter_GetCandles_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetCandles'
type MockAdapter_GetCandles_Call struct {
	*mock.Call
}

// GetCandles is a helper method to define mock.On call
//  - limit int
//  - symbol string
//  - interval string
func (_e *MockAdapter_Expecter) GetCandles(limit interface{}, symbol interface{}, interval interface{}) *MockAdapter_GetCandles_Call {
	return &MockAdapter_GetCandles_Call{Call: _e.mock.On("GetCandles", limit, symbol, interval)}
}

func (_c *MockAdapter_GetCandles_Call) Run(run func(limit int, symbol string, interval string)) *MockAdapter_GetCandles_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(int), args[1].(string), args[2].(string))
	})
	return _c
}

func (_c *MockAdapter_GetCandles_Call) Return(_a0 []workers.CandleData, _a1 error) *MockAdapter_GetCandles_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// GetID provides a mock function with given fields:
func (_m *MockAdapter) GetID() int {
	ret := _m.Called()

	var r0 int
	if rf, ok := ret.Get(0).(func() int); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int)
	}

	return r0
}

// MockAdapter_GetID_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetID'
type MockAdapter_GetID_Call struct {
	*mock.Call
}

// GetID is a helper method to define mock.On call
func (_e *MockAdapter_Expecter) GetID() *MockAdapter_GetID_Call {
	return &MockAdapter_GetID_Call{Call: _e.mock.On("GetID")}
}

func (_c *MockAdapter_GetID_Call) Run(run func()) *MockAdapter_GetID_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockAdapter_GetID_Call) Return(_a0 int) *MockAdapter_GetID_Call {
	_c.Call.Return(_a0)
	return _c
}

// GetName provides a mock function with given fields:
func (_m *MockAdapter) GetName() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// MockAdapter_GetName_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetName'
type MockAdapter_GetName_Call struct {
	*mock.Call
}

// GetName is a helper method to define mock.On call
func (_e *MockAdapter_Expecter) GetName() *MockAdapter_GetName_Call {
	return &MockAdapter_GetName_Call{Call: _e.mock.On("GetName")}
}

func (_c *MockAdapter_GetName_Call) Run(run func()) *MockAdapter_GetName_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockAdapter_GetName_Call) Return(_a0 string) *MockAdapter_GetName_Call {
	_c.Call.Return(_a0)
	return _c
}

// GetOrderByClientOrderID provides a mock function with given fields: pairSymbol, clientOrderID
func (_m *MockAdapter) GetOrderByClientOrderID(pairSymbol string, clientOrderID string) (internalstructs.OrderData, error) {
	ret := _m.Called(pairSymbol, clientOrderID)

	var r0 internalstructs.OrderData
	if rf, ok := ret.Get(0).(func(string, string) internalstructs.OrderData); ok {
		r0 = rf(pairSymbol, clientOrderID)
	} else {
		r0 = ret.Get(0).(internalstructs.OrderData)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(pairSymbol, clientOrderID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockAdapter_GetOrderByClientOrderID_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetOrderByClientOrderID'
type MockAdapter_GetOrderByClientOrderID_Call struct {
	*mock.Call
}

// GetOrderByClientOrderID is a helper method to define mock.On call
//  - pairSymbol string
//  - clientOrderID string
func (_e *MockAdapter_Expecter) GetOrderByClientOrderID(pairSymbol interface{}, clientOrderID interface{}) *MockAdapter_GetOrderByClientOrderID_Call {
	return &MockAdapter_GetOrderByClientOrderID_Call{Call: _e.mock.On("GetOrderByClientOrderID", pairSymbol, clientOrderID)}
}

func (_c *MockAdapter_GetOrderByClientOrderID_Call) Run(run func(pairSymbol string, clientOrderID string)) *MockAdapter_GetOrderByClientOrderID_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(string))
	})
	return _c
}

func (_c *MockAdapter_GetOrderByClientOrderID_Call) Return(_a0 internalstructs.OrderData, _a1 error) *MockAdapter_GetOrderByClientOrderID_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// GetOrderData provides a mock function with given fields: pairSymbol, orderID
func (_m *MockAdapter) GetOrderData(pairSymbol string, orderID int64) (internalstructs.OrderData, error) {
	ret := _m.Called(pairSymbol, orderID)

	var r0 internalstructs.OrderData
	if rf, ok := ret.Get(0).(func(string, int64) internalstructs.OrderData); ok {
		r0 = rf(pairSymbol, orderID)
	} else {
		r0 = ret.Get(0).(internalstructs.OrderData)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, int64) error); ok {
		r1 = rf(pairSymbol, orderID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockAdapter_GetOrderData_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetOrderData'
type MockAdapter_GetOrderData_Call struct {
	*mock.Call
}

// GetOrderData is a helper method to define mock.On call
//  - pairSymbol string
//  - orderID int64
func (_e *MockAdapter_Expecter) GetOrderData(pairSymbol interface{}, orderID interface{}) *MockAdapter_GetOrderData_Call {
	return &MockAdapter_GetOrderData_Call{Call: _e.mock.On("GetOrderData", pairSymbol, orderID)}
}

func (_c *MockAdapter_GetOrderData_Call) Run(run func(pairSymbol string, orderID int64)) *MockAdapter_GetOrderData_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(int64))
	})
	return _c
}

func (_c *MockAdapter_GetOrderData_Call) Return(_a0 internalstructs.OrderData, _a1 error) *MockAdapter_GetOrderData_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// GetOrderExecFee provides a mock function with given fields: baseAssetTicker, quoteAssetTicker, orderSide, orderID
func (_m *MockAdapter) GetOrderExecFee(baseAssetTicker string, quoteAssetTicker string, orderSide string, orderID int64) (internalstructs.OrderFees, error) {
	ret := _m.Called(baseAssetTicker, quoteAssetTicker, orderSide, orderID)

	var r0 internalstructs.OrderFees
	if rf, ok := ret.Get(0).(func(string, string, string, int64) internalstructs.OrderFees); ok {
		r0 = rf(baseAssetTicker, quoteAssetTicker, orderSide, orderID)
	} else {
		r0 = ret.Get(0).(internalstructs.OrderFees)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string, string, int64) error); ok {
		r1 = rf(baseAssetTicker, quoteAssetTicker, orderSide, orderID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockAdapter_GetOrderExecFee_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetOrderExecFee'
type MockAdapter_GetOrderExecFee_Call struct {
	*mock.Call
}

// GetOrderExecFee is a helper method to define mock.On call
//  - baseAssetTicker string
//  - quoteAssetTicker string
//  - orderSide string
//  - orderID int64
func (_e *MockAdapter_Expecter) GetOrderExecFee(baseAssetTicker interface{}, quoteAssetTicker interface{}, orderSide interface{}, orderID interface{}) *MockAdapter_GetOrderExecFee_Call {
	return &MockAdapter_GetOrderExecFee_Call{Call: _e.mock.On("GetOrderExecFee", baseAssetTicker, quoteAssetTicker, orderSide, orderID)}
}

func (_c *MockAdapter_GetOrderExecFee_Call) Run(run func(baseAssetTicker string, quoteAssetTicker string, orderSide string, orderID int64)) *MockAdapter_GetOrderExecFee_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(string), args[2].(string), args[3].(int64))
	})
	return _c
}

func (_c *MockAdapter_GetOrderExecFee_Call) Return(_a0 internalstructs.OrderFees, _a1 error) *MockAdapter_GetOrderExecFee_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// GetPairBalance provides a mock function with given fields: pair
func (_m *MockAdapter) GetPairBalance(pair internalstructs.PairSymbolData) (internalstructs.PairBalance, error) {
	ret := _m.Called(pair)

	var r0 internalstructs.PairBalance
	if rf, ok := ret.Get(0).(func(internalstructs.PairSymbolData) internalstructs.PairBalance); ok {
		r0 = rf(pair)
	} else {
		r0 = ret.Get(0).(internalstructs.PairBalance)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(internalstructs.PairSymbolData) error); ok {
		r1 = rf(pair)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockAdapter_GetPairBalance_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetPairBalance'
type MockAdapter_GetPairBalance_Call struct {
	*mock.Call
}

// GetPairBalance is a helper method to define mock.On call
//  - pair internalstructs.PairSymbolData
func (_e *MockAdapter_Expecter) GetPairBalance(pair interface{}) *MockAdapter_GetPairBalance_Call {
	return &MockAdapter_GetPairBalance_Call{Call: _e.mock.On("GetPairBalance", pair)}
}

func (_c *MockAdapter_GetPairBalance_Call) Run(run func(pair internalstructs.PairSymbolData)) *MockAdapter_GetPairBalance_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(internalstructs.PairSymbolData))
	})
	return _c
}

func (_c *MockAdapter_GetPairBalance_Call) Return(_a0 internalstructs.PairBalance, _a1 error) *MockAdapter_GetPairBalance_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// GetPairData provides a mock function with given fields: pairSymbol
func (_m *MockAdapter) GetPairData(pairSymbol string) (internalstructs.ExchangePairData, error) {
	ret := _m.Called(pairSymbol)

	var r0 internalstructs.ExchangePairData
	if rf, ok := ret.Get(0).(func(string) internalstructs.ExchangePairData); ok {
		r0 = rf(pairSymbol)
	} else {
		r0 = ret.Get(0).(internalstructs.ExchangePairData)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(pairSymbol)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockAdapter_GetPairData_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetPairData'
type MockAdapter_GetPairData_Call struct {
	*mock.Call
}

// GetPairData is a helper method to define mock.On call
//  - pairSymbol string
func (_e *MockAdapter_Expecter) GetPairData(pairSymbol interface{}) *MockAdapter_GetPairData_Call {
	return &MockAdapter_GetPairData_Call{Call: _e.mock.On("GetPairData", pairSymbol)}
}

func (_c *MockAdapter_GetPairData_Call) Run(run func(pairSymbol string)) *MockAdapter_GetPairData_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *MockAdapter_GetPairData_Call) Return(_a0 internalstructs.ExchangePairData, _a1 error) *MockAdapter_GetPairData_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// GetPairLastPrice provides a mock function with given fields: pairSymbol
func (_m *MockAdapter) GetPairLastPrice(pairSymbol string) (float64, error) {
	ret := _m.Called(pairSymbol)

	var r0 float64
	if rf, ok := ret.Get(0).(func(string) float64); ok {
		r0 = rf(pairSymbol)
	} else {
		r0 = ret.Get(0).(float64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(pairSymbol)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockAdapter_GetPairLastPrice_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetPairLastPrice'
type MockAdapter_GetPairLastPrice_Call struct {
	*mock.Call
}

// GetPairLastPrice is a helper method to define mock.On call
//  - pairSymbol string
func (_e *MockAdapter_Expecter) GetPairLastPrice(pairSymbol interface{}) *MockAdapter_GetPairLastPrice_Call {
	return &MockAdapter_GetPairLastPrice_Call{Call: _e.mock.On("GetPairLastPrice", pairSymbol)}
}

func (_c *MockAdapter_GetPairLastPrice_Call) Run(run func(pairSymbol string)) *MockAdapter_GetPairLastPrice_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *MockAdapter_GetPairLastPrice_Call) Return(_a0 float64, _a1 error) *MockAdapter_GetPairLastPrice_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// GetPairOpenOrders provides a mock function with given fields: pairSymbol
func (_m *MockAdapter) GetPairOpenOrders(pairSymbol string) ([]internalstructs.OrderData, error) {
	ret := _m.Called(pairSymbol)

	var r0 []internalstructs.OrderData
	if rf, ok := ret.Get(0).(func(string) []internalstructs.OrderData); ok {
		r0 = rf(pairSymbol)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]internalstructs.OrderData)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(pairSymbol)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockAdapter_GetPairOpenOrders_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetPairOpenOrders'
type MockAdapter_GetPairOpenOrders_Call struct {
	*mock.Call
}

// GetPairOpenOrders is a helper method to define mock.On call
//  - pairSymbol string
func (_e *MockAdapter_Expecter) GetPairOpenOrders(pairSymbol interface{}) *MockAdapter_GetPairOpenOrders_Call {
	return &MockAdapter_GetPairOpenOrders_Call{Call: _e.mock.On("GetPairOpenOrders", pairSymbol)}
}

func (_c *MockAdapter_GetPairOpenOrders_Call) Run(run func(pairSymbol string)) *MockAdapter_GetPairOpenOrders_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *MockAdapter_GetPairOpenOrders_Call) Return(_a0 []internalstructs.OrderData, _a1 error) *MockAdapter_GetPairOpenOrders_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// GetPairs provides a mock function with given fields:
func (_m *MockAdapter) GetPairs() ([]internalstructs.ExchangePairData, error) {
	ret := _m.Called()

	var r0 []internalstructs.ExchangePairData
	if rf, ok := ret.Get(0).(func() []internalstructs.ExchangePairData); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]internalstructs.ExchangePairData)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockAdapter_GetPairs_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetPairs'
type MockAdapter_GetPairs_Call struct {
	*mock.Call
}

// GetPairs is a helper method to define mock.On call
func (_e *MockAdapter_Expecter) GetPairs() *MockAdapter_GetPairs_Call {
	return &MockAdapter_GetPairs_Call{Call: _e.mock.On("GetPairs")}
}

func (_c *MockAdapter_GetPairs_Call) Run(run func()) *MockAdapter_GetPairs_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockAdapter_GetPairs_Call) Return(_a0 []internalstructs.ExchangePairData, _a1 error) *MockAdapter_GetPairs_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// GetPriceWorker provides a mock function with given fields: callback
func (_m *MockAdapter) GetPriceWorker(callback workers.PriceEventCallback) workers.IPriceWorker {
	ret := _m.Called(callback)

	var r0 workers.IPriceWorker
	if rf, ok := ret.Get(0).(func(workers.PriceEventCallback) workers.IPriceWorker); ok {
		r0 = rf(callback)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(workers.IPriceWorker)
		}
	}

	return r0
}

// MockAdapter_GetPriceWorker_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetPriceWorker'
type MockAdapter_GetPriceWorker_Call struct {
	*mock.Call
}

// GetPriceWorker is a helper method to define mock.On call
//  - callback workers.PriceEventCallback
func (_e *MockAdapter_Expecter) GetPriceWorker(callback interface{}) *MockAdapter_GetPriceWorker_Call {
	return &MockAdapter_GetPriceWorker_Call{Call: _e.mock.On("GetPriceWorker", callback)}
}

func (_c *MockAdapter_GetPriceWorker_Call) Run(run func(callback workers.PriceEventCallback)) *MockAdapter_GetPriceWorker_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(workers.PriceEventCallback))
	})
	return _c
}

func (_c *MockAdapter_GetPriceWorker_Call) Return(_a0 workers.IPriceWorker) *MockAdapter_GetPriceWorker_Call {
	_c.Call.Return(_a0)
	return _c
}

// GetTag provides a mock function with given fields:
func (_m *MockAdapter) GetTag() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// MockAdapter_GetTag_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetTag'
type MockAdapter_GetTag_Call struct {
	*mock.Call
}

// GetTag is a helper method to define mock.On call
func (_e *MockAdapter_Expecter) GetTag() *MockAdapter_GetTag_Call {
	return &MockAdapter_GetTag_Call{Call: _e.mock.On("GetTag")}
}

func (_c *MockAdapter_GetTag_Call) Run(run func()) *MockAdapter_GetTag_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockAdapter_GetTag_Call) Return(_a0 string) *MockAdapter_GetTag_Call {
	_c.Call.Return(_a0)
	return _c
}

// GetTradeEventsWorker provides a mock function with given fields:
func (_m *MockAdapter) GetTradeEventsWorker() workers.ITradeEventWorker {
	ret := _m.Called()

	var r0 workers.ITradeEventWorker
	if rf, ok := ret.Get(0).(func() workers.ITradeEventWorker); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(workers.ITradeEventWorker)
		}
	}

	return r0
}

// MockAdapter_GetTradeEventsWorker_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetTradeEventsWorker'
type MockAdapter_GetTradeEventsWorker_Call struct {
	*mock.Call
}

// GetTradeEventsWorker is a helper method to define mock.On call
func (_e *MockAdapter_Expecter) GetTradeEventsWorker() *MockAdapter_GetTradeEventsWorker_Call {
	return &MockAdapter_GetTradeEventsWorker_Call{Call: _e.mock.On("GetTradeEventsWorker")}
}

func (_c *MockAdapter_GetTradeEventsWorker_Call) Run(run func()) *MockAdapter_GetTradeEventsWorker_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockAdapter_GetTradeEventsWorker_Call) Return(_a0 workers.ITradeEventWorker) *MockAdapter_GetTradeEventsWorker_Call {
	_c.Call.Return(_a0)
	return _c
}

// PlaceOrder provides a mock function with given fields: ctx, order
func (_m *MockAdapter) PlaceOrder(ctx context.Context, order internalstructs.BotOrderAdjusted) (internalstructs.CreateOrderResponse, error) {
	ret := _m.Called(ctx, order)

	var r0 internalstructs.CreateOrderResponse
	if rf, ok := ret.Get(0).(func(context.Context, internalstructs.BotOrderAdjusted) internalstructs.CreateOrderResponse); ok {
		r0 = rf(ctx, order)
	} else {
		r0 = ret.Get(0).(internalstructs.CreateOrderResponse)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, internalstructs.BotOrderAdjusted) error); ok {
		r1 = rf(ctx, order)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockAdapter_PlaceOrder_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'PlaceOrder'
type MockAdapter_PlaceOrder_Call struct {
	*mock.Call
}

// PlaceOrder is a helper method to define mock.On call
//  - ctx context.Context
//  - order internalstructs.BotOrderAdjusted
func (_e *MockAdapter_Expecter) PlaceOrder(ctx interface{}, order interface{}) *MockAdapter_PlaceOrder_Call {
	return &MockAdapter_PlaceOrder_Call{Call: _e.mock.On("PlaceOrder", ctx, order)}
}

func (_c *MockAdapter_PlaceOrder_Call) Run(run func(ctx context.Context, order internalstructs.BotOrderAdjusted)) *MockAdapter_PlaceOrder_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(internalstructs.BotOrderAdjusted))
	})
	return _c
}

func (_c *MockAdapter_PlaceOrder_Call) Return(_a0 internalstructs.CreateOrderResponse, _a1 error) *MockAdapter_PlaceOrder_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// VerifyAPIKeys provides a mock function with given fields: keyPublic, keySecret
func (_m *MockAdapter) VerifyAPIKeys(keyPublic string, keySecret string) error {
	ret := _m.Called(keyPublic, keySecret)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string) error); ok {
		r0 = rf(keyPublic, keySecret)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockAdapter_VerifyAPIKeys_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'VerifyAPIKeys'
type MockAdapter_VerifyAPIKeys_Call struct {
	*mock.Call
}

// VerifyAPIKeys is a helper method to define mock.On call
//  - keyPublic string
//  - keySecret string
func (_e *MockAdapter_Expecter) VerifyAPIKeys(keyPublic interface{}, keySecret interface{}) *MockAdapter_VerifyAPIKeys_Call {
	return &MockAdapter_VerifyAPIKeys_Call{Call: _e.mock.On("VerifyAPIKeys", keyPublic, keySecret)}
}

func (_c *MockAdapter_VerifyAPIKeys_Call) Run(run func(keyPublic string, keySecret string)) *MockAdapter_VerifyAPIKeys_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(string))
	})
	return _c
}

func (_c *MockAdapter_VerifyAPIKeys_Call) Return(_a0 error) *MockAdapter_VerifyAPIKeys_Call {
	_c.Call.Return(_a0)
	return _c
}

type mockConstructorTestingTNewMockAdapter interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockAdapter creates a new instance of MockAdapter. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockAdapter(t mockConstructorTestingTNewMockAdapter) *MockAdapter {
	mock := &MockAdapter{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
