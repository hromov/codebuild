// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/hromov/notification/notification (interfaces: TopicSender)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockTopicSender is a mock of TopicSender interface.
type MockTopicSender struct {
	ctrl     *gomock.Controller
	recorder *MockTopicSenderMockRecorder
}

// MockTopicSenderMockRecorder is the mock recorder for MockTopicSender.
type MockTopicSenderMockRecorder struct {
	mock *MockTopicSender
}

// NewMockTopicSender creates a new mock instance.
func NewMockTopicSender(ctrl *gomock.Controller) *MockTopicSender {
	mock := &MockTopicSender{ctrl: ctrl}
	mock.recorder = &MockTopicSenderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTopicSender) EXPECT() *MockTopicSenderMockRecorder {
	return m.recorder
}

// Send mocks base method.
func (m *MockTopicSender) Send(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Send", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Send indicates an expected call of Send.
func (mr *MockTopicSenderMockRecorder) Send(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Send", reflect.TypeOf((*MockTopicSender)(nil).Send), arg0)
}
