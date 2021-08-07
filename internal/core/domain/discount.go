package domain

const (
	EqualOperator = "=="
)

type DiscountResult struct {
	Percentage int64
}

type DiscountRequest map[string]interface{}

type Discount struct {
	FieldName string      `json:"field_name"`
	Operator  string      `json:"operator"`
	Value     interface{} `json:"value"`
	ValueType string      `json:"value_type"`
	Priority  int         `json:"priority"`
	Apply     int64       `json:"apply"`
}

func NewDiscount() Discount {
	return Discount{}
}

func (d Discount) GetDiscount(dr DiscountRequest) *DiscountResult {
	if len(dr) == 0 {
		return nil
	}

	drVal, ok := dr[d.FieldName]
	if !ok {
		return nil
	}

	isTypeOk := checkInterfaceType(drVal, d.ValueType)
	if !isTypeOk {
		return nil
	}

	var discountResult DiscountResult

	switch d.Operator {
	case EqualOperator:
		if d.Value == drVal {
			discountResult.Percentage = d.Apply
			return &discountResult
		}
	}

	return nil
}

func checkInterfaceType(val interface{}, valueType string) bool {
	switch valueType {
	case "long":
		_, ok := val.(int64)
		return ok
	case "string":
		_, ok := val.(string)
		return ok
	default:
		return false
	}
}
