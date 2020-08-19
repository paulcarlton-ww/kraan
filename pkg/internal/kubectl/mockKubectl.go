// Code generated by MockGen. DO NOT EDIT.
// Source: kubectl.go

// Package kubectl is a generated GoMock package.
package kubectl

import (
	reflect "reflect"

	logr "github.com/go-logr/logr"
	gomock "github.com/golang/mock/gomock"
)

// MockKubectl is a mock of Kubectl interface
type MockKubectl struct {
	ctrl     *gomock.Controller
	recorder *MockKubectlMockRecorder
}

// MockKubectlMockRecorder is the mock recorder for MockKubectl
type MockKubectlMockRecorder struct {
	mock *MockKubectl
}

// NewMockKubectl creates a new mock instance
func NewMockKubectl(ctrl *gomock.Controller) *MockKubectl {
	mock := &MockKubectl{ctrl: ctrl}
	mock.recorder = &MockKubectlMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockKubectl) EXPECT() *MockKubectlMockRecorder {
	return m.recorder
}

// Apply mocks base method
func (m *MockKubectl) Apply(path string) Command {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Apply", path)
	ret0, _ := ret[0].(Command)
	return ret0
}

// Apply indicates an expected call of Apply
func (mr *MockKubectlMockRecorder) Apply(path interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Apply", reflect.TypeOf((*MockKubectl)(nil).Apply), path)
}

// Delete mocks base method
func (m *MockKubectl) Delete(args ...string) Command {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Delete", varargs...)
	ret0, _ := ret[0].(Command)
	return ret0
}

// Delete indicates an expected call of Delete
func (mr *MockKubectlMockRecorder) Delete(args ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockKubectl)(nil).Delete), args...)
}

// Get mocks base method
func (m *MockKubectl) Get(args ...string) Command {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Get", varargs...)
	ret0, _ := ret[0].(Command)
	return ret0
}

// Get indicates an expected call of Get
func (mr *MockKubectlMockRecorder) Get(args ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockKubectl)(nil).Get), args...)
}

// getLogger mocks base method
func (m *MockKubectl) getLogger() logr.Logger {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "getLogger")
	ret0, _ := ret[0].(logr.Logger)
	return ret0
}

// getLogger indicates an expected call of getLogger
func (mr *MockKubectlMockRecorder) getLogger() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "getLogger", reflect.TypeOf((*MockKubectl)(nil).getLogger))
}

// getPath mocks base method
func (m *MockKubectl) getPath() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "getPath")
	ret0, _ := ret[0].(string)
	return ret0
}

// getPath indicates an expected call of getPath
func (mr *MockKubectlMockRecorder) getPath() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "getPath", reflect.TypeOf((*MockKubectl)(nil).getPath))
}

// MockCommand is a mock of Command interface
type MockCommand struct {
	ctrl     *gomock.Controller
	recorder *MockCommandMockRecorder
}

// MockCommandMockRecorder is the mock recorder for MockCommand
type MockCommandMockRecorder struct {
	mock *MockCommand
}

// NewMockCommand creates a new mock instance
func NewMockCommand(ctrl *gomock.Controller) *MockCommand {
	mock := &MockCommand{ctrl: ctrl}
	mock.recorder = &MockCommandMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockCommand) EXPECT() *MockCommandMockRecorder {
	return m.recorder
}

// Run mocks base method
func (m *MockCommand) Run() ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Run")
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Run indicates an expected call of Run
func (mr *MockCommandMockRecorder) Run() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Run", reflect.TypeOf((*MockCommand)(nil).Run))
}

// DryRun mocks base method
func (m *MockCommand) DryRun() ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DryRun")
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DryRun indicates an expected call of DryRun
func (mr *MockCommandMockRecorder) DryRun() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DryRun", reflect.TypeOf((*MockCommand)(nil).DryRun))
}

// WithLogger mocks base method
func (m *MockCommand) WithLogger(logger logr.Logger) *abstractCommand {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WithLogger", logger)
	ret0, _ := ret[0].(*abstractCommand)
	return ret0
}

// WithLogger indicates an expected call of WithLogger
func (mr *MockCommandMockRecorder) WithLogger(logger interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WithLogger", reflect.TypeOf((*MockCommand)(nil).WithLogger), logger)
}
