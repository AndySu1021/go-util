// Code generated by MockGen. DO NOT EDIT.
// Source: interface/storage.go

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	io "io"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockIStorage is a mock of IStorage interface.
type MockIStorage struct {
	ctrl     *gomock.Controller
	recorder *MockIStorageMockRecorder
}

// MockIStorageMockRecorder is the mock recorder for MockIStorage.
type MockIStorageMockRecorder struct {
	mock *MockIStorage
}

// NewMockIStorage creates a new mock instance.
func NewMockIStorage(ctrl *gomock.Controller) *MockIStorage {
	mock := &MockIStorage{ctrl: ctrl}
	mock.recorder = &MockIStorageMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIStorage) EXPECT() *MockIStorageMockRecorder {
	return m.recorder
}

// Upload mocks base method.
func (m *MockIStorage) Upload(ctx context.Context, reader io.Reader, filename string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Upload", ctx, reader, filename)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Upload indicates an expected call of Upload.
func (mr *MockIStorageMockRecorder) Upload(ctx, reader, filename interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Upload", reflect.TypeOf((*MockIStorage)(nil).Upload), ctx, reader, filename)
}
