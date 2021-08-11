package discountsrv

import (
	"github.com/mendezdev/mytheresa-api/internal/core/domain"
	"github.com/mendezdev/mytheresa-api/internal/core/ports"
)

type service struct {
	discountRepo ports.DiscountRepository
}

func New(dr ports.DiscountRepository) ports.DiscountService {
	return &service{
		discountRepo: dr,
	}
}

func (srv *service) GetDiscounts() ([]domain.Discount, error) {
	discounts, err := srv.discountRepo.GetDiscounts()
	if err != nil {
		return nil, err
	}

	return discounts, nil
}