// open_closed_principle is a file to demonstrate what is the idea behind the Open-Closed Principle including some examples.
// The file simulates a sales Website that needs to filter the products for some attributes.
package main

import (
	"fmt"
)

// Color is the type for the product colors that will be used as constant.
type Color int

const (
	red Color = iota
	green
	blue
)

// Size is the type for the product sizes that will be used as constant.
type Size int

const (
	small Size = iota
	medium
	large
)

// Product is a struct to represent the product model with your attributes.
type Product struct {
	name  string
	color Color
	size  Size
}

// Filter is a struct that will implement the function to filter products given the parameter of filter.
// This way the principle will be broken given if some other filter needs to be added,
// the implemented functions also need to be updated,
// what means the function IS CLOSED FOR EXTENSION, BUT OPEN TO MODIFICATION.
type Filter struct{}

// ByColorAndSize is a function that filters a list of products given the parameters of color and size.
// The function is a wrong implementation that breaks the Open-Closed Principle given that if any other filter
// needs to be made, and this function needs to be updated.
func (f Filter) ByColorAndSize(products []Product, color Color, size Size) []Product {
	result := make([]Product, 0)

	for _, product := range products {
		if product.color == color && product.size == size {
			result = append(result, product)
		}
	}

	return result
}

// Filterer is an interface that represents the abstraction for filtering.
// All types of filters needed can implement this interface.
// So the code is OPEN FOR EXTENSION AND CLOSED FOR MODIFICATION.
type Filterer interface {
	// Filter is a function that will be implemented by concrete struct for any specific filter type.
	Filter([]Product) []Product
}

// ColorFilter is a struct that will implement the Filterer interface to filter products by color.
type ColorFilter struct {
	color Color
}

// Filter is an implementation of the Filterer interface to filter products by color.
func (c ColorFilter) Filter(products []Product) []Product {
	result := make([]Product, 0)

	for _, product := range products {
		if product.color == c.color {
			result = append(result, product)
		}
	}

	return result
}

// ColorAndSizeFilter is a struct that will implement the Filterer interface to filter products by color and size.
type ColorAndSizeFilter struct {
	color Color
	size  Size
}

// Filter is an implementation of the Filterer interface to filter products by color and size.
func (c ColorAndSizeFilter) Filter(products []Product) []Product {
	result := make([]Product, 0)

	for _, product := range products {
		if product.color == c.color && product.size == c.size {
			result = append(result, product)
		}
	}

	return result
}

func main() {
	apple := Product{name: "Apple", color: green, size: small}
	tree := Product{name: "Tree", color: green, size: large}
	house := Product{name: "House", color: blue, size: small}
	products := []Product{apple, tree, house}

	// Using the wrong implementation.
	filter := Filter{}
	productsFiltered := filter.ByColorAndSize(products, green, large)

	fmt.Println("Large and Green products (wrong implementation)")
	for _, productFiltered := range productsFiltered {
		fmt.Printf(" - %s is large and green\n", productFiltered.name)
	}

	// Using the right implementation to filter by color.
	colorFilterer := ColorFilter{color: green}
	productsColorFiltered := colorFilterer.Filter(products)

	fmt.Println("Green products (right implementation)")
	for _, productColorFiltered := range productsColorFiltered {
		fmt.Printf(" - %s is green\n", productColorFiltered.name)
	}

	// Using the right implementation to filter by color and size.
	colorAndSizeFilterer := ColorAndSizeFilter{color: green, size: large}
	productsColorAndSizeFiltered := colorAndSizeFilterer.Filter(products)

	fmt.Println("Large and Green products (right implementation)")
	for _, productColorAndSizeFiltered := range productsColorAndSizeFiltered {
		fmt.Printf(" - %s is large and green\n", productColorAndSizeFiltered.name)
	}
}
