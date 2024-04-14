package products

import (
	"errors"
	"fmt"
	"strings"
)

type Products struct {
	Products []Product `json:"products"`
}

type Product struct {
	ID       string  `json:"id"`
	Name     string  `json:"name"`
	Category string  `json:"category"`
	Price    float64 `json:"price"`
}

// Create Products
func Create(id string, name string, category string, price float64) Product {
	return Product{id, name, category, price}
}

// Display all Products
func (products *Products) Display() {
	for _, item := range products.Products {
		fmt.Printf("%s : %s\n", item.ID, strings.ToUpper(item.Name))
	}
}

func (products *Products) GetProductByID(productID string) (Product, error) {
	for _, item := range products.Products {
		if item.ID == productID {
			return item, nil
		}
	}
	return Product{}, errors.New("INVALID PRODUCT ID")
}
