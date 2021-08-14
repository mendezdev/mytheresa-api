package domain

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type ProductAPIResponse struct {
	Sku      string           `json:"sku"`
	Name     string           `json:"name"`
	Category string           `json:"category"`
	Price    PriceAPIResponse `json:"price"`
}

type PriceAPIResponse struct {
	Original           int64   `json:"original"`
	Final              *int64  `json:"final"`
	DiscountPercentage *string `json:"discount_percentage"`
	Currency           string  `json:"currency"`
}

type Product struct {
	Sku      string `json:"sku"`
	Name     string `json:"name"`
	Category string `json:"category"`
	Price    int64  `json:"price"`
}

func NewProduct() Product {
	return Product{}
}

func (p Product) BuildProductResponse(dr *DiscountResult) (*ProductAPIResponse, error) {
	price := PriceAPIResponse{
		Original: p.Price,
		Currency: "EUR",
	}

	if dr != nil {
		discountPercentage := fmt.Sprintf("%d%s", dr.Percentage, "%")
		price.DiscountPercentage = &discountPercentage

		priceParsed, parseErr := fromLongToDecimal(p.Price)
		if parseErr != nil {
			return nil, errors.New("error trying to parse original price to decimal")
		}
		percetage := float64(1.0) - (float64(dr.Percentage) / 100)
		finalPrice := float64(priceParsed) * percetage
		finalPriceParsed, parseErr := formatFinalPriceResponse(finalPrice)
		if parseErr != nil {
			return nil, errors.New("error trying to parse discount price")
		}
		price.Final = &finalPriceParsed

	}

	return &ProductAPIResponse{
		Sku:      p.Sku,
		Name:     p.Name,
		Category: p.Category,
		Price:    price,
	}, nil
}

func (p Product) ToDiscountRequest() DiscountRequest {
	return DiscountRequest{
		"category": p.Category,
		"price":    p.Price,
	}
}

func formatFinalPriceResponse(price float64) (int64, error) {
	parsed := fmt.Sprintf("%.2f", price)
	parsed = strings.Replace(parsed, ".", "", -1)
	result, err := strconv.ParseInt(parsed, 10, 64)
	if err != nil {
		return 0, err
	}
	return result, nil
}

func fromLongToDecimal(val int64) (float64, error) {
	parsed := fmt.Sprintf("%d", val)
	splitted := strings.Split(parsed, "")
	minimunDigits := 3
	if len(splitted) < minimunDigits {
		return 0, fmt.Errorf("error parsing long to decimal, should have %d digits minimum", minimunDigits)
	}
	base := splitted[0 : len(splitted)-2]
	decimals := splitted[len(splitted)-2:]

	var result string
	for _, s := range base {
		result += s
	}
	result += "."
	for _, s := range decimals {
		result += s
	}

	return strconv.ParseFloat(result, 64)
}
