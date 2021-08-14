package productrepo

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
	path        = "./store/productsdb.json"
	searchLimit = 5
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

func (mem *inmemory) GetByCategory(category string, lessThan *int64) ([]domain.Product, apierrors.ApiErr) {
	productCategory, ok := mem.byCategory[category]
	products := make([]domain.Product, 0)
	if !ok {
		return products, nil
	}

	for _, product := range productCategory {
		if len(products) == searchLimit {
			break
		}
		if lessThan != nil {
			if product.Price < *lessThan {
				products = append(products, product)
			} else {
				break
			}
			continue
		}
		products = append(products, product)
	}
	return products, nil
}

func init() {
	fmt.Println("trying to initialize products inmemory db...")
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
	byCategory = make(map[string][]domain.Product)
	for _, product := range items.Products {
		_, ok := byCategory[product.Category]
		if !ok {
			byCategory[product.Category] = make([]domain.Product, 0)
		}
		byCategory[product.Category] = append(byCategory[product.Category], product)
	}
	fmt.Println("structuring ok!")

	fmt.Println("sorting products categories by price")
	for category, _ := range byCategory {
		sort.Slice(byCategory[category], func(i, j int) bool {
			return byCategory[category][i].Price < byCategory[category][j].Price
		})
	}
	fmt.Println("sorting finished!")
}
