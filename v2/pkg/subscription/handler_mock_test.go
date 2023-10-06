// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/wundergraph/graphql-go-tools/v2/pkg/subscription (interfaces: Protocol,EventHandler)

// Package subscription is a generated GoMock package.
package subscription

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockProtocol is a mock of Protocol interface.
type MockProtocol struct {
	ctrl     *gomock.Controller
	recorder *MockProtocolMockRecorder
}

// MockProtocolMockRecorder is the mock recorder for MockProtocol.
type MockProtocolMockRecorder struct {
	mock *MockProtocol
}

// NewMockProtocol creates a new mock instance.
func NewMockProtocol(ctrl *gomock.Controller) *MockProtocol {
	mock := &MockProtocol{ctrl: ctrl}
	mock.recorder = &MockProtocolMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockProtocol) EXPECT() *MockProtocolMockRecorder {
	return m.recorder
}

// EventHandler mocks base method.
func (m *MockProtocol) EventHandler() EventHandler {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EventHandler")
	ret0, _ := ret[0].(EventHandler)
	return ret0
}

// EventHandler indicates an expected call of EventHandler.
func (mr *MockProtocolMockRecorder) EventHandler() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EventHandler", reflect.TypeOf((*MockProtocol)(nil).EventHandler))
}

// Handle mocks base method.
func (m *MockProtocol) Handle(arg0 context.Context, arg1 Engine, arg2 []byte) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Handle", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// Handle indicates an expected call of Handle.
func (mr *MockProtocolMockRecorder) Handle(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Handle", reflect.TypeOf((*MockProtocol)(nil).Handle), arg0, arg1, arg2)
}

// MockEventHandler is a mock of EventHandler interface.
type MockEventHandler struct {
	ctrl     *gomock.Controller
	recorder *MockEventHandlerMockRecorder
}

// MockEventHandlerMockRecorder is the mock recorder for MockEventHandler.
type MockEventHandlerMockRecorder struct {
	mock *MockEventHandler
}

// NewMockEventHandler creates a new mock instance.
func NewMockEventHandler(ctrl *gomock.Controller) *MockEventHandler {
	mock := &MockEventHandler{ctrl: ctrl}
	mock.recorder = &MockEventHandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockEventHandler) EXPECT() *MockEventHandlerMockRecorder {
	return m.recorder
}

// Emit mocks base method.
func (m *MockEventHandler) Emit(arg0 EventType, arg1 string, arg2 []byte, arg3 error) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Emit", arg0, arg1, arg2, arg3)
}

// Emit indicates an expected call of Emit.
func (mr *MockEventHandlerMockRecorder) Emit(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Emit", reflect.TypeOf((*MockEventHandler)(nil).Emit), arg0, arg1, arg2, arg3)
}
