package cart

import (
	"github.com/chethanbhat/go-shopping-cart/products"
)

type Cart struct {
	Items []CartItem `json:"items"`
}

type CartItem struct {
	products.Product
	Qty            int     `json:"qty"`
	TotalBeforeTax float64 `json:"total_before_tax"`
	TotalAfterTax  float64 `json:"total_after_tax"`
}

type Order struct {
	Cart            `json:"order"`
	TotalOrderValue float64 `json:"total_order_value"`
}

// Create Cart
func CreateCart() Cart {
	return Cart{}
}

// Add to Cart
func AddToCart(o *Cart, p products.Product, taxRate float64) {
	// Initialize the Items slice if it's nil
	if o.Items == nil {
		o.Items = make([]CartItem, 0)
	}

	// Check if the item already exists in the cart
	exists := false
	for idx, item := range o.Items {
		if item.Product.Name == p.Name && item.Product.Category == p.Category {
			// Item already exists, increase quantity and update total
			o.Items[idx].Qty += 1
			o.Items[idx].TotalBeforeTax = float64(o.Items[idx].Qty) * p.Price
			o.Items[idx].TotalAfterTax = float64(o.Items[idx].Qty) * p.Price * (1 + taxRate)
			exists = true
			break
		}
	}

	// If the item is new, add it to the cart
	if !exists {
		qty := 1
		newItem := CartItem{
			p,
			qty,
			float64(qty) * p.Price,
			float64(qty) * p.Price * (1 + taxRate),
		}

		o.Items = append(o.Items, newItem)
	}

}

// Remove from Cart
func RemoveFromCart(o *Cart, p products.Product) {
	if o.Items == nil {
		return
	}

	for idx, item := range o.Items {
		if item.Product.Name == p.Name && item.Product.Category == p.Category {
			// Remove the item from the slice using slicing and append
			o.Items = append(o.Items[:idx], o.Items[idx+1:]...)
			break
		}
	}

}

// Get Cart Total
func GetCartTotal(o *Cart) float64 {
	total := 0.0
	for _, item := range o.Items {
		total += item.TotalAfterTax
	}
	return total
}
