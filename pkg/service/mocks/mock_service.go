// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/service/service.go

// Package mocks is a generated GoMock package.
package mocks

import (
	"github.com/kossadda/wallet-service/models"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockOperators is a mock of Operators interface.
type MockOperators struct {
	ctrl     *gomock.Controller
	recorder *MockOperatorsMockRecorder
}

// MockOperatorsMockRecorder is the mock recorder for MockOperators.
type MockOperatorsMockRecorder struct {
	mock *MockOperators
}

// NewMockOperators creates a new mock instance.
func NewMockOperators(ctrl *gomock.Controller) *MockOperators {
	mock := &MockOperators{ctrl: ctrl}
	mock.recorder = &MockOperatorsMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOperators) EXPECT() *MockOperatorsMockRecorder {
	return m.recorder
}

// BalanceChange mocks base method.
func (m *MockOperators) BalanceChange(req models.Request) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BalanceChange", req)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BalanceChange indicates an expected call of BalanceChange.
func (mr *MockOperatorsMockRecorder) BalanceChange(req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BalanceChange", reflect.TypeOf((*MockOperators)(nil).BalanceChange), req)
}

// BalanceCheck mocks base method.
func (m *MockOperators) BalanceCheck(id string) (float64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BalanceCheck", id)
	ret0, _ := ret[0].(float64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BalanceCheck indicates an expected call of BalanceCheck.
func (mr *MockOperatorsMockRecorder) BalanceCheck(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BalanceCheck", reflect.TypeOf((*MockOperators)(nil).BalanceCheck), id)
}
