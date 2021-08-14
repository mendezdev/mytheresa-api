package ports

import (
	"github.com/mendezdev/mytheresa-api/internal/core/domain"
	"github.com/mendezdev/mytheresa-api/pkg/apierrors"
)

//go:generate mockgen -destination=mock_product_service.go -package=ports -source=product_service.go ProductService

type ProductService interface {
	GetProductsByCategory(category string, lessThan *int64) ([]domain.ProductAPIResponse, apierrors.ApiErr)
}
