// Code generated by mockery v2.14.1. DO NOT EDIT.

package wrapper

import (
	context "context"

	binance "github.com/adshao/go-binance/v2"

	mock "github.com/stretchr/testify/mock"

	workers "github.com/matrixbotio/exchange-gates-lib/internal/workers"
)

// MockBinanceAPIWrapper is an autogenerated mock type for the BinanceAPIWrapper type
type MockBinanceAPIWrapper struct {
	mock.Mock
}

type MockBinanceAPIWrapper_Expecter struct {
	mock *mock.Mock
}

func (_m *MockBinanceAPIWrapper) EXPECT() *MockBinanceAPIWrapper_Expecter {
	return &MockBinanceAPIWrapper_Expecter{mock: &_m.Mock}
}

// CancelOrderByClientOrderID provides a mock function with given fields: ctx, pairSymbol, clientOrderID
func (_m *MockBinanceAPIWrapper) CancelOrderByClientOrderID(ctx context.Context, pairSymbol string, clientOrderID string) error {
	ret := _m.Called(ctx, pairSymbol, clientOrderID)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) error); ok {
		r0 = rf(ctx, pairSymbol, clientOrderID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockBinanceAPIWrapper_CancelOrderByClientOrderID_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CancelOrderByClientOrderID'
type MockBinanceAPIWrapper_CancelOrderByClientOrderID_Call struct {
	*mock.Call
}

// CancelOrderByClientOrderID is a helper method to define mock.On call
//  - ctx context.Context
//  - pairSymbol string
//  - clientOrderID string
func (_e *MockBinanceAPIWrapper_Expecter) CancelOrderByClientOrderID(ctx interface{}, pairSymbol interface{}, clientOrderID interface{}) *MockBinanceAPIWrapper_CancelOrderByClientOrderID_Call {
	return &MockBinanceAPIWrapper_CancelOrderByClientOrderID_Call{Call: _e.mock.On("CancelOrderByClientOrderID", ctx, pairSymbol, clientOrderID)}
}

func (_c *MockBinanceAPIWrapper_CancelOrderByClientOrderID_Call) Run(run func(ctx context.Context, pairSymbol string, clientOrderID string)) *MockBinanceAPIWrapper_CancelOrderByClientOrderID_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(string))
	})
	return _c
}

func (_c *MockBinanceAPIWrapper_CancelOrderByClientOrderID_Call) Return(_a0 error) *MockBinanceAPIWrapper_CancelOrderByClientOrderID_Call {
	_c.Call.Return(_a0)
	return _c
}

// CancelOrderByID provides a mock function with given fields: ctx, pairSymbol, orderID
func (_m *MockBinanceAPIWrapper) CancelOrderByID(ctx context.Context, pairSymbol string, orderID int64) error {
	ret := _m.Called(ctx, pairSymbol, orderID)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, int64) error); ok {
		r0 = rf(ctx, pairSymbol, orderID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockBinanceAPIWrapper_CancelOrderByID_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CancelOrderByID'
type MockBinanceAPIWrapper_CancelOrderByID_Call struct {
	*mock.Call
}

// CancelOrderByID is a helper method to define mock.On call
//  - ctx context.Context
//  - pairSymbol string
//  - orderID int64
func (_e *MockBinanceAPIWrapper_Expecter) CancelOrderByID(ctx interface{}, pairSymbol interface{}, orderID interface{}) *MockBinanceAPIWrapper_CancelOrderByID_Call {
	return &MockBinanceAPIWrapper_CancelOrderByID_Call{Call: _e.mock.On("CancelOrderByID", ctx, pairSymbol, orderID)}
}

func (_c *MockBinanceAPIWrapper_CancelOrderByID_Call) Run(run func(ctx context.Context, pairSymbol string, orderID int64)) *MockBinanceAPIWrapper_CancelOrderByID_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(int64))
	})
	return _c
}

func (_c *MockBinanceAPIWrapper_CancelOrderByID_Call) Return(_a0 error) *MockBinanceAPIWrapper_CancelOrderByID_Call {
	_c.Call.Return(_a0)
	return _c
}

// Connect provides a mock function with given fields: ctx, keyPublic, keySecret
func (_m *MockBinanceAPIWrapper) Connect(ctx context.Context, keyPublic string, keySecret string) error {
	ret := _m.Called(ctx, keyPublic, keySecret)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) error); ok {
		r0 = rf(ctx, keyPublic, keySecret)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockBinanceAPIWrapper_Connect_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Connect'
type MockBinanceAPIWrapper_Connect_Call struct {
	*mock.Call
}

// Connect is a helper method to define mock.On call
//  - ctx context.Context
//  - keyPublic string
//  - keySecret string
func (_e *MockBinanceAPIWrapper_Expecter) Connect(ctx interface{}, keyPublic interface{}, keySecret interface{}) *MockBinanceAPIWrapper_Connect_Call {
	return &MockBinanceAPIWrapper_Connect_Call{Call: _e.mock.On("Connect", ctx, keyPublic, keySecret)}
}

func (_c *MockBinanceAPIWrapper_Connect_Call) Run(run func(ctx context.Context, keyPublic string, keySecret string)) *MockBinanceAPIWrapper_Connect_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(string))
	})
	return _c
}

func (_c *MockBinanceAPIWrapper_Connect_Call) Return(_a0 error) *MockBinanceAPIWrapper_Connect_Call {
	_c.Call.Return(_a0)
	return _c
}

// GetAccountData provides a mock function with given fields: _a0
func (_m *MockBinanceAPIWrapper) GetAccountData(_a0 context.Context) (*binance.Account, error) {
	ret := _m.Called(_a0)

	var r0 *binance.Account
	if rf, ok := ret.Get(0).(func(context.Context) *binance.Account); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*binance.Account)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockBinanceAPIWrapper_GetAccountData_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetAccountData'
type MockBinanceAPIWrapper_GetAccountData_Call struct {
	*mock.Call
}

// GetAccountData is a helper method to define mock.On call
//  - _a0 context.Context
func (_e *MockBinanceAPIWrapper_Expecter) GetAccountData(_a0 interface{}) *MockBinanceAPIWrapper_GetAccountData_Call {
	return &MockBinanceAPIWrapper_GetAccountData_Call{Call: _e.mock.On("GetAccountData", _a0)}
}

func (_c *MockBinanceAPIWrapper_GetAccountData_Call) Run(run func(_a0 context.Context)) *MockBinanceAPIWrapper_GetAccountData_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *MockBinanceAPIWrapper_GetAccountData_Call) Return(_a0 *binance.Account, _a1 error) *MockBinanceAPIWrapper_GetAccountData_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// GetExchangeInfo provides a mock function with given fields: ctx, pairSymbol
func (_m *MockBinanceAPIWrapper) GetExchangeInfo(ctx context.Context, pairSymbol string) (*binance.ExchangeInfo, error) {
	ret := _m.Called(ctx, pairSymbol)

	var r0 *binance.ExchangeInfo
	if rf, ok := ret.Get(0).(func(context.Context, string) *binance.ExchangeInfo); ok {
		r0 = rf(ctx, pairSymbol)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*binance.ExchangeInfo)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, pairSymbol)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockBinanceAPIWrapper_GetExchangeInfo_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetExchangeInfo'
type MockBinanceAPIWrapper_GetExchangeInfo_Call struct {
	*mock.Call
}

// GetExchangeInfo is a helper method to define mock.On call
//  - ctx context.Context
//  - pairSymbol string
func (_e *MockBinanceAPIWrapper_Expecter) GetExchangeInfo(ctx interface{}, pairSymbol interface{}) *MockBinanceAPIWrapper_GetExchangeInfo_Call {
	return &MockBinanceAPIWrapper_GetExchangeInfo_Call{Call: _e.mock.On("GetExchangeInfo", ctx, pairSymbol)}
}

func (_c *MockBinanceAPIWrapper_GetExchangeInfo_Call) Run(run func(ctx context.Context, pairSymbol string)) *MockBinanceAPIWrapper_GetExchangeInfo_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *MockBinanceAPIWrapper_GetExchangeInfo_Call) Return(_a0 *binance.ExchangeInfo, _a1 error) *MockBinanceAPIWrapper_GetExchangeInfo_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// GetKlines provides a mock function with given fields: ctx, pairSymbol, interval, limit
func (_m *MockBinanceAPIWrapper) GetKlines(ctx context.Context, pairSymbol string, interval string, limit int) ([]*binance.Kline, error) {
	ret := _m.Called(ctx, pairSymbol, interval, limit)

	var r0 []*binance.Kline
	if rf, ok := ret.Get(0).(func(context.Context, string, string, int) []*binance.Kline); ok {
		r0 = rf(ctx, pairSymbol, interval, limit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*binance.Kline)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string, int) error); ok {
		r1 = rf(ctx, pairSymbol, interval, limit)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockBinanceAPIWrapper_GetKlines_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetKlines'
type MockBinanceAPIWrapper_GetKlines_Call struct {
	*mock.Call
}

// GetKlines is a helper method to define mock.On call
//  - ctx context.Context
//  - pairSymbol string
//  - interval string
//  - limit int
func (_e *MockBinanceAPIWrapper_Expecter) GetKlines(ctx interface{}, pairSymbol interface{}, interval interface{}, limit interface{}) *MockBinanceAPIWrapper_GetKlines_Call {
	return &MockBinanceAPIWrapper_GetKlines_Call{Call: _e.mock.On("GetKlines", ctx, pairSymbol, interval, limit)}
}

func (_c *MockBinanceAPIWrapper_GetKlines_Call) Run(run func(ctx context.Context, pairSymbol string, interval string, limit int)) *MockBinanceAPIWrapper_GetKlines_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(string), args[3].(int))
	})
	return _c
}

func (_c *MockBinanceAPIWrapper_GetKlines_Call) Return(_a0 []*binance.Kline, _a1 error) *MockBinanceAPIWrapper_GetKlines_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// GetOpenOrders provides a mock function with given fields: ctx, pairSymbol
func (_m *MockBinanceAPIWrapper) GetOpenOrders(ctx context.Context, pairSymbol string) ([]*binance.Order, error) {
	ret := _m.Called(ctx, pairSymbol)

	var r0 []*binance.Order
	if rf, ok := ret.Get(0).(func(context.Context, string) []*binance.Order); ok {
		r0 = rf(ctx, pairSymbol)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*binance.Order)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, pairSymbol)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockBinanceAPIWrapper_GetOpenOrders_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetOpenOrders'
type MockBinanceAPIWrapper_GetOpenOrders_Call struct {
	*mock.Call
}

// GetOpenOrders is a helper method to define mock.On call
//  - ctx context.Context
//  - pairSymbol string
func (_e *MockBinanceAPIWrapper_Expecter) GetOpenOrders(ctx interface{}, pairSymbol interface{}) *MockBinanceAPIWrapper_GetOpenOrders_Call {
	return &MockBinanceAPIWrapper_GetOpenOrders_Call{Call: _e.mock.On("GetOpenOrders", ctx, pairSymbol)}
}

func (_c *MockBinanceAPIWrapper_GetOpenOrders_Call) Run(run func(ctx context.Context, pairSymbol string)) *MockBinanceAPIWrapper_GetOpenOrders_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *MockBinanceAPIWrapper_GetOpenOrders_Call) Return(_a0 []*binance.Order, _a1 error) *MockBinanceAPIWrapper_GetOpenOrders_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// GetOrderDataByClientOrderID provides a mock function with given fields: ctx, pairSymbol, clientOrderID
func (_m *MockBinanceAPIWrapper) GetOrderDataByClientOrderID(ctx context.Context, pairSymbol string, clientOrderID string) (*binance.Order, error) {
	ret := _m.Called(ctx, pairSymbol, clientOrderID)

	var r0 *binance.Order
	if rf, ok := ret.Get(0).(func(context.Context, string, string) *binance.Order); ok {
		r0 = rf(ctx, pairSymbol, clientOrderID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*binance.Order)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, pairSymbol, clientOrderID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockBinanceAPIWrapper_GetOrderDataByClientOrderID_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetOrderDataByClientOrderID'
type MockBinanceAPIWrapper_GetOrderDataByClientOrderID_Call struct {
	*mock.Call
}

// GetOrderDataByClientOrderID is a helper method to define mock.On call
//  - ctx context.Context
//  - pairSymbol string
//  - clientOrderID string
func (_e *MockBinanceAPIWrapper_Expecter) GetOrderDataByClientOrderID(ctx interface{}, pairSymbol interface{}, clientOrderID interface{}) *MockBinanceAPIWrapper_GetOrderDataByClientOrderID_Call {
	return &MockBinanceAPIWrapper_GetOrderDataByClientOrderID_Call{Call: _e.mock.On("GetOrderDataByClientOrderID", ctx, pairSymbol, clientOrderID)}
}

func (_c *MockBinanceAPIWrapper_GetOrderDataByClientOrderID_Call) Run(run func(ctx context.Context, pairSymbol string, clientOrderID string)) *MockBinanceAPIWrapper_GetOrderDataByClientOrderID_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(string))
	})
	return _c
}

func (_c *MockBinanceAPIWrapper_GetOrderDataByClientOrderID_Call) Return(_a0 *binance.Order, _a1 error) *MockBinanceAPIWrapper_GetOrderDataByClientOrderID_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// GetOrderDataByOrderID provides a mock function with given fields: ctx, pairSymbol, orderID
func (_m *MockBinanceAPIWrapper) GetOrderDataByOrderID(ctx context.Context, pairSymbol string, orderID int64) (*binance.Order, error) {
	ret := _m.Called(ctx, pairSymbol, orderID)

	var r0 *binance.Order
	if rf, ok := ret.Get(0).(func(context.Context, string, int64) *binance.Order); ok {
		r0 = rf(ctx, pairSymbol, orderID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*binance.Order)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, int64) error); ok {
		r1 = rf(ctx, pairSymbol, orderID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockBinanceAPIWrapper_GetOrderDataByOrderID_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetOrderDataByOrderID'
type MockBinanceAPIWrapper_GetOrderDataByOrderID_Call struct {
	*mock.Call
}

// GetOrderDataByOrderID is a helper method to define mock.On call
//  - ctx context.Context
//  - pairSymbol string
//  - orderID int64
func (_e *MockBinanceAPIWrapper_Expecter) GetOrderDataByOrderID(ctx interface{}, pairSymbol interface{}, orderID interface{}) *MockBinanceAPIWrapper_GetOrderDataByOrderID_Call {
	return &MockBinanceAPIWrapper_GetOrderDataByOrderID_Call{Call: _e.mock.On("GetOrderDataByOrderID", ctx, pairSymbol, orderID)}
}

func (_c *MockBinanceAPIWrapper_GetOrderDataByOrderID_Call) Run(run func(ctx context.Context, pairSymbol string, orderID int64)) *MockBinanceAPIWrapper_GetOrderDataByOrderID_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(int64))
	})
	return _c
}

func (_c *MockBinanceAPIWrapper_GetOrderDataByOrderID_Call) Return(_a0 *binance.Order, _a1 error) *MockBinanceAPIWrapper_GetOrderDataByOrderID_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// GetOrderTradeHistory provides a mock function with given fields: ctx, orderID, pairSymbol
func (_m *MockBinanceAPIWrapper) GetOrderTradeHistory(ctx context.Context, orderID int64, pairSymbol string) ([]*binance.TradeV3, error) {
	ret := _m.Called(ctx, orderID, pairSymbol)

	var r0 []*binance.TradeV3
	if rf, ok := ret.Get(0).(func(context.Context, int64, string) []*binance.TradeV3); ok {
		r0 = rf(ctx, orderID, pairSymbol)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*binance.TradeV3)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int64, string) error); ok {
		r1 = rf(ctx, orderID, pairSymbol)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockBinanceAPIWrapper_GetOrderTradeHistory_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetOrderTradeHistory'
type MockBinanceAPIWrapper_GetOrderTradeHistory_Call struct {
	*mock.Call
}

// GetOrderTradeHistory is a helper method to define mock.On call
//  - ctx context.Context
//  - orderID int64
//  - pairSymbol string
func (_e *MockBinanceAPIWrapper_Expecter) GetOrderTradeHistory(ctx interface{}, orderID interface{}, pairSymbol interface{}) *MockBinanceAPIWrapper_GetOrderTradeHistory_Call {
	return &MockBinanceAPIWrapper_GetOrderTradeHistory_Call{Call: _e.mock.On("GetOrderTradeHistory", ctx, orderID, pairSymbol)}
}

func (_c *MockBinanceAPIWrapper_GetOrderTradeHistory_Call) Run(run func(ctx context.Context, orderID int64, pairSymbol string)) *MockBinanceAPIWrapper_GetOrderTradeHistory_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(int64), args[2].(string))
	})
	return _c
}

func (_c *MockBinanceAPIWrapper_GetOrderTradeHistory_Call) Return(_a0 []*binance.TradeV3, _a1 error) *MockBinanceAPIWrapper_GetOrderTradeHistory_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// GetPrices provides a mock function with given fields: ctx, pairSymbol
func (_m *MockBinanceAPIWrapper) GetPrices(ctx context.Context, pairSymbol string) ([]*binance.SymbolPrice, error) {
	ret := _m.Called(ctx, pairSymbol)

	var r0 []*binance.SymbolPrice
	if rf, ok := ret.Get(0).(func(context.Context, string) []*binance.SymbolPrice); ok {
		r0 = rf(ctx, pairSymbol)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*binance.SymbolPrice)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, pairSymbol)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockBinanceAPIWrapper_GetPrices_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetPrices'
type MockBinanceAPIWrapper_GetPrices_Call struct {
	*mock.Call
}

// GetPrices is a helper method to define mock.On call
//  - ctx context.Context
//  - pairSymbol string
func (_e *MockBinanceAPIWrapper_Expecter) GetPrices(ctx interface{}, pairSymbol interface{}) *MockBinanceAPIWrapper_GetPrices_Call {
	return &MockBinanceAPIWrapper_GetPrices_Call{Call: _e.mock.On("GetPrices", ctx, pairSymbol)}
}

func (_c *MockBinanceAPIWrapper_GetPrices_Call) Run(run func(ctx context.Context, pairSymbol string)) *MockBinanceAPIWrapper_GetPrices_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *MockBinanceAPIWrapper_GetPrices_Call) Return(_a0 []*binance.SymbolPrice, _a1 error) *MockBinanceAPIWrapper_GetPrices_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// Ping provides a mock function with given fields: _a0
func (_m *MockBinanceAPIWrapper) Ping(_a0 context.Context) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockBinanceAPIWrapper_Ping_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Ping'
type MockBinanceAPIWrapper_Ping_Call struct {
	*mock.Call
}

// Ping is a helper method to define mock.On call
//  - _a0 context.Context
func (_e *MockBinanceAPIWrapper_Expecter) Ping(_a0 interface{}) *MockBinanceAPIWrapper_Ping_Call {
	return &MockBinanceAPIWrapper_Ping_Call{Call: _e.mock.On("Ping", _a0)}
}

func (_c *MockBinanceAPIWrapper_Ping_Call) Run(run func(_a0 context.Context)) *MockBinanceAPIWrapper_Ping_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *MockBinanceAPIWrapper_Ping_Call) Return(_a0 error) *MockBinanceAPIWrapper_Ping_Call {
	_c.Call.Return(_a0)
	return _c
}

// PlaceLimitOrder provides a mock function with given fields: ctx, pairSymbol, orderSide, qty, price, optionalClientOrderID
func (_m *MockBinanceAPIWrapper) PlaceLimitOrder(ctx context.Context, pairSymbol string, orderSide binance.SideType, qty string, price string, optionalClientOrderID string) (*binance.CreateOrderResponse, error) {
	ret := _m.Called(ctx, pairSymbol, orderSide, qty, price, optionalClientOrderID)

	var r0 *binance.CreateOrderResponse
	if rf, ok := ret.Get(0).(func(context.Context, string, binance.SideType, string, string, string) *binance.CreateOrderResponse); ok {
		r0 = rf(ctx, pairSymbol, orderSide, qty, price, optionalClientOrderID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*binance.CreateOrderResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, binance.SideType, string, string, string) error); ok {
		r1 = rf(ctx, pairSymbol, orderSide, qty, price, optionalClientOrderID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockBinanceAPIWrapper_PlaceLimitOrder_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'PlaceLimitOrder'
type MockBinanceAPIWrapper_PlaceLimitOrder_Call struct {
	*mock.Call
}

// PlaceLimitOrder is a helper method to define mock.On call
//  - ctx context.Context
//  - pairSymbol string
//  - orderSide binance.SideType
//  - qty string
//  - price string
//  - optionalClientOrderID string
func (_e *MockBinanceAPIWrapper_Expecter) PlaceLimitOrder(ctx interface{}, pairSymbol interface{}, orderSide interface{}, qty interface{}, price interface{}, optionalClientOrderID interface{}) *MockBinanceAPIWrapper_PlaceLimitOrder_Call {
	return &MockBinanceAPIWrapper_PlaceLimitOrder_Call{Call: _e.mock.On("PlaceLimitOrder", ctx, pairSymbol, orderSide, qty, price, optionalClientOrderID)}
}

func (_c *MockBinanceAPIWrapper_PlaceLimitOrder_Call) Run(run func(ctx context.Context, pairSymbol string, orderSide binance.SideType, qty string, price string, optionalClientOrderID string)) *MockBinanceAPIWrapper_PlaceLimitOrder_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(binance.SideType), args[3].(string), args[4].(string), args[5].(string))
	})
	return _c
}

func (_c *MockBinanceAPIWrapper_PlaceLimitOrder_Call) Return(_a0 *binance.CreateOrderResponse, _a1 error) *MockBinanceAPIWrapper_PlaceLimitOrder_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// SubscribeToCandle provides a mock function with given fields: pairSymbol, interval, eventCallback, errorHandler
func (_m *MockBinanceAPIWrapper) SubscribeToCandle(pairSymbol string, interval string, eventCallback func(workers.CandleEvent), errorHandler func(error)) (chan struct{}, chan struct{}, error) {
	ret := _m.Called(pairSymbol, interval, eventCallback, errorHandler)

	var r0 chan struct{}
	if rf, ok := ret.Get(0).(func(string, string, func(workers.CandleEvent), func(error)) chan struct{}); ok {
		r0 = rf(pairSymbol, interval, eventCallback, errorHandler)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(chan struct{})
		}
	}

	var r1 chan struct{}
	if rf, ok := ret.Get(1).(func(string, string, func(workers.CandleEvent), func(error)) chan struct{}); ok {
		r1 = rf(pairSymbol, interval, eventCallback, errorHandler)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(chan struct{})
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(string, string, func(workers.CandleEvent), func(error)) error); ok {
		r2 = rf(pairSymbol, interval, eventCallback, errorHandler)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// MockBinanceAPIWrapper_SubscribeToCandle_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SubscribeToCandle'
type MockBinanceAPIWrapper_SubscribeToCandle_Call struct {
	*mock.Call
}

// SubscribeToCandle is a helper method to define mock.On call
//  - pairSymbol string
//  - interval string
//  - eventCallback func(workers.CandleEvent)
//  - errorHandler func(error)
func (_e *MockBinanceAPIWrapper_Expecter) SubscribeToCandle(pairSymbol interface{}, interval interface{}, eventCallback interface{}, errorHandler interface{}) *MockBinanceAPIWrapper_SubscribeToCandle_Call {
	return &MockBinanceAPIWrapper_SubscribeToCandle_Call{Call: _e.mock.On("SubscribeToCandle", pairSymbol, interval, eventCallback, errorHandler)}
}

func (_c *MockBinanceAPIWrapper_SubscribeToCandle_Call) Run(run func(pairSymbol string, interval string, eventCallback func(workers.CandleEvent), errorHandler func(error))) *MockBinanceAPIWrapper_SubscribeToCandle_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(string), args[2].(func(workers.CandleEvent)), args[3].(func(error)))
	})
	return _c
}

func (_c *MockBinanceAPIWrapper_SubscribeToCandle_Call) Return(doneC chan struct{}, stopC chan struct{}, err error) *MockBinanceAPIWrapper_SubscribeToCandle_Call {
	_c.Call.Return(doneC, stopC, err)
	return _c
}

// SubscribeToCandlesList provides a mock function with given fields: intervalsPerPair, eventCallback, errorHandler
func (_m *MockBinanceAPIWrapper) SubscribeToCandlesList(intervalsPerPair map[string]string, eventCallback func(workers.CandleEvent), errorHandler func(error)) (chan struct{}, chan struct{}, error) {
	ret := _m.Called(intervalsPerPair, eventCallback, errorHandler)

	var r0 chan struct{}
	if rf, ok := ret.Get(0).(func(map[string]string, func(workers.CandleEvent), func(error)) chan struct{}); ok {
		r0 = rf(intervalsPerPair, eventCallback, errorHandler)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(chan struct{})
		}
	}

	var r1 chan struct{}
	if rf, ok := ret.Get(1).(func(map[string]string, func(workers.CandleEvent), func(error)) chan struct{}); ok {
		r1 = rf(intervalsPerPair, eventCallback, errorHandler)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(chan struct{})
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(map[string]string, func(workers.CandleEvent), func(error)) error); ok {
		r2 = rf(intervalsPerPair, eventCallback, errorHandler)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// MockBinanceAPIWrapper_SubscribeToCandlesList_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SubscribeToCandlesList'
type MockBinanceAPIWrapper_SubscribeToCandlesList_Call struct {
	*mock.Call
}

// SubscribeToCandlesList is a helper method to define mock.On call
//  - intervalsPerPair map[string]string
//  - eventCallback func(workers.CandleEvent)
//  - errorHandler func(error)
func (_e *MockBinanceAPIWrapper_Expecter) SubscribeToCandlesList(intervalsPerPair interface{}, eventCallback interface{}, errorHandler interface{}) *MockBinanceAPIWrapper_SubscribeToCandlesList_Call {
	return &MockBinanceAPIWrapper_SubscribeToCandlesList_Call{Call: _e.mock.On("SubscribeToCandlesList", intervalsPerPair, eventCallback, errorHandler)}
}

func (_c *MockBinanceAPIWrapper_SubscribeToCandlesList_Call) Run(run func(intervalsPerPair map[string]string, eventCallback func(workers.CandleEvent), errorHandler func(error))) *MockBinanceAPIWrapper_SubscribeToCandlesList_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(map[string]string), args[1].(func(workers.CandleEvent)), args[2].(func(error)))
	})
	return _c
}

func (_c *MockBinanceAPIWrapper_SubscribeToCandlesList_Call) Return(doneC chan struct{}, stopC chan struct{}, err error) *MockBinanceAPIWrapper_SubscribeToCandlesList_Call {
	_c.Call.Return(doneC, stopC, err)
	return _c
}

// SubscribeToPriceEvents provides a mock function with given fields: pairSymbol, eventCallback, errorHandler
func (_m *MockBinanceAPIWrapper) SubscribeToPriceEvents(pairSymbol string, eventCallback binance.WsBookTickerHandler, errorHandler func(error)) (chan struct{}, chan struct{}, error) {
	ret := _m.Called(pairSymbol, eventCallback, errorHandler)

	var r0 chan struct{}
	if rf, ok := ret.Get(0).(func(string, binance.WsBookTickerHandler, func(error)) chan struct{}); ok {
		r0 = rf(pairSymbol, eventCallback, errorHandler)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(chan struct{})
		}
	}

	var r1 chan struct{}
	if rf, ok := ret.Get(1).(func(string, binance.WsBookTickerHandler, func(error)) chan struct{}); ok {
		r1 = rf(pairSymbol, eventCallback, errorHandler)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(chan struct{})
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(string, binance.WsBookTickerHandler, func(error)) error); ok {
		r2 = rf(pairSymbol, eventCallback, errorHandler)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// MockBinanceAPIWrapper_SubscribeToPriceEvents_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SubscribeToPriceEvents'
type MockBinanceAPIWrapper_SubscribeToPriceEvents_Call struct {
	*mock.Call
}

// SubscribeToPriceEvents is a helper method to define mock.On call
//  - pairSymbol string
//  - eventCallback binance.WsBookTickerHandler
//  - errorHandler func(error)
func (_e *MockBinanceAPIWrapper_Expecter) SubscribeToPriceEvents(pairSymbol interface{}, eventCallback interface{}, errorHandler interface{}) *MockBinanceAPIWrapper_SubscribeToPriceEvents_Call {
	return &MockBinanceAPIWrapper_SubscribeToPriceEvents_Call{Call: _e.mock.On("SubscribeToPriceEvents", pairSymbol, eventCallback, errorHandler)}
}

func (_c *MockBinanceAPIWrapper_SubscribeToPriceEvents_Call) Run(run func(pairSymbol string, eventCallback binance.WsBookTickerHandler, errorHandler func(error))) *MockBinanceAPIWrapper_SubscribeToPriceEvents_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(binance.WsBookTickerHandler), args[2].(func(error)))
	})
	return _c
}

func (_c *MockBinanceAPIWrapper_SubscribeToPriceEvents_Call) Return(doneC chan struct{}, stopC chan struct{}, err error) *MockBinanceAPIWrapper_SubscribeToPriceEvents_Call {
	_c.Call.Return(doneC, stopC, err)
	return _c
}

// SubscribeToTradeEvents provides a mock function with given fields: pairSymbol, exchangeTag, eventCallback, errorHandler
func (_m *MockBinanceAPIWrapper) SubscribeToTradeEvents(pairSymbol string, exchangeTag string, eventCallback func(workers.TradeEvent), errorHandler func(error)) (chan struct{}, chan struct{}, error) {
	ret := _m.Called(pairSymbol, exchangeTag, eventCallback, errorHandler)

	var r0 chan struct{}
	if rf, ok := ret.Get(0).(func(string, string, func(workers.TradeEvent), func(error)) chan struct{}); ok {
		r0 = rf(pairSymbol, exchangeTag, eventCallback, errorHandler)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(chan struct{})
		}
	}

	var r1 chan struct{}
	if rf, ok := ret.Get(1).(func(string, string, func(workers.TradeEvent), func(error)) chan struct{}); ok {
		r1 = rf(pairSymbol, exchangeTag, eventCallback, errorHandler)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(chan struct{})
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(string, string, func(workers.TradeEvent), func(error)) error); ok {
		r2 = rf(pairSymbol, exchangeTag, eventCallback, errorHandler)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// MockBinanceAPIWrapper_SubscribeToTradeEvents_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SubscribeToTradeEvents'
type MockBinanceAPIWrapper_SubscribeToTradeEvents_Call struct {
	*mock.Call
}

// SubscribeToTradeEvents is a helper method to define mock.On call
//  - pairSymbol string
//  - exchangeTag string
//  - eventCallback func(workers.TradeEvent)
//  - errorHandler func(error)
func (_e *MockBinanceAPIWrapper_Expecter) SubscribeToTradeEvents(pairSymbol interface{}, exchangeTag interface{}, eventCallback interface{}, errorHandler interface{}) *MockBinanceAPIWrapper_SubscribeToTradeEvents_Call {
	return &MockBinanceAPIWrapper_SubscribeToTradeEvents_Call{Call: _e.mock.On("SubscribeToTradeEvents", pairSymbol, exchangeTag, eventCallback, errorHandler)}
}

func (_c *MockBinanceAPIWrapper_SubscribeToTradeEvents_Call) Run(run func(pairSymbol string, exchangeTag string, eventCallback func(workers.TradeEvent), errorHandler func(error))) *MockBinanceAPIWrapper_SubscribeToTradeEvents_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(string), args[2].(func(workers.TradeEvent)), args[3].(func(error)))
	})
	return _c
}

func (_c *MockBinanceAPIWrapper_SubscribeToTradeEvents_Call) Return(doneC chan struct{}, stopC chan struct{}, err error) *MockBinanceAPIWrapper_SubscribeToTradeEvents_Call {
	_c.Call.Return(doneC, stopC, err)
	return _c
}

// SubscribeToTradeEventsPrivate provides a mock function with given fields: exchangeTag, callback, handler
func (_m *MockBinanceAPIWrapper) SubscribeToTradeEventsPrivate(exchangeTag string, callback workers.TradeEventPrivateCallback, handler func(error)) (chan struct{}, chan struct{}, error) {
	ret := _m.Called(exchangeTag, callback, handler)

	var r0 chan struct{}
	if rf, ok := ret.Get(0).(func(string, workers.TradeEventPrivateCallback, func(error)) chan struct{}); ok {
		r0 = rf(exchangeTag, callback, handler)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(chan struct{})
		}
	}

	var r1 chan struct{}
	if rf, ok := ret.Get(1).(func(string, workers.TradeEventPrivateCallback, func(error)) chan struct{}); ok {
		r1 = rf(exchangeTag, callback, handler)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(chan struct{})
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(string, workers.TradeEventPrivateCallback, func(error)) error); ok {
		r2 = rf(exchangeTag, callback, handler)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// MockBinanceAPIWrapper_SubscribeToTradeEventsPrivate_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SubscribeToTradeEventsPrivate'
type MockBinanceAPIWrapper_SubscribeToTradeEventsPrivate_Call struct {
	*mock.Call
}

// SubscribeToTradeEventsPrivate is a helper method to define mock.On call
//  - exchangeTag string
//  - callback workers.TradeEventPrivateCallback
//  - handler func(error)
func (_e *MockBinanceAPIWrapper_Expecter) SubscribeToTradeEventsPrivate(exchangeTag interface{}, callback interface{}, handler interface{}) *MockBinanceAPIWrapper_SubscribeToTradeEventsPrivate_Call {
	return &MockBinanceAPIWrapper_SubscribeToTradeEventsPrivate_Call{Call: _e.mock.On("SubscribeToTradeEventsPrivate", exchangeTag, callback, handler)}
}

func (_c *MockBinanceAPIWrapper_SubscribeToTradeEventsPrivate_Call) Run(run func(exchangeTag string, callback workers.TradeEventPrivateCallback, handler func(error))) *MockBinanceAPIWrapper_SubscribeToTradeEventsPrivate_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(workers.TradeEventPrivateCallback), args[2].(func(error)))
	})
	return _c
}

func (_c *MockBinanceAPIWrapper_SubscribeToTradeEventsPrivate_Call) Return(doneC chan struct{}, stopC chan struct{}, err error) *MockBinanceAPIWrapper_SubscribeToTradeEventsPrivate_Call {
	_c.Call.Return(doneC, stopC, err)
	return _c
}

// Sync provides a mock function with given fields: _a0
func (_m *MockBinanceAPIWrapper) Sync(_a0 context.Context) {
	_m.Called(_a0)
}

// MockBinanceAPIWrapper_Sync_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Sync'
type MockBinanceAPIWrapper_Sync_Call struct {
	*mock.Call
}

// Sync is a helper method to define mock.On call
//  - _a0 context.Context
func (_e *MockBinanceAPIWrapper_Expecter) Sync(_a0 interface{}) *MockBinanceAPIWrapper_Sync_Call {
	return &MockBinanceAPIWrapper_Sync_Call{Call: _e.mock.On("Sync", _a0)}
}

func (_c *MockBinanceAPIWrapper_Sync_Call) Run(run func(_a0 context.Context)) *MockBinanceAPIWrapper_Sync_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *MockBinanceAPIWrapper_Sync_Call) Return() *MockBinanceAPIWrapper_Sync_Call {
	_c.Call.Return()
	return _c
}

type mockConstructorTestingTNewMockBinanceAPIWrapper interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockBinanceAPIWrapper creates a new instance of MockBinanceAPIWrapper. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockBinanceAPIWrapper(t mockConstructorTestingTNewMockBinanceAPIWrapper) *MockBinanceAPIWrapper {
	mock := &MockBinanceAPIWrapper{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
