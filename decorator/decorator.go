package main

import "fmt"

// Pizza is an interface that represents a pizza and returns its price.
type Pizza interface {
	getPrice() int
}

// VeggieMania is a struct that represents a veggie pizza.
type VeggieMania struct{}

func (p *VeggieMania) getPrice() int {
	return 15
}

// TomatoTopping is a struct that represents a pizza with tomato topping.
type TomatoTopping struct {
	pizza Pizza
}

func (t *TomatoTopping) getPrice() int {
	pizzaPrice := t.pizza.getPrice()
	return pizzaPrice + 7
}

// CheeseTopping is a struct that represents a pizza with cheese topping.
type CheeseTopping struct {
	pizza Pizza
}

func (c *CheeseTopping) getPrice() int {
	pizzaPrice := c.pizza.getPrice()
	return pizzaPrice + 10
}

func main() {
	pizza := &VeggieMania{}
	// Add cheese topping.
	pizzaWithCheese := &CheeseTopping{pizza: pizza}
	// Add tomato topping.
	pizzaWithCheeseAndTomato := &TomatoTopping{pizza: pizzaWithCheese}

	fmt.Printf("Price of VeggieMania with tomato and cheese topping is %d.\n", pizzaWithCheeseAndTomato.getPrice())
}
