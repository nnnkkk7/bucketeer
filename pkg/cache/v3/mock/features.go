// Code generated by MockGen. DO NOT EDIT.
// Source: features.go

// Package mock is a generated GoMock package.
package mock

import (
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"

	feature "github.com/bucketeer-io/bucketeer/proto/feature"
)

// MockFeaturesCache is a mock of FeaturesCache interface.
type MockFeaturesCache struct {
	ctrl     *gomock.Controller
	recorder *MockFeaturesCacheMockRecorder
}

// MockFeaturesCacheMockRecorder is the mock recorder for MockFeaturesCache.
type MockFeaturesCacheMockRecorder struct {
	mock *MockFeaturesCache
}

// NewMockFeaturesCache creates a new mock instance.
func NewMockFeaturesCache(ctrl *gomock.Controller) *MockFeaturesCache {
	mock := &MockFeaturesCache{ctrl: ctrl}
	mock.recorder = &MockFeaturesCacheMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockFeaturesCache) EXPECT() *MockFeaturesCacheMockRecorder {
	return m.recorder
}

// Get mocks base method.
func (m *MockFeaturesCache) Get(environmentNamespace string) (*feature.Features, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", environmentNamespace)
	ret0, _ := ret[0].(*feature.Features)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockFeaturesCacheMockRecorder) Get(environmentNamespace interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockFeaturesCache)(nil).Get), environmentNamespace)
}

// Put mocks base method.
func (m *MockFeaturesCache) Put(features *feature.Features, environmentNamespace string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Put", features, environmentNamespace)
	ret0, _ := ret[0].(error)
	return ret0
}

// Put indicates an expected call of Put.
func (mr *MockFeaturesCacheMockRecorder) Put(features, environmentNamespace interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Put", reflect.TypeOf((*MockFeaturesCache)(nil).Put), features, environmentNamespace)
}
