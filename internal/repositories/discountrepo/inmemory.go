package discountrepo

import (
	"github.com/mendezdev/mytheresa-api/internal/core/domain"
	"github.com/mendezdev/mytheresa-api/internal/core/ports"
)

type inmemory struct{}

func NewInMemory() ports.DiscountRepository {
	return &inmemory{}
}

func (mem *inmemory) GetDiscounts() ([]domain.Discount, error) {
	return nil, nil
}
