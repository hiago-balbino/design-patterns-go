package main

import "fmt"

type Size int

const (
	P Size = iota
	M
	G
	GG
)

type TShirt struct {
	Color string
	Size  Size
}

type tshirtMod func(*TShirt)

type TShirtBuilder struct {
	actions []tshirtMod
}

func (tsb *TShirtBuilder) WithColor(color string) *TShirtBuilder {
	tsb.actions = append(tsb.actions, func(ts *TShirt) {
		ts.Color = color
	})
	return tsb
}

func (tsb *TShirtBuilder) WithSize(size Size) *TShirtBuilder {
	tsb.actions = append(tsb.actions, func(ts *TShirt) {
		ts.Size = size
	})
	return tsb
}

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
