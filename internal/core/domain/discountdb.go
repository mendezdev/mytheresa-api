package domain

type DiscountConditionDB struct {
	FieldName string      `json:"field_name"`
	Operator  string      `json:"operator"`
	Value     interface{} `json:"value"`
	ValueType interface{} `json:"value_type"`
	Apply     int64       `json:"apply"`
}
