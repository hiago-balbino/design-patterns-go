package main

import "fmt"

// OfficeAddress is a model that represents an office's address with basic details.
type OfficeAddress struct {
	StreetAddress, City string
	Floor               int
}

// Employee is a model that represents a person that works at an office and needs to have the office address registered.
type Employee struct {
	Name          string
	OfficeAddress OfficeAddress
}

// DeepCopy is a function of an employee struct to copy values to another one.
// The Employee that will be copied is considered a Prototype.
func (e *Employee) DeepCopy() *Employee {
	copiedEmployee := *e
	return &copiedEmployee
}

// It is an employee that works at the main office.
var baseEmployee = Employee{
	Name: "",
	OfficeAddress: OfficeAddress{
		StreetAddress: "123 East Dr",
		City:          "London",
		Floor:         0,
	},
}

// It is an employee that works at the auxiliary office.
var auxEmployee = Employee{
	Name: "",
	OfficeAddress: OfficeAddress{
		StreetAddress: "66 West Dr",
		City:          "London",
		Floor:         0,
	},
}

// Utility function to copy prototype values and update the necessary attributes
// like employee name and office floor.
func newEmployee(prototype *Employee, name string, floor int) *Employee {
	copiedEmployee := prototype.DeepCopy()
	copiedEmployee.Name = name
	copiedEmployee.OfficeAddress.Floor = floor
	return copiedEmployee
}

// NewBaseEmployee creates a new employee that works at the main office.
func NewBaseEmployee(name string, floor int) *Employee {
	return newEmployee(&baseEmployee, name, floor)
}

// NewAuxEmployee creates a new employee that works at the auxiliary office.
func NewAuxEmployee(name string, floor int) *Employee {
	return newEmployee(&auxEmployee, name, floor)
}

func main() {
	jane := NewBaseEmployee("Jane", 100)
	john := NewAuxEmployee("John", 200)

	fmt.Println(jane)
	fmt.Println(john)
}
