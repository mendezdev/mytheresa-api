package producthdl

type GetProductDto struct {
	Category string `json:"category"`
	LessThan *int64 `json:"less_than"`
}
