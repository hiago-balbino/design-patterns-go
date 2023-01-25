// dependency_inversion_principle is a file to demonstrate what is the idea behind the Dependency Inversion Principle including some examples.
// The file simulates a Bakery.
package main

import "fmt"

// CookOrder is a struct to represent the model of a cooking order to prepare.
type CookOrder struct{}

// Prepare is a function that implements the preparation of a cooking order.
func (c CookOrder) Prepare(order string) {
	fmt.Printf("cooked a %s", order)
}

// Kitchen is an interface that contains the function that handles orders.
type Kitchen interface {
	// Cook is a function that cooks an order received through a parameter.
	Cook(order string)
}

// Baker is a struct that represents a baker person and implements the Kitchen interface to handle orders.
type Baker struct {
	name string
}

// Cook is a function that cooks an order received through a parameter.
func (b Baker) Cook(order string) {
	fmt.Printf("cooked a %s", order)
}

// Bakery is a model to represent a bakery that cooks orders.
//type Bakery struct {
//	cookOrder CookOrder
//}

// NewBakery is a constructor to create a new instance of Bakery to cook orders.
// This breaks the principle given the Bakery is depending on a concrete struct instead of an abstraction.
//func NewBakery(cookOrder CookOrder) Bakery {
//	return Bakery{
//		cookOrder: cookOrder,
//	}
//}

// Bakery is a model to represent a bakery that cooks orders in a kitchen.
type Bakery struct {
	kitchen Kitchen
}

// NewBakery is a constructor to create a new instance of a Bakery to cook orders in a Kitchen.
// This is a good way and does not break the principle given the Bakery is depending on a Kitchen abstraction instead of a concrete struct.
func NewBakery(kitchen Kitchen) Bakery {
	return Bakery{
		kitchen: kitchen,
	}
}

func main() {
	order := "chocolate cake"

	// Using the constructor that breaks the principle.
	//	bakery := NewBakery(CookOrder{})
	//	bakery.cookOrder.Prepare(order)

	// Using the constructor that does not break the principle.
	baker := Baker{name: "Jhon"}
	bakery := NewBakery(baker)
	bakery.kitchen.Cook(order)
}
