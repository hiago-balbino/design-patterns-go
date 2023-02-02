package main

import "fmt"

// Person is a model to represent a Person with basic details(name, age, eye count).
type Person struct {
	Name     string
	Age      int
	EyeCount int // Just to be used as an example created with the default value.
}

// NewPerson is a constructor/factory to create a new instance of Person.
// This factory validates and throws a panic when the age is less than zero.
// The EyeCount field is filled with default value=2.
func NewPerson(name string, age int) Person {
	if age < 0 {
		panic("some exception")
	}

	return Person{
		Name:     name,
		Age:      age,
		EyeCount: 2, // Default value.
	}
}

func main() {
	// Initialize directly.
	personDirectly := Person{Name: "Jane", Age: 21}
	fmt.Println("created directly: ", personDirectly)

	// Initialize using constructor/factory.
	personFactory := NewPerson("Jane", 21)
	fmt.Println("created using constructor/factory: ", personFactory)
}
