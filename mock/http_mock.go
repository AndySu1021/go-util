// Code generated by MockGen. DO NOT EDIT.
// Source: interface/http.go

// Package mock is a generated GoMock package.
package mock

import (
	http "net/http"
	reflect "reflect"
	time "time"

	gomock "github.com/golang/mock/gomock"
)

// MockIHttpClient is a mock of IHttpClient interface.
type MockIHttpClient struct {
	ctrl     *gomock.Controller
	recorder *MockIHttpClientMockRecorder
}

// MockIHttpClientMockRecorder is the mock recorder for MockIHttpClient.
type MockIHttpClientMockRecorder struct {
	mock *MockIHttpClient
}

// NewMockIHttpClient creates a new mock instance.
func NewMockIHttpClient(ctrl *gomock.Controller) *MockIHttpClient {
	mock := &MockIHttpClient{ctrl: ctrl}
	mock.recorder = &MockIHttpClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIHttpClient) EXPECT() *MockIHttpClientMockRecorder {
	return m.recorder
}

// Send mocks base method.
func (m *MockIHttpClient) Send() (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Send")
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Send indicates an expected call of Send.
func (mr *MockIHttpClientMockRecorder) Send() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Send", reflect.TypeOf((*MockIHttpClient)(nil).Send))
}

// SetRequest mocks base method.
func (m *MockIHttpClient) SetRequest(request *http.Request) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetRequest", request)
}

// SetRequest indicates an expected call of SetRequest.
func (mr *MockIHttpClientMockRecorder) SetRequest(request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetRequest", reflect.TypeOf((*MockIHttpClient)(nil).SetRequest), request)
}

// SetTimeout mocks base method.
func (m *MockIHttpClient) SetTimeout(timeout time.Duration) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetTimeout", timeout)
}

// SetTimeout indicates an expected call of SetTimeout.
func (mr *MockIHttpClientMockRecorder) SetTimeout(timeout interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetTimeout", reflect.TypeOf((*MockIHttpClient)(nil).SetTimeout), timeout)
}
