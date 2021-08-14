package ports

import (
	"github.com/mendezdev/mytheresa-api/internal/core/domain"
	"github.com/mendezdev/mytheresa-api/pkg/apierrors"
)

//go:generate mockgen -destination=mock_discount_service.go -package=ports -source=discount_service.go DiscountService

type DiscountService interface {
	GetDiscounts() ([]domain.Discount, apierrors.ApiErr)
}
