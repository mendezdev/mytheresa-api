package ports

import "github.com/mendezdev/mytheresa-api/internal/core/domain"

//go:generate mockgen -destination=mock_product_service.go -package=ports -source=product_service.go ProductService

type ProductService interface {
	GetProductsByCategory(domain.ProductFilter) ([]domain.ProductAPIResponse, error)
}
