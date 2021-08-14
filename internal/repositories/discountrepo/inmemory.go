package discountrepo

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"

	"github.com/mendezdev/mytheresa-api/internal/core/domain"
	"github.com/mendezdev/mytheresa-api/internal/core/ports"
	"github.com/mendezdev/mytheresa-api/pkg/apierrors"
)

const (
	path = "./store/discountdb.json"
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

func (mem *inmemory) GetDiscounts() ([]domain.Discount, apierrors.ApiErr) {
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

	fmt.Println("sorting discount by priority")
	sort.Slice(discounts, func(i, j int) bool {
		return discounts[i].Priority < discounts[j].Priority
	})
	fmt.Println("sorting finished!")

	fmt.Printf("discounts: %d\n", len(discounts))
}
