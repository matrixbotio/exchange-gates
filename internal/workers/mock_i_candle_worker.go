// Code generated by mockery v2.32.4. DO NOT EDIT.

package workers

import mock "github.com/stretchr/testify/mock"

// MockICandleWorker is an autogenerated mock type for the ICandleWorker type
type MockICandleWorker struct {
	mock.Mock
}

type MockICandleWorker_Expecter struct {
	mock *mock.Mock
}

func (_m *MockICandleWorker) EXPECT() *MockICandleWorker_Expecter {
	return &MockICandleWorker_Expecter{mock: &_m.Mock}
}

// GetExchangeTag provides a mock function with given fields:
func (_m *MockICandleWorker) GetExchangeTag() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// MockICandleWorker_GetExchangeTag_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetExchangeTag'
type MockICandleWorker_GetExchangeTag_Call struct {
	*mock.Call
}

// GetExchangeTag is a helper method to define mock.On call
func (_e *MockICandleWorker_Expecter) GetExchangeTag() *MockICandleWorker_GetExchangeTag_Call {
	return &MockICandleWorker_GetExchangeTag_Call{Call: _e.mock.On("GetExchangeTag")}
}

func (_c *MockICandleWorker_GetExchangeTag_Call) Run(run func()) *MockICandleWorker_GetExchangeTag_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockICandleWorker_GetExchangeTag_Call) Return(_a0 string) *MockICandleWorker_GetExchangeTag_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockICandleWorker_GetExchangeTag_Call) RunAndReturn(run func() string) *MockICandleWorker_GetExchangeTag_Call {
	_c.Call.Return(run)
	return _c
}

// Stop provides a mock function with given fields:
func (_m *MockICandleWorker) Stop() {
	_m.Called()
}

// MockICandleWorker_Stop_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Stop'
type MockICandleWorker_Stop_Call struct {
	*mock.Call
}

// Stop is a helper method to define mock.On call
func (_e *MockICandleWorker_Expecter) Stop() *MockICandleWorker_Stop_Call {
	return &MockICandleWorker_Stop_Call{Call: _e.mock.On("Stop")}
}

func (_c *MockICandleWorker_Stop_Call) Run(run func()) *MockICandleWorker_Stop_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockICandleWorker_Stop_Call) Return() *MockICandleWorker_Stop_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockICandleWorker_Stop_Call) RunAndReturn(run func()) *MockICandleWorker_Stop_Call {
	_c.Call.Return(run)
	return _c
}

// SubscribeToCandle provides a mock function with given fields: pairSymbol, eventCallback, errorHandler
func (_m *MockICandleWorker) SubscribeToCandle(pairSymbol string, eventCallback func(CandleEvent), errorHandler func(error)) error {
	ret := _m.Called(pairSymbol, eventCallback, errorHandler)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, func(CandleEvent), func(error)) error); ok {
		r0 = rf(pairSymbol, eventCallback, errorHandler)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockICandleWorker_SubscribeToCandle_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SubscribeToCandle'
type MockICandleWorker_SubscribeToCandle_Call struct {
	*mock.Call
}

// SubscribeToCandle is a helper method to define mock.On call
//   - pairSymbol string
//   - eventCallback func(CandleEvent)
//   - errorHandler func(error)
func (_e *MockICandleWorker_Expecter) SubscribeToCandle(pairSymbol interface{}, eventCallback interface{}, errorHandler interface{}) *MockICandleWorker_SubscribeToCandle_Call {
	return &MockICandleWorker_SubscribeToCandle_Call{Call: _e.mock.On("SubscribeToCandle", pairSymbol, eventCallback, errorHandler)}
}

func (_c *MockICandleWorker_SubscribeToCandle_Call) Run(run func(pairSymbol string, eventCallback func(CandleEvent), errorHandler func(error))) *MockICandleWorker_SubscribeToCandle_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(func(CandleEvent)), args[2].(func(error)))
	})
	return _c
}

func (_c *MockICandleWorker_SubscribeToCandle_Call) Return(_a0 error) *MockICandleWorker_SubscribeToCandle_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockICandleWorker_SubscribeToCandle_Call) RunAndReturn(run func(string, func(CandleEvent), func(error)) error) *MockICandleWorker_SubscribeToCandle_Call {
	_c.Call.Return(run)
	return _c
}

// SubscribeToCandlesList provides a mock function with given fields: intervalsPerPair, eventCallback, errorHandler
func (_m *MockICandleWorker) SubscribeToCandlesList(intervalsPerPair map[string]string, eventCallback func(CandleEvent), errorHandler func(error)) error {
	ret := _m.Called(intervalsPerPair, eventCallback, errorHandler)

	var r0 error
	if rf, ok := ret.Get(0).(func(map[string]string, func(CandleEvent), func(error)) error); ok {
		r0 = rf(intervalsPerPair, eventCallback, errorHandler)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockICandleWorker_SubscribeToCandlesList_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SubscribeToCandlesList'
type MockICandleWorker_SubscribeToCandlesList_Call struct {
	*mock.Call
}

// SubscribeToCandlesList is a helper method to define mock.On call
//   - intervalsPerPair map[string]string
//   - eventCallback func(CandleEvent)
//   - errorHandler func(error)
func (_e *MockICandleWorker_Expecter) SubscribeToCandlesList(intervalsPerPair interface{}, eventCallback interface{}, errorHandler interface{}) *MockICandleWorker_SubscribeToCandlesList_Call {
	return &MockICandleWorker_SubscribeToCandlesList_Call{Call: _e.mock.On("SubscribeToCandlesList", intervalsPerPair, eventCallback, errorHandler)}
}

func (_c *MockICandleWorker_SubscribeToCandlesList_Call) Run(run func(intervalsPerPair map[string]string, eventCallback func(CandleEvent), errorHandler func(error))) *MockICandleWorker_SubscribeToCandlesList_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(map[string]string), args[1].(func(CandleEvent)), args[2].(func(error)))
	})
	return _c
}

func (_c *MockICandleWorker_SubscribeToCandlesList_Call) Return(_a0 error) *MockICandleWorker_SubscribeToCandlesList_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockICandleWorker_SubscribeToCandlesList_Call) RunAndReturn(run func(map[string]string, func(CandleEvent), func(error)) error) *MockICandleWorker_SubscribeToCandlesList_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockICandleWorker creates a new instance of MockICandleWorker. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockICandleWorker(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockICandleWorker {
	mock := &MockICandleWorker{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
