// Code generated by MockGen. DO NOT EDIT.
// Source: gathering_structs.go

// Package gathering is a generated GoMock package.
package gathering

import (
	repository "anomaly_detector/internal/repository"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockICollector is a mock of ICollector interface
type MockICollector struct {
	ctrl     *gomock.Controller
	recorder *MockICollectorMockRecorder
}

// MockICollectorMockRecorder is the mock recorder for MockICollector
type MockICollectorMockRecorder struct {
	mock *MockICollector
}

// NewMockICollector creates a new mock instance
func NewMockICollector(ctrl *gomock.Controller) *MockICollector {
	mock := &MockICollector{ctrl: ctrl}
	mock.recorder = &MockICollectorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockICollector) EXPECT() *MockICollectorMockRecorder {
	return m.recorder
}

// HandleEvent mocks base method
func (m *MockICollector) HandleEvent(entityID int32, eventLabel string, eventValue int32) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "HandleEvent", entityID, eventLabel, eventValue)
}

// HandleEvent indicates an expected call of HandleEvent
func (mr *MockICollectorMockRecorder) HandleEvent(entityID, eventLabel, eventValue interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HandleEvent", reflect.TypeOf((*MockICollector)(nil).HandleEvent), entityID, eventLabel, eventValue)
}

// GetState mocks base method
func (m *MockICollector) GetState(data repository.StatRequestMessage) (repository.EventPreparedList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetState", data)
	ret0, _ := ret[0].(repository.EventPreparedList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetState indicates an expected call of GetState
func (mr *MockICollectorMockRecorder) GetState(data interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetState", reflect.TypeOf((*MockICollector)(nil).GetState), data)
}
