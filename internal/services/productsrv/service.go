package productsrv

import (
	"github.com/mendezdev/mytheresa-api/internal/core/domain"
	"github.com/mendezdev/mytheresa-api/internal/ports"
)

const (
	productsLimit = 5
)

type service struct {
	productRepo     ports.ProductRepository
	discountService ports.DiscountService
}

func New(pr ports.ProductRepository, ds ports.DiscountService) ports.ProductService {
	return &service{
		productRepo: pr,
	}
}

func (srv *service) GetProductsByCategory(pf domain.ProductFilter) ([]domain.ProductAPIResponse, error) {
	productsByCategory, err := srv.productRepo.GetByCategory(pf.Category, productsLimit, pf.LesstThan)
	if err != nil {
		return nil, err
	}

	responses := make([]domain.ProductAPIResponse, 0)
	if len(productsByCategory) == 0 {
		return responses, nil
	}

	discounts, discountErr := srv.discountService.GetDiscounts()
	if discountErr != nil {
		return nil, discountErr
	}

	for _, product := range productsByCategory {
		for _, discount := range discounts {
			discountResult := discount.GetDiscount(product.ToDiscountRequest())
			productResponse, productErr := product.BuildProductResponse(discountResult)
			if productErr != nil {
				return nil, productErr
			}
			responses = append(responses, *productResponse)
		}
	}
	return responses, nil
}
