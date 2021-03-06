// Code generated by MockGen. DO NOT EDIT.
// Source: environment.go

// Package mocks is a generated GoMock package.
package mocks

import (
	io "io"
	reflect "reflect"
	time "time"

	gomock "github.com/golang/mock/gomock"
	common "github.com/stackrox/rox/roxctl/common"
	io2 "github.com/stackrox/rox/roxctl/common/io"
	"github.com/stackrox/rox/roxctl/common/logger"
	grpc "google.golang.org/grpc"
)

// MockEnvironment is a mock of Environment interface.
type MockEnvironment struct {
	ctrl     *gomock.Controller
	recorder *MockEnvironmentMockRecorder
}

// MockEnvironmentMockRecorder is the mock recorder for MockEnvironment.
type MockEnvironmentMockRecorder struct {
	mock *MockEnvironment
}

// NewMockEnvironment creates a new mock instance.
func NewMockEnvironment(ctrl *gomock.Controller) *MockEnvironment {
	mock := &MockEnvironment{ctrl: ctrl}
	mock.recorder = &MockEnvironmentMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockEnvironment) EXPECT() *MockEnvironmentMockRecorder {
	return m.recorder
}

// ColorWriter mocks base method.
func (m *MockEnvironment) ColorWriter() io.Writer {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ColorWriter")
	ret0, _ := ret[0].(io.Writer)
	return ret0
}

// ColorWriter indicates an expected call of ColorWriter.
func (mr *MockEnvironmentMockRecorder) ColorWriter() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ColorWriter", reflect.TypeOf((*MockEnvironment)(nil).ColorWriter))
}

// ConnectNames mocks base method.
func (m *MockEnvironment) ConnectNames() (string, string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ConnectNames")
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(string)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ConnectNames indicates an expected call of ConnectNames.
func (mr *MockEnvironmentMockRecorder) ConnectNames() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ConnectNames", reflect.TypeOf((*MockEnvironment)(nil).ConnectNames))
}

// GRPCConnection mocks base method.
func (m *MockEnvironment) GRPCConnection() (*grpc.ClientConn, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GRPCConnection")
	ret0, _ := ret[0].(*grpc.ClientConn)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GRPCConnection indicates an expected call of GRPCConnection.
func (mr *MockEnvironmentMockRecorder) GRPCConnection() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GRPCConnection", reflect.TypeOf((*MockEnvironment)(nil).GRPCConnection))
}

// HTTPClient mocks base method.
func (m *MockEnvironment) HTTPClient(timeout time.Duration) (common.RoxctlHTTPClient, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HTTPClient", timeout)
	ret0, _ := ret[0].(common.RoxctlHTTPClient)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// HTTPClient indicates an expected call of HTTPClient.
func (mr *MockEnvironmentMockRecorder) HTTPClient(timeout interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HTTPClient", reflect.TypeOf((*MockEnvironment)(nil).HTTPClient), timeout)
}

// InputOutput mocks base method.
func (m *MockEnvironment) InputOutput() io2.IO {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InputOutput")
	ret0, _ := ret[0].(io2.IO)
	return ret0
}

// InputOutput indicates an expected call of InputOutput.
func (mr *MockEnvironmentMockRecorder) InputOutput() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InputOutput", reflect.TypeOf((*MockEnvironment)(nil).InputOutput))
}

// Logger mocks base method.
func (m *MockEnvironment) Logger() logger.Logger {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Logger")
	ret0, _ := ret[0].(logger.Logger)
	return ret0
}

// Logger indicates an expected call of Logger.
func (mr *MockEnvironmentMockRecorder) Logger() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Logger", reflect.TypeOf((*MockEnvironment)(nil).Logger))
}

// MockIO is a mock of IO interface.
type MockIO struct {
	ctrl     *gomock.Controller
	recorder *MockIOMockRecorder
}

// MockIOMockRecorder is the mock recorder for MockIO.
type MockIOMockRecorder struct {
	mock *MockIO
}

// NewMockIO creates a new mock instance.
func NewMockIO(ctrl *gomock.Controller) *MockIO {
	mock := &MockIO{ctrl: ctrl}
	mock.recorder = &MockIOMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIO) EXPECT() *MockIOMockRecorder {
	return m.recorder
}

// ErrOut mocks base method.
func (m *MockIO) ErrOut() io.Writer {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ErrOut")
	ret0, _ := ret[0].(io.Writer)
	return ret0
}

// ErrOut indicates an expected call of ErrOut.
func (mr *MockIOMockRecorder) ErrOut() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ErrOut", reflect.TypeOf((*MockIO)(nil).ErrOut))
}

// In mocks base method.
func (m *MockIO) In() io.Reader {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "In")
	ret0, _ := ret[0].(io.Reader)
	return ret0
}

// In indicates an expected call of In.
func (mr *MockIOMockRecorder) In() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "In", reflect.TypeOf((*MockIO)(nil).In))
}

// Out mocks base method.
func (m *MockIO) Out() io.Writer {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Out")
	ret0, _ := ret[0].(io.Writer)
	return ret0
}

// Out indicates an expected call of Out.
func (mr *MockIOMockRecorder) Out() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Out", reflect.TypeOf((*MockIO)(nil).Out))
}

// MockLogger is a mock of Logger interface.
type MockLogger struct {
	ctrl     *gomock.Controller
	recorder *MockLoggerMockRecorder
}

// MockLoggerMockRecorder is the mock recorder for MockLogger.
type MockLoggerMockRecorder struct {
	mock *MockLogger
}

// NewMockLogger creates a new mock instance.
func NewMockLogger(ctrl *gomock.Controller) *MockLogger {
	mock := &MockLogger{ctrl: ctrl}
	mock.recorder = &MockLoggerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockLogger) EXPECT() *MockLoggerMockRecorder {
	return m.recorder
}

// ErrfLn mocks base method.
func (m *MockLogger) ErrfLn(format string, a ...interface{}) {
	m.ctrl.T.Helper()
	varargs := []interface{}{format}
	for _, a_2 := range a {
		varargs = append(varargs, a_2)
	}
	m.ctrl.Call(m, "ErrfLn", varargs...)
}

// ErrfLn indicates an expected call of ErrfLn.
func (mr *MockLoggerMockRecorder) ErrfLn(format interface{}, a ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{format}, a...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ErrfLn", reflect.TypeOf((*MockLogger)(nil).ErrfLn), varargs...)
}

// InfofLn mocks base method.
func (m *MockLogger) InfofLn(format string, a ...interface{}) {
	m.ctrl.T.Helper()
	varargs := []interface{}{format}
	for _, a_2 := range a {
		varargs = append(varargs, a_2)
	}
	m.ctrl.Call(m, "InfofLn", varargs...)
}

// InfofLn indicates an expected call of InfofLn.
func (mr *MockLoggerMockRecorder) InfofLn(format interface{}, a ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{format}, a...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InfofLn", reflect.TypeOf((*MockLogger)(nil).InfofLn), varargs...)
}

// PrintfLn mocks base method.
func (m *MockLogger) PrintfLn(format string, a ...interface{}) {
	m.ctrl.T.Helper()
	varargs := []interface{}{format}
	for _, a_2 := range a {
		varargs = append(varargs, a_2)
	}
	m.ctrl.Call(m, "PrintfLn", varargs...)
}

// PrintfLn indicates an expected call of PrintfLn.
func (mr *MockLoggerMockRecorder) PrintfLn(format interface{}, a ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{format}, a...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PrintfLn", reflect.TypeOf((*MockLogger)(nil).PrintfLn), varargs...)
}

// WarnfLn mocks base method.
func (m *MockLogger) WarnfLn(format string, a ...interface{}) {
	m.ctrl.T.Helper()
	varargs := []interface{}{format}
	for _, a_2 := range a {
		varargs = append(varargs, a_2)
	}
	m.ctrl.Call(m, "WarnfLn", varargs...)
}

// WarnfLn indicates an expected call of WarnfLn.
func (mr *MockLoggerMockRecorder) WarnfLn(format interface{}, a ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{format}, a...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WarnfLn", reflect.TypeOf((*MockLogger)(nil).WarnfLn), varargs...)
}
