// Code generated by mockery v2.34.2. DO NOT EDIT.

package initialstate

import (
	types "github.com/scaling-lightning/scaling-lightning/pkg/types"
	mock "github.com/stretchr/testify/mock"
)

// MockSLNetworkInterface is an autogenerated mock type for the SLNetworkInterface type
type MockSLNetworkInterface struct {
	mock.Mock
}

type MockSLNetworkInterface_Expecter struct {
	mock *mock.Mock
}

func (_m *MockSLNetworkInterface) EXPECT() *MockSLNetworkInterface_Expecter {
	return &MockSLNetworkInterface_Expecter{mock: &_m.Mock}
}

// ChannelBalance provides a mock function with given fields: nodeName
func (_m *MockSLNetworkInterface) ChannelBalance(nodeName string) (types.Amount, error) {
	ret := _m.Called(nodeName)

	var r0 types.Amount
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (types.Amount, error)); ok {
		return rf(nodeName)
	}
	if rf, ok := ret.Get(0).(func(string) types.Amount); ok {
		r0 = rf(nodeName)
	} else {
		r0 = ret.Get(0).(types.Amount)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(nodeName)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockSLNetworkInterface_ChannelBalance_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ChannelBalance'
type MockSLNetworkInterface_ChannelBalance_Call struct {
	*mock.Call
}

// ChannelBalance is a helper method to define mock.On call
//   - nodeName string
func (_e *MockSLNetworkInterface_Expecter) ChannelBalance(nodeName interface{}) *MockSLNetworkInterface_ChannelBalance_Call {
	return &MockSLNetworkInterface_ChannelBalance_Call{Call: _e.mock.On("ChannelBalance", nodeName)}
}

func (_c *MockSLNetworkInterface_ChannelBalance_Call) Run(run func(nodeName string)) *MockSLNetworkInterface_ChannelBalance_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *MockSLNetworkInterface_ChannelBalance_Call) Return(_a0 types.Amount, _a1 error) *MockSLNetworkInterface_ChannelBalance_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockSLNetworkInterface_ChannelBalance_Call) RunAndReturn(run func(string) (types.Amount, error)) *MockSLNetworkInterface_ChannelBalance_Call {
	_c.Call.Return(run)
	return _c
}

// ConnectPeer provides a mock function with given fields: fromNodeName, toNodeName
func (_m *MockSLNetworkInterface) ConnectPeer(fromNodeName string, toNodeName string) (string, error) {
	ret := _m.Called(fromNodeName, toNodeName)

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(string, string) (string, error)); ok {
		return rf(fromNodeName, toNodeName)
	}

	if rf, ok := ret.Get(0).(func(string, string) string); ok {
		r0 = rf(fromNodeName, toNodeName)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(fromNodeName, toNodeName)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockSLNetworkInterface_ConnectPeer_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ConnectPeer'
type MockSLNetworkInterface_ConnectPeer_Call struct {
	*mock.Call
}

// ConnectPeer is a helper method to define mock.On call
//   - fromNodeName string
//   - toNodeName string
func (_e *MockSLNetworkInterface_Expecter) ConnectPeer(fromNodeName interface{}, toNodeName interface{}) *MockSLNetworkInterface_ConnectPeer_Call {
	return &MockSLNetworkInterface_ConnectPeer_Call{Call: _e.mock.On("ConnectPeer", fromNodeName, toNodeName)}
}

func (_c *MockSLNetworkInterface_ConnectPeer_Call) Run(run func(fromNodeName string, toNodeName string)) *MockSLNetworkInterface_ConnectPeer_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(string))
	})
	return _c
}

func (_c *MockSLNetworkInterface_ConnectPeer_Call) Return(_a0 string, _a1 error) *MockSLNetworkInterface_ConnectPeer_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockSLNetworkInterface_ConnectPeer_Call) RunAndReturn(run func(string, string) (string, error)) *MockSLNetworkInterface_ConnectPeer_Call {
	_c.Call.Return(run)
	return _c
}

// CreateInvoice provides a mock function with given fields: nodeName, amountSats
func (_m *MockSLNetworkInterface) CreateInvoice(nodeName string, amountSats uint64) (string, error) {
	ret := _m.Called(nodeName, amountSats)

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(string, uint64) (string, error)); ok {
		return rf(nodeName, amountSats)
	}
	if rf, ok := ret.Get(0).(func(string, uint64) string); ok {
		r0 = rf(nodeName, amountSats)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(string, uint64) error); ok {
		r1 = rf(nodeName, amountSats)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockSLNetworkInterface_CreateInvoice_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateInvoice'
type MockSLNetworkInterface_CreateInvoice_Call struct {
	*mock.Call
}

// CreateInvoice is a helper method to define mock.On call
//   - nodeName string
//   - amountSats uint64
func (_e *MockSLNetworkInterface_Expecter) CreateInvoice(nodeName interface{}, amountSats interface{}) *MockSLNetworkInterface_CreateInvoice_Call {
	return &MockSLNetworkInterface_CreateInvoice_Call{Call: _e.mock.On("CreateInvoice", nodeName, amountSats)}
}

func (_c *MockSLNetworkInterface_CreateInvoice_Call) Run(run func(nodeName string, amountSats uint64)) *MockSLNetworkInterface_CreateInvoice_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(uint64))
	})
	return _c
}

func (_c *MockSLNetworkInterface_CreateInvoice_Call) Return(_a0 string, _a1 error) *MockSLNetworkInterface_CreateInvoice_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockSLNetworkInterface_CreateInvoice_Call) RunAndReturn(run func(string, uint64) (string, error)) *MockSLNetworkInterface_CreateInvoice_Call {
	_c.Call.Return(run)
	return _c
}

// OpenChannel provides a mock function with given fields: fromNodeName, toNodeName, localAmt
func (_m *MockSLNetworkInterface) OpenChannel(fromNodeName string, toNodeName string, localAmt uint64) (types.ChannelPoint, error) {
	ret := _m.Called(fromNodeName, toNodeName, localAmt)

	var r0 types.ChannelPoint
	var r1 error
	if rf, ok := ret.Get(0).(func(string, string, uint64) (types.ChannelPoint, error)); ok {
		return rf(fromNodeName, toNodeName, localAmt)
	}
	if rf, ok := ret.Get(0).(func(string, string, uint64) types.ChannelPoint); ok {
		r0 = rf(fromNodeName, toNodeName, localAmt)
	} else {
		r0 = ret.Get(0).(types.ChannelPoint)
	}

	if rf, ok := ret.Get(1).(func(string, string, uint64) error); ok {
		r1 = rf(fromNodeName, toNodeName, localAmt)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockSLNetworkInterface_OpenChannel_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'OpenChannel'
type MockSLNetworkInterface_OpenChannel_Call struct {
	*mock.Call
}

// OpenChannel is a helper method to define mock.On call
//   - fromNodeName string
//   - toNodeName string
//   - localAmt uint64
func (_e *MockSLNetworkInterface_Expecter) OpenChannel(fromNodeName interface{}, toNodeName interface{}, localAmt interface{}) *MockSLNetworkInterface_OpenChannel_Call {
	return &MockSLNetworkInterface_OpenChannel_Call{Call: _e.mock.On("OpenChannel", fromNodeName, toNodeName, localAmt)}
}

func (_c *MockSLNetworkInterface_OpenChannel_Call) Run(run func(fromNodeName string, toNodeName string, localAmt uint64)) *MockSLNetworkInterface_OpenChannel_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(string), args[2].(uint64))
	})
	return _c
}

func (_c *MockSLNetworkInterface_OpenChannel_Call) Return(_a0 types.ChannelPoint, _a1 error) *MockSLNetworkInterface_OpenChannel_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockSLNetworkInterface_OpenChannel_Call) RunAndReturn(run func(string, string, uint64) (types.ChannelPoint, error)) *MockSLNetworkInterface_OpenChannel_Call {
	_c.Call.Return(run)
	return _c
}

// PayInvoice provides a mock function with given fields: nodeName, invoice
func (_m *MockSLNetworkInterface) PayInvoice(nodeName string, invoice string) (string, error) {
	ret := _m.Called(nodeName, invoice)

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(string, string) (string, error)); ok {
		return rf(nodeName, invoice)
	}
	if rf, ok := ret.Get(0).(func(string, string) string); ok {
		r0 = rf(nodeName, invoice)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(nodeName, invoice)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockSLNetworkInterface_PayInvoice_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'PayInvoice'
type MockSLNetworkInterface_PayInvoice_Call struct {
	*mock.Call
}

// PayInvoice is a helper method to define mock.On call
//   - nodeName string
//   - invoice string
func (_e *MockSLNetworkInterface_Expecter) PayInvoice(nodeName interface{}, invoice interface{}) *MockSLNetworkInterface_PayInvoice_Call {
	return &MockSLNetworkInterface_PayInvoice_Call{Call: _e.mock.On("PayInvoice", nodeName, invoice)}
}

func (_c *MockSLNetworkInterface_PayInvoice_Call) Run(run func(nodeName string, invoice string)) *MockSLNetworkInterface_PayInvoice_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(string))
	})
	return _c
}

func (_c *MockSLNetworkInterface_PayInvoice_Call) Return(_a0 string, _a1 error) *MockSLNetworkInterface_PayInvoice_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockSLNetworkInterface_PayInvoice_Call) RunAndReturn(run func(string, string) (string, error)) *MockSLNetworkInterface_PayInvoice_Call {
	_c.Call.Return(run)
	return _c
}

// Send provides a mock function with given fields: fromNodeName, toNodeName, amountSats
func (_m *MockSLNetworkInterface) Send(fromNodeName string, toNodeName string, amountSats uint64) (string, error) {
	ret := _m.Called(fromNodeName, toNodeName, amountSats)

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(string, string, uint64) (string, error)); ok {
		return rf(fromNodeName, toNodeName, amountSats)
	}
	if rf, ok := ret.Get(0).(func(string, string, uint64) string); ok {
		r0 = rf(fromNodeName, toNodeName, amountSats)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(string, string, uint64) error); ok {
		r1 = rf(fromNodeName, toNodeName, amountSats)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockSLNetworkInterface_Send_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Send'
type MockSLNetworkInterface_Send_Call struct {
	*mock.Call
}

// Send is a helper method to define mock.On call
//   - fromNodeName string
//   - toNodeName string
//   - amountSats uint64
func (_e *MockSLNetworkInterface_Expecter) Send(fromNodeName interface{}, toNodeName interface{}, amountSats interface{}) *MockSLNetworkInterface_Send_Call {
	return &MockSLNetworkInterface_Send_Call{Call: _e.mock.On("Send", fromNodeName, toNodeName, amountSats)}
}

func (_c *MockSLNetworkInterface_Send_Call) Run(run func(fromNodeName string, toNodeName string, amountSats uint64)) *MockSLNetworkInterface_Send_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(string), args[2].(uint64))
	})
	return _c
}

func (_c *MockSLNetworkInterface_Send_Call) Return(_a0 string, _a1 error) *MockSLNetworkInterface_Send_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockSLNetworkInterface_Send_Call) RunAndReturn(run func(string, string, uint64) (string, error)) *MockSLNetworkInterface_Send_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockSLNetworkInterface creates a new instance of MockSLNetworkInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockSLNetworkInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockSLNetworkInterface {
	mock := &MockSLNetworkInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
