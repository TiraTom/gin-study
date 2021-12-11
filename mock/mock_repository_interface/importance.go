// Code generated by MockGen. DO NOT EDIT.
// Source: importance.go

// Package mock_repository_interface is a generated GoMock package.
package mock_repository_interface

import (
	reflect "reflect"

	domain_obj "github.com/Tiratom/gin-study/domain/domain_obj"
	gomock "github.com/golang/mock/gomock"
)

// MockImportance is a mock of Importance interface.
type MockImportance struct {
	ctrl     *gomock.Controller
	recorder *MockImportanceMockRecorder
}

// MockImportanceMockRecorder is the mock recorder for MockImportance.
type MockImportanceMockRecorder struct {
	mock *MockImportance
}

// NewMockImportance creates a new mock instance.
func NewMockImportance(ctrl *gomock.Controller) *MockImportance {
	mock := &MockImportance{ctrl: ctrl}
	mock.recorder = &MockImportanceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockImportance) EXPECT() *MockImportanceMockRecorder {
	return m.recorder
}

// GetAll mocks base method.
func (m *MockImportance) GetAll() ([]*domain_obj.Importance, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAll")
	ret0, _ := ret[0].([]*domain_obj.Importance)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAll indicates an expected call of GetAll.
func (mr *MockImportanceMockRecorder) GetAll() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockImportance)(nil).GetAll))
}