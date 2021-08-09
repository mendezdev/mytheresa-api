package ports

import "github.com/mendezdev/mytheresa-api/internal/core/domain"

type ProductService interface {
	GetProductsByCategory(domain.ProductFilter) ([]domain.ProductAPIResponse, error)
}

type DiscountService interface {
	GetDiscounts() ([]domain.Discount, error)
}
