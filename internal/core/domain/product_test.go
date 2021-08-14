package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBuildProductResponseOk(t *testing.T) {
	dr := DiscountResult{
		Percentage: 10,
	}

	product := Product{
		Sku:      "00001",
		Name:     "test name",
		Category: "test category",
		Price:    10000,
	}

	response, err := product.BuildProductResponse(&dr)
	assert.Nil(t, err)
	assert.NotNil(t, response)
	assert.Equal(t, product.Sku, response.Sku)
	assert.Equal(t, product.Name, response.Name)
	assert.Equal(t, product.Category, response.Category)
	assert.Equal(t, product.Price, response.Price.Original)
	assert.Equal(t, int64(9000), *response.Price.Final)
	assert.Equal(t, "EUR", response.Price.Currency)
	assert.Equal(t, "10%", *response.Price.DiscountPercentage)
}

func TestFromLongToDecimalOk(t *testing.T) {
	var price int64 = 10050

	var expected float64 = 100.50

	result, err := fromLongToDecimal(price)
	assert.Nil(t, err)
	assert.Equal(t, expected, result)
}

func TestFromLongToDecimalError(t *testing.T) {
	var price int64 = 10

	_, err := fromLongToDecimal(price)
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "should have 3 digits minimum")
}

func TestFormatFinalPriceResponseOk(t *testing.T) {
	testCases := []struct {
		Name          string
		Price         float64
		PriceExpected int64
		Err           error
	}{
		{"with zero decimals", 100.00, 10000, nil},
		{"without zero decimals", 100.67, 10067, nil},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			result, err := formatFinalPriceResponse(tc.Price)
			assert.Equal(t, err, tc.Err)
			assert.Equal(t, tc.PriceExpected, result)
		})
	}
}
