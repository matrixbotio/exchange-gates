// Code generated by mockery v2.32.4. DO NOT EDIT.

package binance

import (
	context "context"

	v2 "github.com/adshao/go-binance/v2"
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
//   - ctx context.Context
//   - pairSymbol string
//   - clientOrderID string
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

func (_c *MockBinanceAPIWrapper_CancelOrderByClientOrderID_Call) RunAndReturn(run func(context.Context, string, string) error) *MockBinanceAPIWrapper_CancelOrderByClientOrderID_Call {
	_c.Call.Return(run)
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
//   - ctx context.Context
//   - pairSymbol string
//   - orderID int64
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

func (_c *MockBinanceAPIWrapper_CancelOrderByID_Call) RunAndReturn(run func(context.Context, string, int64) error) *MockBinanceAPIWrapper_CancelOrderByID_Call {
	_c.Call.Return(run)
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
//   - ctx context.Context
//   - keyPublic string
//   - keySecret string
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

func (_c *MockBinanceAPIWrapper_Connect_Call) RunAndReturn(run func(context.Context, string, string) error) *MockBinanceAPIWrapper_Connect_Call {
	_c.Call.Return(run)
	return _c
}

// GetAccountData provides a mock function with given fields: _a0
func (_m *MockBinanceAPIWrapper) GetAccountData(_a0 context.Context) (*v2.Account, error) {
	ret := _m.Called(_a0)

	var r0 *v2.Account
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (*v2.Account, error)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func(context.Context) *v2.Account); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v2.Account)
		}
	}

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
//   - _a0 context.Context
func (_e *MockBinanceAPIWrapper_Expecter) GetAccountData(_a0 interface{}) *MockBinanceAPIWrapper_GetAccountData_Call {
	return &MockBinanceAPIWrapper_GetAccountData_Call{Call: _e.mock.On("GetAccountData", _a0)}
}

func (_c *MockBinanceAPIWrapper_GetAccountData_Call) Run(run func(_a0 context.Context)) *MockBinanceAPIWrapper_GetAccountData_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *MockBinanceAPIWrapper_GetAccountData_Call) Return(_a0 *v2.Account, _a1 error) *MockBinanceAPIWrapper_GetAccountData_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockBinanceAPIWrapper_GetAccountData_Call) RunAndReturn(run func(context.Context) (*v2.Account, error)) *MockBinanceAPIWrapper_GetAccountData_Call {
	_c.Call.Return(run)
	return _c
}

// GetExchangeInfo provides a mock function with given fields: ctx, pairSymbol
func (_m *MockBinanceAPIWrapper) GetExchangeInfo(ctx context.Context, pairSymbol string) (*v2.ExchangeInfo, error) {
	ret := _m.Called(ctx, pairSymbol)

	var r0 *v2.ExchangeInfo
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*v2.ExchangeInfo, error)); ok {
		return rf(ctx, pairSymbol)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *v2.ExchangeInfo); ok {
		r0 = rf(ctx, pairSymbol)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v2.ExchangeInfo)
		}
	}

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
//   - ctx context.Context
//   - pairSymbol string
func (_e *MockBinanceAPIWrapper_Expecter) GetExchangeInfo(ctx interface{}, pairSymbol interface{}) *MockBinanceAPIWrapper_GetExchangeInfo_Call {
	return &MockBinanceAPIWrapper_GetExchangeInfo_Call{Call: _e.mock.On("GetExchangeInfo", ctx, pairSymbol)}
}

func (_c *MockBinanceAPIWrapper_GetExchangeInfo_Call) Run(run func(ctx context.Context, pairSymbol string)) *MockBinanceAPIWrapper_GetExchangeInfo_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *MockBinanceAPIWrapper_GetExchangeInfo_Call) Return(_a0 *v2.ExchangeInfo, _a1 error) *MockBinanceAPIWrapper_GetExchangeInfo_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockBinanceAPIWrapper_GetExchangeInfo_Call) RunAndReturn(run func(context.Context, string) (*v2.ExchangeInfo, error)) *MockBinanceAPIWrapper_GetExchangeInfo_Call {
	_c.Call.Return(run)
	return _c
}

// GetKlines provides a mock function with given fields: ctx, pairSymbol, interval, limit
func (_m *MockBinanceAPIWrapper) GetKlines(ctx context.Context, pairSymbol string, interval string, limit int) ([]*v2.Kline, error) {
	ret := _m.Called(ctx, pairSymbol, interval, limit)

	var r0 []*v2.Kline
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, int) ([]*v2.Kline, error)); ok {
		return rf(ctx, pairSymbol, interval, limit)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string, int) []*v2.Kline); ok {
		r0 = rf(ctx, pairSymbol, interval, limit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*v2.Kline)
		}
	}

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
//   - ctx context.Context
//   - pairSymbol string
//   - interval string
//   - limit int
func (_e *MockBinanceAPIWrapper_Expecter) GetKlines(ctx interface{}, pairSymbol interface{}, interval interface{}, limit interface{}) *MockBinanceAPIWrapper_GetKlines_Call {
	return &MockBinanceAPIWrapper_GetKlines_Call{Call: _e.mock.On("GetKlines", ctx, pairSymbol, interval, limit)}
}

func (_c *MockBinanceAPIWrapper_GetKlines_Call) Run(run func(ctx context.Context, pairSymbol string, interval string, limit int)) *MockBinanceAPIWrapper_GetKlines_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(string), args[3].(int))
	})
	return _c
}

func (_c *MockBinanceAPIWrapper_GetKlines_Call) Return(_a0 []*v2.Kline, _a1 error) *MockBinanceAPIWrapper_GetKlines_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockBinanceAPIWrapper_GetKlines_Call) RunAndReturn(run func(context.Context, string, string, int) ([]*v2.Kline, error)) *MockBinanceAPIWrapper_GetKlines_Call {
	_c.Call.Return(run)
	return _c
}

// GetOpenOrders provides a mock function with given fields: ctx, pairSymbol
func (_m *MockBinanceAPIWrapper) GetOpenOrders(ctx context.Context, pairSymbol string) ([]*v2.Order, error) {
	ret := _m.Called(ctx, pairSymbol)

	var r0 []*v2.Order
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) ([]*v2.Order, error)); ok {
		return rf(ctx, pairSymbol)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) []*v2.Order); ok {
		r0 = rf(ctx, pairSymbol)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*v2.Order)
		}
	}

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
//   - ctx context.Context
//   - pairSymbol string
func (_e *MockBinanceAPIWrapper_Expecter) GetOpenOrders(ctx interface{}, pairSymbol interface{}) *MockBinanceAPIWrapper_GetOpenOrders_Call {
	return &MockBinanceAPIWrapper_GetOpenOrders_Call{Call: _e.mock.On("GetOpenOrders", ctx, pairSymbol)}
}

func (_c *MockBinanceAPIWrapper_GetOpenOrders_Call) Run(run func(ctx context.Context, pairSymbol string)) *MockBinanceAPIWrapper_GetOpenOrders_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *MockBinanceAPIWrapper_GetOpenOrders_Call) Return(_a0 []*v2.Order, _a1 error) *MockBinanceAPIWrapper_GetOpenOrders_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockBinanceAPIWrapper_GetOpenOrders_Call) RunAndReturn(run func(context.Context, string) ([]*v2.Order, error)) *MockBinanceAPIWrapper_GetOpenOrders_Call {
	_c.Call.Return(run)
	return _c
}

// GetOrderDataByClientOrderID provides a mock function with given fields: ctx, pairSymbol, clientOrderID
func (_m *MockBinanceAPIWrapper) GetOrderDataByClientOrderID(ctx context.Context, pairSymbol string, clientOrderID string) (*v2.Order, error) {
	ret := _m.Called(ctx, pairSymbol, clientOrderID)

	var r0 *v2.Order
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) (*v2.Order, error)); ok {
		return rf(ctx, pairSymbol, clientOrderID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string) *v2.Order); ok {
		r0 = rf(ctx, pairSymbol, clientOrderID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v2.Order)
		}
	}

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
//   - ctx context.Context
//   - pairSymbol string
//   - clientOrderID string
func (_e *MockBinanceAPIWrapper_Expecter) GetOrderDataByClientOrderID(ctx interface{}, pairSymbol interface{}, clientOrderID interface{}) *MockBinanceAPIWrapper_GetOrderDataByClientOrderID_Call {
	return &MockBinanceAPIWrapper_GetOrderDataByClientOrderID_Call{Call: _e.mock.On("GetOrderDataByClientOrderID", ctx, pairSymbol, clientOrderID)}
}

func (_c *MockBinanceAPIWrapper_GetOrderDataByClientOrderID_Call) Run(run func(ctx context.Context, pairSymbol string, clientOrderID string)) *MockBinanceAPIWrapper_GetOrderDataByClientOrderID_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(string))
	})
	return _c
}

func (_c *MockBinanceAPIWrapper_GetOrderDataByClientOrderID_Call) Return(_a0 *v2.Order, _a1 error) *MockBinanceAPIWrapper_GetOrderDataByClientOrderID_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockBinanceAPIWrapper_GetOrderDataByClientOrderID_Call) RunAndReturn(run func(context.Context, string, string) (*v2.Order, error)) *MockBinanceAPIWrapper_GetOrderDataByClientOrderID_Call {
	_c.Call.Return(run)
	return _c
}

// GetOrderDataByOrderID provides a mock function with given fields: ctx, pairSymbol, orderID
func (_m *MockBinanceAPIWrapper) GetOrderDataByOrderID(ctx context.Context, pairSymbol string, orderID int64) (*v2.Order, error) {
	ret := _m.Called(ctx, pairSymbol, orderID)

	var r0 *v2.Order
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, int64) (*v2.Order, error)); ok {
		return rf(ctx, pairSymbol, orderID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, int64) *v2.Order); ok {
		r0 = rf(ctx, pairSymbol, orderID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v2.Order)
		}
	}

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
//   - ctx context.Context
//   - pairSymbol string
//   - orderID int64
func (_e *MockBinanceAPIWrapper_Expecter) GetOrderDataByOrderID(ctx interface{}, pairSymbol interface{}, orderID interface{}) *MockBinanceAPIWrapper_GetOrderDataByOrderID_Call {
	return &MockBinanceAPIWrapper_GetOrderDataByOrderID_Call{Call: _e.mock.On("GetOrderDataByOrderID", ctx, pairSymbol, orderID)}
}

func (_c *MockBinanceAPIWrapper_GetOrderDataByOrderID_Call) Run(run func(ctx context.Context, pairSymbol string, orderID int64)) *MockBinanceAPIWrapper_GetOrderDataByOrderID_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(int64))
	})
	return _c
}

func (_c *MockBinanceAPIWrapper_GetOrderDataByOrderID_Call) Return(_a0 *v2.Order, _a1 error) *MockBinanceAPIWrapper_GetOrderDataByOrderID_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockBinanceAPIWrapper_GetOrderDataByOrderID_Call) RunAndReturn(run func(context.Context, string, int64) (*v2.Order, error)) *MockBinanceAPIWrapper_GetOrderDataByOrderID_Call {
	_c.Call.Return(run)
	return _c
}

// GetPrices provides a mock function with given fields: ctx, pairSymbol
func (_m *MockBinanceAPIWrapper) GetPrices(ctx context.Context, pairSymbol string) ([]*v2.SymbolPrice, error) {
	ret := _m.Called(ctx, pairSymbol)

	var r0 []*v2.SymbolPrice
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) ([]*v2.SymbolPrice, error)); ok {
		return rf(ctx, pairSymbol)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) []*v2.SymbolPrice); ok {
		r0 = rf(ctx, pairSymbol)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*v2.SymbolPrice)
		}
	}

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
//   - ctx context.Context
//   - pairSymbol string
func (_e *MockBinanceAPIWrapper_Expecter) GetPrices(ctx interface{}, pairSymbol interface{}) *MockBinanceAPIWrapper_GetPrices_Call {
	return &MockBinanceAPIWrapper_GetPrices_Call{Call: _e.mock.On("GetPrices", ctx, pairSymbol)}
}

func (_c *MockBinanceAPIWrapper_GetPrices_Call) Run(run func(ctx context.Context, pairSymbol string)) *MockBinanceAPIWrapper_GetPrices_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *MockBinanceAPIWrapper_GetPrices_Call) Return(_a0 []*v2.SymbolPrice, _a1 error) *MockBinanceAPIWrapper_GetPrices_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockBinanceAPIWrapper_GetPrices_Call) RunAndReturn(run func(context.Context, string) ([]*v2.SymbolPrice, error)) *MockBinanceAPIWrapper_GetPrices_Call {
	_c.Call.Return(run)
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
//   - _a0 context.Context
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

func (_c *MockBinanceAPIWrapper_Ping_Call) RunAndReturn(run func(context.Context) error) *MockBinanceAPIWrapper_Ping_Call {
	_c.Call.Return(run)
	return _c
}

// PlaceLimitOrder provides a mock function with given fields: ctx, pairSymbol, orderSide, qty, price, optionalClientOrderID
func (_m *MockBinanceAPIWrapper) PlaceLimitOrder(ctx context.Context, pairSymbol string, orderSide v2.SideType, qty string, price string, optionalClientOrderID string) (*v2.CreateOrderResponse, error) {
	ret := _m.Called(ctx, pairSymbol, orderSide, qty, price, optionalClientOrderID)

	var r0 *v2.CreateOrderResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, v2.SideType, string, string, string) (*v2.CreateOrderResponse, error)); ok {
		return rf(ctx, pairSymbol, orderSide, qty, price, optionalClientOrderID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, v2.SideType, string, string, string) *v2.CreateOrderResponse); ok {
		r0 = rf(ctx, pairSymbol, orderSide, qty, price, optionalClientOrderID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v2.CreateOrderResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, v2.SideType, string, string, string) error); ok {
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
//   - ctx context.Context
//   - pairSymbol string
//   - orderSide v2.SideType
//   - qty string
//   - price string
//   - optionalClientOrderID string
func (_e *MockBinanceAPIWrapper_Expecter) PlaceLimitOrder(ctx interface{}, pairSymbol interface{}, orderSide interface{}, qty interface{}, price interface{}, optionalClientOrderID interface{}) *MockBinanceAPIWrapper_PlaceLimitOrder_Call {
	return &MockBinanceAPIWrapper_PlaceLimitOrder_Call{Call: _e.mock.On("PlaceLimitOrder", ctx, pairSymbol, orderSide, qty, price, optionalClientOrderID)}
}

func (_c *MockBinanceAPIWrapper_PlaceLimitOrder_Call) Run(run func(ctx context.Context, pairSymbol string, orderSide v2.SideType, qty string, price string, optionalClientOrderID string)) *MockBinanceAPIWrapper_PlaceLimitOrder_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(v2.SideType), args[3].(string), args[4].(string), args[5].(string))
	})
	return _c
}

func (_c *MockBinanceAPIWrapper_PlaceLimitOrder_Call) Return(_a0 *v2.CreateOrderResponse, _a1 error) *MockBinanceAPIWrapper_PlaceLimitOrder_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockBinanceAPIWrapper_PlaceLimitOrder_Call) RunAndReturn(run func(context.Context, string, v2.SideType, string, string, string) (*v2.CreateOrderResponse, error)) *MockBinanceAPIWrapper_PlaceLimitOrder_Call {
	_c.Call.Return(run)
	return _c
}

// SubscribeToCandle provides a mock function with given fields: pairSymbol, interval, eventCallback, errorHandler
func (_m *MockBinanceAPIWrapper) SubscribeToCandle(pairSymbol string, interval string, eventCallback func(workers.CandleEvent), errorHandler func(error)) (chan struct{}, chan struct{}, error) {
	ret := _m.Called(pairSymbol, interval, eventCallback, errorHandler)

	var r0 chan struct{}
	var r1 chan struct{}
	var r2 error
	if rf, ok := ret.Get(0).(func(string, string, func(workers.CandleEvent), func(error)) (chan struct{}, chan struct{}, error)); ok {
		return rf(pairSymbol, interval, eventCallback, errorHandler)
	}
	if rf, ok := ret.Get(0).(func(string, string, func(workers.CandleEvent), func(error)) chan struct{}); ok {
		r0 = rf(pairSymbol, interval, eventCallback, errorHandler)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(chan struct{})
		}
	}

	if rf, ok := ret.Get(1).(func(string, string, func(workers.CandleEvent), func(error)) chan struct{}); ok {
		r1 = rf(pairSymbol, interval, eventCallback, errorHandler)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(chan struct{})
		}
	}

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
//   - pairSymbol string
//   - interval string
//   - eventCallback func(workers.CandleEvent)
//   - errorHandler func(error)
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

func (_c *MockBinanceAPIWrapper_SubscribeToCandle_Call) RunAndReturn(run func(string, string, func(workers.CandleEvent), func(error)) (chan struct{}, chan struct{}, error)) *MockBinanceAPIWrapper_SubscribeToCandle_Call {
	_c.Call.Return(run)
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
//   - _a0 context.Context
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

func (_c *MockBinanceAPIWrapper_Sync_Call) RunAndReturn(run func(context.Context)) *MockBinanceAPIWrapper_Sync_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockBinanceAPIWrapper creates a new instance of MockBinanceAPIWrapper. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockBinanceAPIWrapper(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockBinanceAPIWrapper {
	mock := &MockBinanceAPIWrapper{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
