package main

import "fmt"

// Person is a model to represent a Person with address and job details.
type Person struct {
	// Address details.
	StreetAddress, Postcode, City string
	// Job details.
	CompanyName, Position string
	AnnualIncome          int
}

// PersonBuilder is a builder with a Person to be built.
type PersonBuilder struct {
	person *Person
}

// NewPersonBuilder is a constructor to create a new instance of PersonBuilder.
func NewPersonBuilder() *PersonBuilder {
	return &PersonBuilder{&Person{}}
}

// Build create a Person given the attributes set.
func (pb *PersonBuilder) Build() Person {
	return *pb.person
}

// Lives is a constructor to create a new instance of PersonAddressBuilder using PersonBuilder as a composite.
func (pb *PersonBuilder) Lives() *PersonAddressBuilder {
	return &PersonAddressBuilder{*pb}
}

// Works is a constructor to create a new instance of PersonJobBuilder using PersonBuilder as a composite.
func (pb *PersonBuilder) Works() *PersonJobBuilder {
	return &PersonJobBuilder{*pb}
}

// PersonAddressBuilder is a builder that uses PersonBuilder as a composite to create a Person Address details.
type PersonAddressBuilder struct {
	PersonBuilder
}

// At is a function to set the street address of the Person.
func (pab *PersonAddressBuilder) At(streetAddress string) *PersonAddressBuilder {
	pab.person.StreetAddress = streetAddress
	return pab
}

// In is a function to set the city of the Person.
func (pab *PersonAddressBuilder) In(city string) *PersonAddressBuilder {
	pab.person.City = city
	return pab
}

// WithPostcode is a function to set the postcode of the Person.
func (pab *PersonAddressBuilder) WithPostcode(postcode string) *PersonAddressBuilder {
	pab.person.Postcode = postcode
	return pab
}

// PersonJobBuilder is a builder that uses PersonBuilder as a composite to create a Person Job details.
type PersonJobBuilder struct {
	PersonBuilder
}

// At is a function to set the job company name to the Person.
func (pjb *PersonJobBuilder) At(companyName string) *PersonJobBuilder {
	pjb.person.CompanyName = companyName
	return pjb
}

// AsA is a function to set the job position to the Person.
func (pjb *PersonJobBuilder) AsA(position string) *PersonJobBuilder {
	pjb.person.Position = position
	return pjb
}

// Earning is a function to set the job earning to the Person.
func (pjb *PersonJobBuilder) Earning(annualIncome int) *PersonJobBuilder {
	pjb.person.AnnualIncome = annualIncome
	return pjb
}

func main() {
	person := NewPersonBuilder().
		Lives().
		At("123 Avenue").
		In("EUA").
		WithPostcode("AV321EA").
		Works().
		At("Tech Company").
		AsA("Software Engineer").
		Earning(90500).
		Build()

	fmt.Println(person)
}
