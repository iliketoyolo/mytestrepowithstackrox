// Code generated by MockGen. DO NOT EDIT.
// Source: client.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	payload "github.com/stackrox/default-authz-plugin/pkg/payload"
)

// MockClient is a mock of Client interface.
type MockClient struct {
	ctrl     *gomock.Controller
	recorder *MockClientMockRecorder
}

// MockClientMockRecorder is the mock recorder for MockClient.
type MockClientMockRecorder struct {
	mock *MockClient
}

// NewMockClient creates a new mock instance.
func NewMockClient(ctrl *gomock.Controller) *MockClient {
	mock := &MockClient{ctrl: ctrl}
	mock.recorder = &MockClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockClient) EXPECT() *MockClientMockRecorder {
	return m.recorder
}

// ForUser mocks base method.
func (m *MockClient) ForUser(ctx context.Context, principal payload.Principal, scopes ...payload.AccessScope) ([]payload.AccessScope, []payload.AccessScope, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, principal}
	for _, a := range scopes {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ForUser", varargs...)
	ret0, _ := ret[0].([]payload.AccessScope)
	ret1, _ := ret[1].([]payload.AccessScope)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ForUser indicates an expected call of ForUser.
func (mr *MockClientMockRecorder) ForUser(ctx, principal interface{}, scopes ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, principal}, scopes...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ForUser", reflect.TypeOf((*MockClient)(nil).ForUser), varargs...)
}
