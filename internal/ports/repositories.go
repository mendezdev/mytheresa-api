package ports

import "github.com/mendezdev/mytheresa-api/internal/core/domain"

type ProductRepository interface {
	GetAll() (domain.ItemDB, error)
}
