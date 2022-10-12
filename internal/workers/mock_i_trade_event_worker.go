// Code generated by mockery v2.14.0. DO NOT EDIT.

package workers

import mock "github.com/stretchr/testify/mock"

// MockITradeEventWorker is an autogenerated mock type for the ITradeEventWorker type
type MockITradeEventWorker struct {
	mock.Mock
}

type MockITradeEventWorker_Expecter struct {
	mock *mock.Mock
}

func (_m *MockITradeEventWorker) EXPECT() *MockITradeEventWorker_Expecter {
	return &MockITradeEventWorker_Expecter{mock: &_m.Mock}
}

// GetExchangeTag provides a mock function with given fields:
func (_m *MockITradeEventWorker) GetExchangeTag() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// MockITradeEventWorker_GetExchangeTag_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetExchangeTag'
type MockITradeEventWorker_GetExchangeTag_Call struct {
	*mock.Call
}

// GetExchangeTag is a helper method to define mock.On call
func (_e *MockITradeEventWorker_Expecter) GetExchangeTag() *MockITradeEventWorker_GetExchangeTag_Call {
	return &MockITradeEventWorker_GetExchangeTag_Call{Call: _e.mock.On("GetExchangeTag")}
}

func (_c *MockITradeEventWorker_GetExchangeTag_Call) Run(run func()) *MockITradeEventWorker_GetExchangeTag_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockITradeEventWorker_GetExchangeTag_Call) Return(_a0 string) *MockITradeEventWorker_GetExchangeTag_Call {
	_c.Call.Return(_a0)
	return _c
}

// Stop provides a mock function with given fields:
func (_m *MockITradeEventWorker) Stop() {
	_m.Called()
}

// MockITradeEventWorker_Stop_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Stop'
type MockITradeEventWorker_Stop_Call struct {
	*mock.Call
}

// Stop is a helper method to define mock.On call
func (_e *MockITradeEventWorker_Expecter) Stop() *MockITradeEventWorker_Stop_Call {
	return &MockITradeEventWorker_Stop_Call{Call: _e.mock.On("Stop")}
}

func (_c *MockITradeEventWorker_Stop_Call) Run(run func()) *MockITradeEventWorker_Stop_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockITradeEventWorker_Stop_Call) Return() *MockITradeEventWorker_Stop_Call {
	_c.Call.Return()
	return _c
}

// SubscribeToTradeEvents provides a mock function with given fields: symbol, eventCallback, errorHandler
func (_m *MockITradeEventWorker) SubscribeToTradeEvents(symbol string, eventCallback func(TradeEvent), errorHandler func(error)) error {
	ret := _m.Called(symbol, eventCallback, errorHandler)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, func(TradeEvent), func(error)) error); ok {
		r0 = rf(symbol, eventCallback, errorHandler)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockITradeEventWorker_SubscribeToTradeEvents_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SubscribeToTradeEvents'
type MockITradeEventWorker_SubscribeToTradeEvents_Call struct {
	*mock.Call
}

// SubscribeToTradeEvents is a helper method to define mock.On call
//  - symbol string
//  - eventCallback func(TradeEvent)
//  - errorHandler func(error)
func (_e *MockITradeEventWorker_Expecter) SubscribeToTradeEvents(symbol interface{}, eventCallback interface{}, errorHandler interface{}) *MockITradeEventWorker_SubscribeToTradeEvents_Call {
	return &MockITradeEventWorker_SubscribeToTradeEvents_Call{Call: _e.mock.On("SubscribeToTradeEvents", symbol, eventCallback, errorHandler)}
}

func (_c *MockITradeEventWorker_SubscribeToTradeEvents_Call) Run(run func(symbol string, eventCallback func(TradeEvent), errorHandler func(error))) *MockITradeEventWorker_SubscribeToTradeEvents_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(func(TradeEvent)), args[2].(func(error)))
	})
	return _c
}

func (_c *MockITradeEventWorker_SubscribeToTradeEvents_Call) Return(_a0 error) *MockITradeEventWorker_SubscribeToTradeEvents_Call {
	_c.Call.Return(_a0)
	return _c
}

type mockConstructorTestingTNewMockITradeEventWorker interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockITradeEventWorker creates a new instance of MockITradeEventWorker. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockITradeEventWorker(t mockConstructorTestingTNewMockITradeEventWorker) *MockITradeEventWorker {
	mock := &MockITradeEventWorker{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
