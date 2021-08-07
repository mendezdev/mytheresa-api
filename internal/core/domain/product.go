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
