// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/repository/repository.go

// Package mocks is a generated GoMock package.
package mocks

import (
	"github.com/kossadda/wallet-service/models"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockOperation is a mock of Operation interface.
type MockOperation struct {
	ctrl     *gomock.Controller
	recorder *MockOperationMockRecorder
}

// MockOperationMockRecorder is the mock recorder for MockOperation.
type MockOperationMockRecorder struct {
	mock *MockOperation
}

// NewMockOperation creates a new mock instance.
func NewMockOperation(ctrl *gomock.Controller) *MockOperation {
	mock := &MockOperation{ctrl: ctrl}
	mock.recorder = &MockOperationMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOperation) EXPECT() *MockOperationMockRecorder {
	return m.recorder
}

// BalanceChange mocks base method.
func (m *MockOperation) BalanceChange(req models.Request) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BalanceChange", req)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BalanceChange indicates an expected call of BalanceChange.
func (mr *MockOperationMockRecorder) BalanceChange(req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BalanceChange", reflect.TypeOf((*MockOperation)(nil).BalanceChange), req)
}

// BalanceCheck mocks base method.
func (m *MockOperation) BalanceCheck(id string) (float64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BalanceCheck", id)
	ret0, _ := ret[0].(float64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BalanceCheck indicates an expected call of BalanceCheck.
func (mr *MockOperationMockRecorder) BalanceCheck(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BalanceCheck", reflect.TypeOf((*MockOperation)(nil).BalanceCheck), id)
}
