// Code generated by MockGen. DO NOT EDIT.
// Source: domain/repository/iuser_repository.go

// Package mock_repository is a generated GoMock package.
package mock_repository

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockIUserRepository is a mock of IUserRepository interface
type MockIUserRepository struct {
	ctrl     *gomock.Controller
	recorder *MockIUserRepositoryMockRecorder
}

// MockIUserRepositoryMockRecorder is the mock recorder for MockIUserRepository
type MockIUserRepositoryMockRecorder struct {
	mock *MockIUserRepository
}

// NewMockIUserRepository creates a new mock instance
func NewMockIUserRepository(ctrl *gomock.Controller) *MockIUserRepository {
	mock := &MockIUserRepository{ctrl: ctrl}
	mock.recorder = &MockIUserRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockIUserRepository) EXPECT() *MockIUserRepositoryMockRecorder {
	return m.recorder
}

// Save mocks base method
func (m *MockIUserRepository) Save(arg0 string) uint {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Save", arg0)
	ret0, _ := ret[0].(uint)
	return ret0
}

// Save indicates an expected call of Save
func (mr *MockIUserRepositoryMockRecorder) Save(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Save", reflect.TypeOf((*MockIUserRepository)(nil).Save), arg0)
}

// FindOne mocks base method
func (m *MockIUserRepository) FindOne(arg0 string) uint {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindOne", arg0)
	ret0, _ := ret[0].(uint)
	return ret0
}

// FindOne indicates an expected call of FindOne
func (mr *MockIUserRepositoryMockRecorder) FindOne(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindOne", reflect.TypeOf((*MockIUserRepository)(nil).FindOne), arg0)
}
