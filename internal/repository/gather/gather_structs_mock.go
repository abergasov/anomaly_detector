// Code generated by MockGen. DO NOT EDIT.
// Source: gather_structs.go

// Package gather is a generated GoMock package.
package gather

import (
	sql "database/sql"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockIStorageSaver is a mock of IStorageSaver interface
type MockIStorageSaver struct {
	ctrl     *gomock.Controller
	recorder *MockIStorageSaverMockRecorder
}

// MockIStorageSaverMockRecorder is the mock recorder for MockIStorageSaver
type MockIStorageSaverMockRecorder struct {
	mock *MockIStorageSaver
}

// NewMockIStorageSaver creates a new mock instance
func NewMockIStorageSaver(ctrl *gomock.Controller) *MockIStorageSaver {
	mock := &MockIStorageSaver{ctrl: ctrl}
	mock.recorder = &MockIStorageSaverMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockIStorageSaver) EXPECT() *MockIStorageSaverMockRecorder {
	return m.recorder
}

// Select mocks base method
func (m *MockIStorageSaver) Select(dest interface{}, query string, args ...interface{}) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{dest, query}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Select", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// Select indicates an expected call of Select
func (mr *MockIStorageSaverMockRecorder) Select(dest, query interface{}, args ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{dest, query}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Select", reflect.TypeOf((*MockIStorageSaver)(nil).Select), varargs...)
}

// Exec mocks base method
func (m *MockIStorageSaver) Exec(query string, args ...interface{}) (sql.Result, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{query}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Exec", varargs...)
	ret0, _ := ret[0].(sql.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Exec indicates an expected call of Exec
func (mr *MockIStorageSaverMockRecorder) Exec(query interface{}, args ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{query}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Exec", reflect.TypeOf((*MockIStorageSaver)(nil).Exec), varargs...)
}
