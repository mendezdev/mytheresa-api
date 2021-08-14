package ports

import (
	"github.com/mendezdev/mytheresa-api/internal/core/domain"
	"github.com/mendezdev/mytheresa-api/pkg/apierrors"
)

//go:generate mockgen -destination=mock_discount_repository.go -package=ports -source=discount_repository.go DiscountRepository

type DiscountRepository interface {
	GetDiscounts() ([]domain.Discount, apierrors.ApiErr)
}
