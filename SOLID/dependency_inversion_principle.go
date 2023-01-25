// dependency_inversion_principle is a file to demonstrate what is the idea behind the Dependency Inversion Principle incluing some example.
// The file simulates a Bakery.
package main

import "fmt"

// CookOrder is a struct to represent the model of a cook order to prepare.
type CookOrder struct{}

// Prepare is a function that implement the preparation of a cook order.
func (c CookOrder) Prepare(order string) {
	fmt.Printf("cooked a %s", order)
}

// Kitchen is an interface that contains the function that handle with orders.
type Kitchen interface {
	// Cook is a function that cook a order received through parameter.
	Cook(order string)
}

// Baker is a struct that represent a baker person and implement the Kitchen interface to handle with orders.
type Baker struct {
	name string
}

// Cook is a function that cook a order received through parameter.
func (b Baker) Cook(order string) {
	fmt.Printf("cooked a %s", order)
}

//type Bakery struct {
//	cookOrder CookOrder
//}

//func NewBakery(cookOrder CookOrder) Bakery {
//	return Bakery{
//		cookOrder: cookOrder,
//	}
//}

type Bakery struct {
	kitchen Kitchen
}

func NewBakery(kitchen Kitchen) Bakery {
	return Bakery{
		kitchen: kitchen,
	}
}

func main() {
	order := "chocolate cake"

	//	bakery := NewBakery(CookOrder{})
	//	bakery.cookOrder.Prepare(order)

	baker := Baker{name: "Jhon"}
	bakery := NewBakery(baker)
	bakery.kitchen.Cook(order)
}
