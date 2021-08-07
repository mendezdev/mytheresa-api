package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDiscountFound(t *testing.T) {
	discount := Discount{
		FieldName: "category",
		Operator:  "==",
		Value:     "boots",
		ValueType: "string",
		Priority:  1,
		Apply:     30,
	}

	discountRequest := DiscountRequest{
		"category": "boots",
	}

	discountResult := discount.GetDiscount(discountRequest)

	assert.NotNil(t, discountResult)
	assert.Equal(t, discount.Apply, discountResult.Percentage)
}

func TestDiscountNotFoundByFieldName(t *testing.T) {
	discount := Discount{
		FieldName: "category",
		Operator:  "==",
		Value:     "shirt",
		ValueType: "string",
		Priority:  1,
		Apply:     30,
	}

	discountRequest := DiscountRequest{
		"category": "boots",
	}

	discountResult := discount.GetDiscount(discountRequest)

	assert.Nil(t, discountResult)
}

func TestDiscountNotFoundByType(t *testing.T) {
	discount := Discount{
		FieldName: "price",
		Operator:  "==",
		Value:     1234,
		ValueType: "long",
		Priority:  1,
		Apply:     30,
	}

	discountRequest := DiscountRequest{
		"price": "1234",
	}

	discountResult := discount.GetDiscount(discountRequest)
	assert.Nil(t, discountResult)
}

func TestDiscountNotFoundByDifferentOperator(t *testing.T) {
	discount := Discount{
		FieldName: "category",
		Operator:  "!=",
		Value:     "boots",
		ValueType: "string",
		Priority:  1,
		Apply:     30,
	}

	discountRequest := DiscountRequest{
		"category": "boots",
	}

	discountResult := discount.GetDiscount(discountRequest)
	assert.Nil(t, discountResult)
}
