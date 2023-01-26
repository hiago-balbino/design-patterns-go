package main

import "fmt"

type House struct {
	windowsType, doorType string
	numFloor              int32
}

func (h House) String() {
	fmt.Printf(
		"this is a house with %s windows type, %s door type and %d floor(s).\n",
		h.windowsType,
		h.doorType,
		h.numFloor,
	)
}

type HouseBuilder interface {
	WithWindowsType(windowsType string) HouseBuilder
	WithDoorType(doorType string) HouseBuilder
	WithNumFloor(numFloor int32) HouseBuilder
	Build() House
}

type NormalBuilder struct {
	windowsType, doorType string
	numFloor              int32
}

func NewNormalBuilder() *NormalBuilder {
	return &NormalBuilder{}
}

func (n *NormalBuilder) WithWindowsType(windowsType string) HouseBuilder {
	n.windowsType = windowsType
	return n
}

func (n *NormalBuilder) WithDoorType(doorType string) HouseBuilder {
	n.doorType = doorType
	return n
}

func (n *NormalBuilder) WithNumFloor(numFloor int32) HouseBuilder {
	n.numFloor = numFloor
	return n
}

func (n *NormalBuilder) Build() House {
	return House{
		windowsType: n.windowsType,
		doorType:    n.doorType,
		numFloor:    n.numFloor,
	}
}

type IglooBuilder struct {
	windowsType, doorType string
	numFloor              int32
}

func NewIglooBuilder() *IglooBuilder {
	return &IglooBuilder{}
}

func (i *IglooBuilder) WithWindowsType(windowsType string) HouseBuilder {
	i.windowsType = windowsType
	return i
}

func (i *IglooBuilder) WithDoorType(doorType string) HouseBuilder {
	i.doorType = doorType
	return i
}

func (i *IglooBuilder) WithNumFloor(numFloor int32) HouseBuilder {
	i.numFloor = numFloor
	return i
}

func (i *IglooBuilder) Build() House {
	return House{
		windowsType: i.windowsType,
		doorType:    i.doorType,
		numFloor:    i.numFloor,
	}
}

func main() {
	normalHouse := NewNormalBuilder().
		WithWindowsType("Wooden").
		WithDoorType("Wooden").
		WithNumFloor(2).
		Build()
	normalHouse.String()

	iglooHouse := NewIglooBuilder().
		WithWindowsType("Snow").
		WithDoorType("Snow").
		WithNumFloor(1).
		Build()
	iglooHouse.String()
}
