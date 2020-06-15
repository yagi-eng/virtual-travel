// Code generated by MockGen. DO NOT EDIT.
// Source: usecases/igateway/igooglemap_gateway.go

// Package mock_igateway is a generated GoMock package.
package mock_igateway

import (
	gomock "github.com/golang/mock/gomock"
	googlemapdto "github.com/yagi-eng/place-search/usecases/dto/googlemapdto"
	reflect "reflect"
)

// MockIGoogleMapGateway is a mock of IGoogleMapGateway interface
type MockIGoogleMapGateway struct {
	ctrl     *gomock.Controller
	recorder *MockIGoogleMapGatewayMockRecorder
}

// MockIGoogleMapGatewayMockRecorder is the mock recorder for MockIGoogleMapGateway
type MockIGoogleMapGatewayMockRecorder struct {
	mock *MockIGoogleMapGateway
}

// NewMockIGoogleMapGateway creates a new mock instance
func NewMockIGoogleMapGateway(ctrl *gomock.Controller) *MockIGoogleMapGateway {
	mock := &MockIGoogleMapGateway{ctrl: ctrl}
	mock.recorder = &MockIGoogleMapGatewayMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockIGoogleMapGateway) EXPECT() *MockIGoogleMapGatewayMockRecorder {
	return m.recorder
}

// GetPlaceDetailsAndPhotoURLsFromQuery mocks base method
func (m *MockIGoogleMapGateway) GetPlaceDetailsAndPhotoURLsFromQuery(arg0 string) []googlemapdto.Output {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPlaceDetailsAndPhotoURLsFromQuery", arg0)
	ret0, _ := ret[0].([]googlemapdto.Output)
	return ret0
}

// GetPlaceDetailsAndPhotoURLsFromQuery indicates an expected call of GetPlaceDetailsAndPhotoURLsFromQuery
func (mr *MockIGoogleMapGatewayMockRecorder) GetPlaceDetailsAndPhotoURLsFromQuery(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPlaceDetailsAndPhotoURLsFromQuery", reflect.TypeOf((*MockIGoogleMapGateway)(nil).GetPlaceDetailsAndPhotoURLsFromQuery), arg0)
}

// GetPlaceDetailsAndPhotoURLsFromLatLng mocks base method
func (m *MockIGoogleMapGateway) GetPlaceDetailsAndPhotoURLsFromLatLng(arg0, arg1 float64) []googlemapdto.Output {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPlaceDetailsAndPhotoURLsFromLatLng", arg0, arg1)
	ret0, _ := ret[0].([]googlemapdto.Output)
	return ret0
}

// GetPlaceDetailsAndPhotoURLsFromLatLng indicates an expected call of GetPlaceDetailsAndPhotoURLsFromLatLng
func (mr *MockIGoogleMapGatewayMockRecorder) GetPlaceDetailsAndPhotoURLsFromLatLng(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPlaceDetailsAndPhotoURLsFromLatLng", reflect.TypeOf((*MockIGoogleMapGateway)(nil).GetPlaceDetailsAndPhotoURLsFromLatLng), arg0, arg1)
}

// GetPlaceDetailsAndPhotoURLs mocks base method
func (m *MockIGoogleMapGateway) GetPlaceDetailsAndPhotoURLs(arg0 []string, arg1 bool) []googlemapdto.Output {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPlaceDetailsAndPhotoURLs", arg0, arg1)
	ret0, _ := ret[0].([]googlemapdto.Output)
	return ret0
}

// GetPlaceDetailsAndPhotoURLs indicates an expected call of GetPlaceDetailsAndPhotoURLs
func (mr *MockIGoogleMapGatewayMockRecorder) GetPlaceDetailsAndPhotoURLs(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPlaceDetailsAndPhotoURLs", reflect.TypeOf((*MockIGoogleMapGateway)(nil).GetPlaceDetailsAndPhotoURLs), arg0, arg1)
}
