package domain

type DiscountResult struct {
	Value     int64
	Percetage int
}

type Discount struct {
	FieldName string      `json:"field_name"`
	Operator  string      `json:"operator"`
	Value     interface{} `json:"value"`
	Priority  int         `json:"priority"`
	ValueType interface{} `json:"value_type"`
	Apply     int64       `json:"apply"`
}

func NewDiscount() Discount {
	return Discount{}
}

func (d Discount) ApplyDiscount(p Product) (*DiscountResult, error) {
	return nil, nil
}
