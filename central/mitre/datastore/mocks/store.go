// Code generated by MockGen. DO NOT EDIT.
// Source: store.go

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	storage "github.com/stackrox/rox/generated/storage"
)

// MockMitreAttackReadOnlyDataStore is a mock of MitreAttackReadOnlyDataStore interface.
type MockMitreAttackReadOnlyDataStore struct {
	ctrl     *gomock.Controller
	recorder *MockMitreAttackReadOnlyDataStoreMockRecorder
}

// MockMitreAttackReadOnlyDataStoreMockRecorder is the mock recorder for MockMitreAttackReadOnlyDataStore.
type MockMitreAttackReadOnlyDataStoreMockRecorder struct {
	mock *MockMitreAttackReadOnlyDataStore
}

// NewMockMitreAttackReadOnlyDataStore creates a new mock instance.
func NewMockMitreAttackReadOnlyDataStore(ctrl *gomock.Controller) *MockMitreAttackReadOnlyDataStore {
	mock := &MockMitreAttackReadOnlyDataStore{ctrl: ctrl}
	mock.recorder = &MockMitreAttackReadOnlyDataStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMitreAttackReadOnlyDataStore) EXPECT() *MockMitreAttackReadOnlyDataStoreMockRecorder {
	return m.recorder
}

// Get mocks base method.
func (m *MockMitreAttackReadOnlyDataStore) Get(id string) (*storage.MitreAttackVector, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", id)
	ret0, _ := ret[0].(*storage.MitreAttackVector)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockMitreAttackReadOnlyDataStoreMockRecorder) Get(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockMitreAttackReadOnlyDataStore)(nil).Get), id)
}

// GetAll mocks base method.
func (m *MockMitreAttackReadOnlyDataStore) GetAll() []*storage.MitreAttackVector {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAll")
	ret0, _ := ret[0].([]*storage.MitreAttackVector)
	return ret0
}

// GetAll indicates an expected call of GetAll.
func (mr *MockMitreAttackReadOnlyDataStoreMockRecorder) GetAll() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockMitreAttackReadOnlyDataStore)(nil).GetAll))
}
