package main

import "fmt"

// House is a model to represent the house to be built.
type House struct {
	windowsType, doorsType string
	numFloors              int32
}

// String implements a Stringer interface to print the output details.
func (h House) String() {
	fmt.Printf(
		"this is a house with %s windows type, %s doors type and %d floor(s).\n",
		h.windowsType,
		h.doorsType,
		h.numFloors,
	)
}

// HouseBuilder is an interface with functions to create a House in some steps.
type HouseBuilder interface {
	// WithWindowsType is a function to set the windows type of the house.
	WithWindowsType(windowsType string) HouseBuilder
	// WithDoorsType is a function to set the doors type of the house.
	WithDoorsType(doorsType string) HouseBuilder
	// WithNumFloors is a function to set the floor numbers of the house.
	WithNumFloors(numFloors int32) HouseBuilder
	// Build is a function to build a house given the attributes.
	Build() House
}

// NormalBuilder is a builder that implements the HouseBuilder interface to create a normal house.
type NormalBuilder struct {
	windowsType, doorsType string
	numFloors              int32
}

// NewNormalBuilder is a constructor to create a new instance of NormalBuilder.
func NewNormalBuilder() *NormalBuilder {
	return &NormalBuilder{}
}

// WithWindowsType is a function to set the windows type of the house.
func (n *NormalBuilder) WithWindowsType(windowsType string) HouseBuilder {
	n.windowsType = windowsType
	return n
}

// WithDoorsType is a function to set the doors type of the house.
func (n *NormalBuilder) WithDoorsType(doorsType string) HouseBuilder {
	n.doorsType = doorsType
	return n
}

// WithNumFloors is a function to set the floor numbers of the house.
func (n *NormalBuilder) WithNumFloors(numFloors int32) HouseBuilder {
	n.numFloors = numFloors
	return n
}

// Build is a function to build a house given the attributes.
func (n *NormalBuilder) Build() House {
	return House{
		windowsType: n.windowsType,
		doorsType:   n.doorsType,
		numFloors:   n.numFloors,
	}
}

// IglooBuilder is a builder that implements the HouseBuilder interface to create an igloo house.
type IglooBuilder struct {
	windowsType, doorsType string
	numFloors              int32
}

// NewIglooBuilder is a constructor to create a new instance of IglooBuilder.
func NewIglooBuilder() *IglooBuilder {
	return &IglooBuilder{}
}

// WithWindowsType is a function to set the windows type of the house.
func (i *IglooBuilder) WithWindowsType(windowsType string) HouseBuilder {
	i.windowsType = windowsType
	return i
}

// WithDoorsType is a function to set the doors type of the house.
func (i *IglooBuilder) WithDoorsType(doorsType string) HouseBuilder {
	i.doorsType = doorsType
	return i
}

// WithNumFloors is a function to set the floor numbers of the house.
func (i *IglooBuilder) WithNumFloors(numFloors int32) HouseBuilder {
	i.numFloors = numFloors
	return i
}

// Build is a function to build a house given the attributes.
func (i *IglooBuilder) Build() House {
	return House{
		windowsType: i.windowsType,
		doorsType:   i.doorsType,
		numFloors:   i.numFloors,
	}
}

func main() {
	normalHouse := NewNormalBuilder().
		WithWindowsType("Wooden").
		WithDoorsType("Wooden").
		WithNumFloors(2).
		Build()
	normalHouse.String()

	iglooHouse := NewIglooBuilder().
		WithWindowsType("Snow").
		WithDoorsType("Snow").
		WithNumFloors(1).
		Build()
	iglooHouse.String()
}
