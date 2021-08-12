package ports

import "github.com/mendezdev/mytheresa-api/internal/core/domain"

//go:generate mockgen -destination=mock_product_repository.go -package=ports -source=product_repository.go ProductRepository

type ProductRepository interface {
	GetByCategory(category string, lessThan *int64) ([]domain.Product, error)
}
