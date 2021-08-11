package productrepo

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/mendezdev/mytheresa-api/internal/core/domain"
	"github.com/mendezdev/mytheresa-api/internal/core/ports"
)

const (
	path = "./store/productsdb.json"
)

var (
	byCategory map[string][]domain.Product
)

type Items struct {
	Products []domain.Product `json:"products"`
}

type inmemory struct {
	byCategory map[string][]domain.Product
}

func NewInMemory() ports.ProductRepository {
	return &inmemory{
		byCategory: byCategory,
	}
}

func (mem *inmemory) GetByCategory(category string, limit int64, lessThan *int64) ([]domain.Product, error) {
	return nil, nil
}

func init() {
	fmt.Println("trying to initialize inmemory db...")
	data, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	fmt.Println("data read ok!")

	fmt.Println("unmarshaling data...")
	var items Items
	jsonErr := json.Unmarshal(data, &items)
	if jsonErr != nil {
		panic(jsonErr)
	}
	fmt.Println("unmarshaling ok!")

	fmt.Printf("products: %d\n", len(items.Products))

	fmt.Println("structuring by category...")
	byCategory := make(map[string][]domain.Product)
	for _, product := range items.Products {
		_, ok := byCategory[product.Category]
		if !ok {
			byCategory[product.Category] = make([]domain.Product, 0)
		}
		byCategory[product.Category] = append(byCategory[product.Category], product)
	}
	fmt.Println("structuring ok!")

}
