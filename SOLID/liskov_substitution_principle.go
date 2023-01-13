// liskov_substitution_principle is a file to demonstrate what is the idea behind the Liskov Substitution Principle including some examples.
// The file simulates Rectangle representations.
// In mathematics, a Square is a Rectangle, indeed it is a specialization of a rectangle, but it makes you want to model this with inheritance. However if in code you made Square derive from Rectangle, then a Square can be used anywhere you expect a Rectangle. This makes you break the principle.
// Go(lang) does not have inheritance but we can explore the principle in a similar way using Composition. Let's say you have SetWidth and SetHeight methods on your Rectangle base struct, this seems perfectly logical. However, if your Rectangle reference pointed to a Square, then SetWidth and SetHeight do not make sense because setting one would change the other to match it. In this case, Square fails the Liskov Substitution with Rectangle and the abstraction of having Square "inherit" from Rectangle is a bad one.
package main

// Area is an interface to calculate the area of the Rectangle.
type Area interface {
	// SetWidth sets the width of the area.
	SetWidth(width int)
	// SetHeight sets the height of the area.
	SetHeight(height int)
	// CalculateArea is the function to calculate the area given the width and height.
	CalculateArea() int
}

// Rectangle is a model to represent the base rectangle.
type Rectangle struct {
	width, height int
}

// SetWidth sets the width of the area.
func (r *Rectangle) SetWidth(width int) {
	r.width = width
}

// SetHeight sets the height of the area.
func (r *Rectangle) SetHeight(height int) {
	r.height = height
}

// CalculateArea is the function to calculate the area given the width and height.
func (r Rectangle) CalculateArea() int {
	return r.width * r.height
}

// Square is a model to represent a rectangle using composition.
type Square struct {
	Rectangle
}

// SetWidth sets the width of the area.
func (s *Square) SetWidth(width int) {
	s.width = width
	s.height = width
}

// SetHeight sets the height of the area.
func (s *Square) SetHeight(height int) {
	s.height = height
	s.width = height
}

// CalculateArea is the function to calculate the area given the width and height.
func (s Square) CalculateArea() int {
	return s.width * s.height
}

func main() {
	var width = 5
	var height = 3
	var expectedArea = width * height

	rectangle := Rectangle{}
	rectangle.SetWidth(width)
	rectangle.SetHeight(height)
	rectangleArea := rectangle.CalculateArea()
	if rectangleArea != expectedArea {
		panic("wrong rectangle area")
	}

	// The LSP will be broken here, given the Square has the same width and height, so this is updated in the Set functions
	// and will affect the result.
	square := Square{}
	square.SetWidth(width)
	square.SetHeight(height)
	squareArea := square.CalculateArea()
	if squareArea != expectedArea {
		panic("wrong square area")
	}
}
