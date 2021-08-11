package discountrepo

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/mendezdev/mytheresa-api/internal/core/domain"
	"github.com/mendezdev/mytheresa-api/internal/core/ports"
)

const (
	path = "../../store/discountdb.json"
)

var (
	discounts []domain.Discount
)

type inmemory struct {
	discounts []domain.Discount
}

func NewInMemory() ports.DiscountRepository {
	return &inmemory{
		discounts: discounts,
	}
}

func (mem *inmemory) GetDiscounts() ([]domain.Discount, error) {
	return mem.discounts, nil
}

func init() {
	fmt.Println("trying to initialize discounts inmemory db...")
	data, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	fmt.Println("data read ok!")

	fmt.Println("unmarshaling data...")
	jsonErr := json.Unmarshal(data, &discounts)
	if jsonErr != nil {
		panic(jsonErr)
	}
	fmt.Println("unmarshaling ok!")

	fmt.Printf("discounts: %d\n", len(discounts))
}
