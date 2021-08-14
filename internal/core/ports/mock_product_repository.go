// Code generated by MockGen. DO NOT EDIT.
// Source: product_repository.go

// Package ports is a generated GoMock package.
package ports

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	domain "github.com/mendezdev/mytheresa-api/internal/core/domain"
	apierrors "github.com/mendezdev/mytheresa-api/pkg/apierrors"
)

// MockProductRepository is a mock of ProductRepository interface.
type MockProductRepository struct {
	ctrl     *gomock.Controller
	recorder *MockProductRepositoryMockRecorder
}

// MockProductRepositoryMockRecorder is the mock recorder for MockProductRepository.
type MockProductRepositoryMockRecorder struct {
	mock *MockProductRepository
}

// NewMockProductRepository creates a new mock instance.
func NewMockProductRepository(ctrl *gomock.Controller) *MockProductRepository {
	mock := &MockProductRepository{ctrl: ctrl}
	mock.recorder = &MockProductRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockProductRepository) EXPECT() *MockProductRepositoryMockRecorder {
	return m.recorder
}

// GetByCategory mocks base method.
func (m *MockProductRepository) GetByCategory(category string, lessThan *int64) ([]domain.Product, apierrors.ApiErr) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByCategory", category, lessThan)
	ret0, _ := ret[0].([]domain.Product)
	ret1, _ := ret[1].(apierrors.ApiErr)
	return ret0, ret1
}

// GetByCategory indicates an expected call of GetByCategory.
func (mr *MockProductRepositoryMockRecorder) GetByCategory(category, lessThan interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByCategory", reflect.TypeOf((*MockProductRepository)(nil).GetByCategory), category, lessThan)
}
