package domain

import "testing"

func TestProductWithoutDiscountOk(t *testing.T) {
	product := Product{
		Sku:      "000001",
		Name:     "BV Lean leather ankle boots",
		Category: "boots",
		Price:    20000,
	}
	discountResult := DiscountResult{
		Percentage: 30,
	}
}
