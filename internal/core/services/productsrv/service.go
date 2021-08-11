package productsrv

import (
	"github.com/mendezdev/mytheresa-api/internal/core/domain"
	"github.com/mendezdev/mytheresa-api/internal/core/ports"
)

const (
	productsLimit = int64(5)
)

type service struct {
	productRepo     ports.ProductRepository
	discountService ports.DiscountService
}

func New(pr ports.ProductRepository, ds ports.DiscountService) ports.ProductService {
	return &service{
		productRepo:     pr,
		discountService: ds,
	}
}

func (srv *service) GetProductsByCategory(category string, lessThan *int64) ([]domain.ProductAPIResponse, error) {
	productsByCategory, err := srv.productRepo.GetByCategory(category, productsLimit, lessThan)
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
		discountResult := getDiscount(discounts, product.ToDiscountRequest())
		productResponse, productErr := product.BuildProductResponse(discountResult)
		if productErr != nil {
			return nil, productErr
		}
		responses = append(responses, *productResponse)
	}
	return responses, nil
}

func getDiscount(discounts []domain.Discount, request domain.DiscountRequest) *domain.DiscountResult {
	var discountResult *domain.DiscountResult

	for _, discount := range discounts {
		discountResult = discount.GetDiscount(request)
		if discountResult != nil {
			break
		}
	}

	return discountResult
}
