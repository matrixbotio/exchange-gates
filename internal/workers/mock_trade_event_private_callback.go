// Code generated by mockery v2.14.1. DO NOT EDIT.

package workers

import mock "github.com/stretchr/testify/mock"

// MockTradeEventPrivateCallback is an autogenerated mock type for the TradeEventPrivateCallback type
type MockTradeEventPrivateCallback struct {
	mock.Mock
}

type MockTradeEventPrivateCallback_Expecter struct {
	mock *mock.Mock
}

func (_m *MockTradeEventPrivateCallback) EXPECT() *MockTradeEventPrivateCallback_Expecter {
	return &MockTradeEventPrivateCallback_Expecter{mock: &_m.Mock}
}

// Execute provides a mock function with given fields: event
func (_m *MockTradeEventPrivateCallback) Execute(event TradeEventPrivate) {
	_m.Called(event)
}

// MockTradeEventPrivateCallback_Execute_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Execute'
type MockTradeEventPrivateCallback_Execute_Call struct {
	*mock.Call
}

// Execute is a helper method to define mock.On call
//  - event TradeEventPrivate
func (_e *MockTradeEventPrivateCallback_Expecter) Execute(event interface{}) *MockTradeEventPrivateCallback_Execute_Call {
	return &MockTradeEventPrivateCallback_Execute_Call{Call: _e.mock.On("Execute", event)}
}

func (_c *MockTradeEventPrivateCallback_Execute_Call) Run(run func(event TradeEventPrivate)) *MockTradeEventPrivateCallback_Execute_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(TradeEventPrivate))
	})
	return _c
}

func (_c *MockTradeEventPrivateCallback_Execute_Call) Return() *MockTradeEventPrivateCallback_Execute_Call {
	_c.Call.Return()
	return _c
}

type mockConstructorTestingTNewMockTradeEventPrivateCallback interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockTradeEventPrivateCallback creates a new instance of MockTradeEventPrivateCallback. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockTradeEventPrivateCallback(t mockConstructorTestingTNewMockTradeEventPrivateCallback) *MockTradeEventPrivateCallback {
	mock := &MockTradeEventPrivateCallback{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
