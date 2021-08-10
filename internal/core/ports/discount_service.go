package ports

import "github.com/mendezdev/mytheresa-api/internal/core/domain"

//go:generate mockgen -destination=mock_discount_service.go -package=ports -source=discount_service.go DiscountService

type DiscountService interface {
	GetDiscounts() ([]domain.Discount, error)
}
