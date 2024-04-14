package main

import (
	"fmt"
	"strconv"

	"github.com/chethanbhat/go-shopping-cart/cart"
	"github.com/chethanbhat/go-shopping-cart/errorhandling"
	"github.com/chethanbhat/go-shopping-cart/fileoperations"
	"github.com/chethanbhat/go-shopping-cart/products"
	"github.com/chethanbhat/go-shopping-cart/taxes"
	"github.com/google/uuid"
)

var PRODUCTS_DATA = "products.json"
var TAX_DATA = "taxes.json"

func main() {
	fmt.Println("Welcome to Shopping Cart")

	// Read the products data from the file
	var allProducts products.Products
	err := fileoperations.ReadJSONData(PRODUCTS_DATA, &allProducts)
	errorhandling.Manage(err)

	// Get tax information
	var taxData taxes.TaxData
	err = fileoperations.ReadJSONData(TAX_DATA, &taxData)
	errorhandling.Manage(err)

	cart := cart.CreateCart()
	taxes := taxData.TaxRates
	option := 0

	for option != 5 {
		showOptions()
		var value string
		fileoperations.ReadUserInput("\nSelect any option", &value)

		option, err = strconv.Atoi(value)

		if err != nil {
			fmt.Println("Please select a valid option")
			continue
		}

		switch option {
		case 1:
			addProducts(allProducts, &cart, taxes)
		case 2:
			removeProducts(allProducts, &cart)
		case 3:
			viewcart(&cart)
			continue
		case 4:
			checkout(&cart)
			return
		case 5:
			showExitMessage()
		default:
			fmt.Println("Invalid option")
		}
	}

}

func displayAllProducts(p products.Products) {
	p.Display()
}

func showOptions() {
	fmt.Println("Select an option")
	fmt.Println("1. Add Products")
	fmt.Println("2. Remove Product")
	fmt.Println("3. View Cart")
	fmt.Println("4. Checkout")
	fmt.Println("5. Exit")
}

func addProducts(p products.Products, c *cart.Cart, t map[string]float64) {
	displayAllProducts(p)
	selectedProductID := ""
	for !checkforValidProductID(p.Products, selectedProductID) {
		fileoperations.ReadUserInput("Enter Product ID", &selectedProductID)
	}
	selectedProduct, err := p.GetProductByID(selectedProductID)
	if err != nil {
		errorhandling.Manage(err)
		return
	}
	taxRate := t[selectedProduct.Category]
	c.AddToCart(selectedProduct, taxRate)

}

func removeProducts(p products.Products, c *cart.Cart) {
	displayAllProducts(p)
	selectedProductID := ""
	for !checkforValidProductID(p.Products, selectedProductID) {
		fileoperations.ReadUserInput("Enter Product ID", &selectedProductID)
	}
	selectedProduct, err := p.GetProductByID(selectedProductID)
	if err != nil {
		errorhandling.Manage(err)
		return
	}
	c.RemoveFromCart(selectedProduct)
}

func viewcart(c *cart.Cart) {
	if len(c.Items) == 0 {
		fmt.Println("Your cart is empty")
		return
	}
	fmt.Println("You have added following items to the cart", len(c.Items))
	for index, item := range c.Items {
		fmt.Println(index+1, "> ", item.Name, " x ", item.Qty)
	}
}

func checkout(c *cart.Cart) {
	cartTotal := c.GetCartTotal()
	fmt.Println("Your Cart Total: ", cartTotal)
	fileID := uuid.New()
	filename := fmt.Sprintf("order-%s.json", &fileID)
	var order = cart.Order{}
	order.Cart = *c
	order.TotalOrderValue = cartTotal
	err := fileoperations.WriteJSONData(filename, &order)
	errorhandling.Manage(err)

	fmt.Println("Order has been saved to file ", filename)
	showExitMessage()
}

func showExitMessage() {
	fmt.Println("Thank you. Visit us again !")
}

func checkforValidProductID(products []products.Product, productID string) bool {
	for _, item := range products {
		if item.ID == productID {
			return true
		}
	}
	return false
}
