package main

import "fmt"

// Size is a constant for all T-Shirt sizes.
type Size int

const (
	P Size = iota
	M
	G
	GG
)

// TShirt is a base model to represent a T-Shirt with basic details.
type TShirt struct {
	Color string
	Size  Size
}

type tshirtAction func(*TShirt)

// TShirtBuilder is a struct that contains a slice of actions to create a T-Shirt.
type TShirtBuilder struct {
	actions []tshirtAction
}

// WithColor is a function to set the T-Shirt color using builder actions.
func (tsb *TShirtBuilder) WithColor(color string) *TShirtBuilder {
	tsb.actions = append(tsb.actions, func(ts *TShirt) {
		ts.Color = color
	})
	return tsb
}

// WithSize is a function to set the T-Shirt size using builder actions.
func (tsb *TShirtBuilder) WithSize(size Size) *TShirtBuilder {
	tsb.actions = append(tsb.actions, func(ts *TShirt) {
		ts.Size = size
	})
	return tsb
}

// Build creates a T-Shirt based on TShirtBuilder actions.
func (tsb *TShirtBuilder) Build() TShirt {
	tshirt := TShirt{}
	for _, action := range tsb.actions {
		action(&tshirt)
	}
	return tshirt
}

func main() {
	builder := TShirtBuilder{}
	tshirt := builder.
		WithColor("Black").
		WithSize(G).
		Build()

	fmt.Println(tshirt)
}
