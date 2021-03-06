// Code generated by MockGen. DO NOT EDIT.
// Source: product_service.go

// Package ports is a generated GoMock package.
package ports

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	domain "github.com/mendezdev/mytheresa-api/internal/core/domain"
	apierrors "github.com/mendezdev/mytheresa-api/pkg/apierrors"
)

// MockProductService is a mock of ProductService interface.
type MockProductService struct {
	ctrl     *gomock.Controller
	recorder *MockProductServiceMockRecorder
}

// MockProductServiceMockRecorder is the mock recorder for MockProductService.
type MockProductServiceMockRecorder struct {
	mock *MockProductService
}

// NewMockProductService creates a new mock instance.
func NewMockProductService(ctrl *gomock.Controller) *MockProductService {
	mock := &MockProductService{ctrl: ctrl}
	mock.recorder = &MockProductServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockProductService) EXPECT() *MockProductServiceMockRecorder {
	return m.recorder
}

// GetProductsByCategory mocks base method.
func (m *MockProductService) GetProductsByCategory(category string, lessThan *int64) ([]domain.ProductAPIResponse, apierrors.ApiErr) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProductsByCategory", category, lessThan)
	ret0, _ := ret[0].([]domain.ProductAPIResponse)
	ret1, _ := ret[1].(apierrors.ApiErr)
	return ret0, ret1
}

// GetProductsByCategory indicates an expected call of GetProductsByCategory.
func (mr *MockProductServiceMockRecorder) GetProductsByCategory(category, lessThan interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProductsByCategory", reflect.TypeOf((*MockProductService)(nil).GetProductsByCategory), category, lessThan)
}
