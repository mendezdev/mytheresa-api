package productsrv

import (
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/mendezdev/mytheresa-api/internal/core/domain"
	"github.com/mendezdev/mytheresa-api/internal/ports"
	"github.com/stretchr/testify/assert"
)

func TestProductOk(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockProductRepo := ports.NewMockProductRepository(mockCtrl)
	mockDiscountSrv := ports.NewMockDiscountService(mockCtrl)

	productSrv := New(mockProductRepo, mockDiscountSrv)

	d1 := domain.Discount{
		FieldName: "category",
		Operator:  "==",
		Value:     "boots",
		ValueType: "string",
		Priority:  1,
		Apply:     30,
	}
	discounts := []domain.Discount{d1}

	mockDiscountSrv.
		EXPECT().
		GetDiscounts().
		Return(discounts, nil).
		Times(1)

	pf := domain.ProductFilter{
		Category: "boots",
	}

	bootProduct := domain.Product{
		Sku:      "00001",
		Name:     "p1",
		Category: "boots",
		Price:    30000,
	}

	shoeProduct := domain.Product{
		Sku:      "00002",
		Name:     "p2",
		Category: "shoes",
		Price:    12000,
	}

	productsExpected := []domain.Product{bootProduct, shoeProduct}
	mockProductRepo.
		EXPECT().
		GetByCategory(pf.Category, productsLimit, pf.LesstThan).
		Return(productsExpected, nil).
		Times(1)

	results, err := productSrv.GetProductsByCategory(pf)
	assert.Nil(t, err)
	assert.Len(t, results, 2)

	for _, result := range results {
		skuOk := false
		nameOk := false
		categoryOk := false
		priceOk := false
		for _, pe := range productsExpected {
			if pe.Sku == result.Sku {
				skuOk = true
			}
			if pe.Name == result.Name {
				nameOk = true
			}
			if pe.Category == result.Category {
				categoryOk = true
			}
			if pe.Price == result.Price.Original {
				priceOk = true
			}
		}
		assert.True(t, skuOk)
		assert.True(t, nameOk)
		assert.True(t, categoryOk)
		assert.True(t, priceOk)

		// check discount price
		if result.Category == bootProduct.Category {
			assert.Equal(t, *result.Price.DiscountPercentage, fmt.Sprintf("%d%s", d1.Apply, "%"))
			assert.Equal(t, *result.Price.Final, int64(21000))
		}
	}
}
