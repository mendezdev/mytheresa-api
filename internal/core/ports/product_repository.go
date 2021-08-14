package ports

import (
	"github.com/mendezdev/mytheresa-api/internal/core/domain"
	"github.com/mendezdev/mytheresa-api/pkg/apierrors"
)

//go:generate mockgen -destination=mock_product_repository.go -package=ports -source=product_repository.go ProductRepository

type ProductRepository interface {
	GetByCategory(category string, lessThan *int64) ([]domain.Product, apierrors.ApiErr)
}
