package domain

const (
	EqualOperator = "equal"
)

type DiscountResult int64

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

func (d Discount) GetDiscount(dr DiscountRequest) (*DiscountResult, error) {
	if len(dr) == 0 {
		return nil, nil
	}

	drVal, ok := dr[d.FieldName]
	if !ok {
		return nil, nil
	}

	isTypeOk := checkInterfaceType(drVal, d.ValueType)
	if !isTypeOk {
		return nil, nil
	}

	var discountResult DiscountResult

	switch d.Operator {
	case EqualOperator:
		if d.Value == drVal {
			discountResult = DiscountResult(d.Apply)
		}
	}

	return &discountResult, nil
}

func checkInterfaceType(val interface{}, valueType string) bool {
	switch valueType {
	case "long":
		_, ok := val.(int64)
		return ok
	case "decimal":
		_, ok := val.(float64)
		return ok
	default:
		return false
	}
}
