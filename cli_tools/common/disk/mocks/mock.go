// Code generated by MockGen. DO NOT EDIT.
// Source: inspect.go

// Package mock_disk is a generated GoMock package.
package mock_disk

import (
	disk "github.com/GoogleCloudPlatform/compute-image-tools/cli_tools/common/disk"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockInspector is a mock of Inspector interface
type MockInspector struct {
	ctrl     *gomock.Controller
	recorder *MockInspectorMockRecorder
}

// MockInspectorMockRecorder is the mock recorder for MockInspector
type MockInspectorMockRecorder struct {
	mock *MockInspector
}

// NewMockInspector creates a new mock instance
func NewMockInspector(ctrl *gomock.Controller) *MockInspector {
	mock := &MockInspector{ctrl: ctrl}
	mock.recorder = &MockInspectorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockInspector) EXPECT() *MockInspectorMockRecorder {
	return m.recorder
}

// Inspect mocks base method
func (m *MockInspector) Inspect(reference string, inspectOS bool) (disk.InspectionResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Inspect", reference, inspectOS)
	ret0, _ := ret[0].(disk.InspectionResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Inspect indicates an expected call of Inspect
func (mr *MockInspectorMockRecorder) Inspect(reference, inspectOS interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Inspect", reflect.TypeOf((*MockInspector)(nil).Inspect), reference, inspectOS)
}

// Cancel mocks base method
func (m *MockInspector) Cancel(reason string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Cancel", reason)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Cancel indicates an expected call of Cancel
func (mr *MockInspectorMockRecorder) Cancel(reason interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Cancel", reflect.TypeOf((*MockInspector)(nil).Cancel), reason)
}

// TraceLogs mocks base method
func (m *MockInspector) TraceLogs() []string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TraceLogs")
	ret0, _ := ret[0].([]string)
	return ret0
}

// TraceLogs indicates an expected call of TraceLogs
func (mr *MockInspectorMockRecorder) TraceLogs() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TraceLogs", reflect.TypeOf((*MockInspector)(nil).TraceLogs))
}

// MockdaisyWorker is a mock of daisyWorker interface
type MockdaisyWorker struct {
	ctrl     *gomock.Controller
	recorder *MockdaisyWorkerMockRecorder
}

// MockdaisyWorkerMockRecorder is the mock recorder for MockdaisyWorker
type MockdaisyWorkerMockRecorder struct {
	mock *MockdaisyWorker
}

// NewMockdaisyWorker creates a new mock instance
func NewMockdaisyWorker(ctrl *gomock.Controller) *MockdaisyWorker {
	mock := &MockdaisyWorker{ctrl: ctrl}
	mock.recorder = &MockdaisyWorkerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockdaisyWorker) EXPECT() *MockdaisyWorkerMockRecorder {
	return m.recorder
}

// runAndReadSerialValue mocks base method
func (m *MockdaisyWorker) runAndReadSerialValue(key string, vars map[string]string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "runAndReadSerialValue", key, vars)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// runAndReadSerialValue indicates an expected call of runAndReadSerialValue
func (mr *MockdaisyWorkerMockRecorder) runAndReadSerialValue(key, vars interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "runAndReadSerialValue", reflect.TypeOf((*MockdaisyWorker)(nil).runAndReadSerialValue), key, vars)
}

// cancel mocks base method
func (m *MockdaisyWorker) cancel(reason string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "cancel", reason)
	ret0, _ := ret[0].(bool)
	return ret0
}

// cancel indicates an expected call of cancel
func (mr *MockdaisyWorkerMockRecorder) cancel(reason interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "cancel", reflect.TypeOf((*MockdaisyWorker)(nil).cancel), reason)
}

// traceLogs mocks base method
func (m *MockdaisyWorker) traceLogs() []string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "traceLogs")
	ret0, _ := ret[0].([]string)
	return ret0
}

// traceLogs indicates an expected call of traceLogs
func (mr *MockdaisyWorkerMockRecorder) traceLogs() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "traceLogs", reflect.TypeOf((*MockdaisyWorker)(nil).traceLogs))
}