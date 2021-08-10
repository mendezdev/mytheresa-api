package productrepo

import (
	"github.com/mendezdev/mytheresa-api/internal/core/domain"
	"github.com/mendezdev/mytheresa-api/internal/core/ports"
)

type inmemory struct{}

func NewInMemory() ports.ProductRepository {
	return &inmemory{}
}

func (mem *inmemory) GetByCategory(category string, limit int64, lessThan *int64) ([]domain.Product, error) {
	return nil, nil
}
