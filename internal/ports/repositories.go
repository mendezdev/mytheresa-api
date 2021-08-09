package ports

import "github.com/mendezdev/mytheresa-api/internal/core/domain"

type ProductRepository interface {
	GetByCategory(category string, limit int64, lessThan *int64) ([]domain.Product, error)
}

type DiscountRepository interface {
	GetDiscounts() ([]domain.Discount, error)
}
