package domain

type ProductAPIResponse struct {
	Sku      string           `json:"sku"`
	Name     string           `json:"name"`
	Category string           `json:"category"`
	Price    PriceAPIResponse `json:"price"`
}

type PriceAPIResponse struct {
	Original           int64  `json:"original"`
	Final              int64  `json:"final"`
	DiscountPercentage string `json:"discount_percentage"`
	Currency           string `json:"currency"`
}

type ProductDB struct {
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

func (p Product) GetProductsBy(category string, priceLessThan *int64) ([]ProductAPIResponse, error) {
	return nil, nil
}
