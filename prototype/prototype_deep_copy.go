package main

import "fmt"

// Address is a model that represents a person's address with basic details.
type Address struct {
	StreetAddress, City, Country string
}

// DeepCopy is a function of address struct to copy values to another one.
// The Address that will be copied is considered a Prototype.
func (a *Address) DeepCopy() *Address {
	return &Address{
		StreetAddress: a.StreetAddress,
		City:          a.City,
		Country:       a.Country,
	}
}

// Person is a model that represents a Person with basic details like name, address and friends.
type Person struct {
	Name    string
	Address *Address
	Friends []string
}

// DeepCopy is a function of a person's struct to copy values to another one.
// The Person that will be copied is considered a Prototype.
func (p *Person) DeepCopy() *Person {
	copiedPerson := *p // Copy the name value.
	copiedPerson.Address = p.Address.DeepCopy()
	copy(copiedPerson.Friends, p.Friends)

	return &copiedPerson
}

func main() {
	john := &Person{
		"John",
		&Address{StreetAddress: "123 London Rd", City: "London", Country: "UK"},
		[]string{"Chris", "Matt"},
	}

	jane := john.DeepCopy()                       // Here the Jane's person was created copying from the existent Jhon's prototype.
	jane.Name = "Jane"                            // Updating the name from prototype.
	jane.Address.StreetAddress = "321 Baker St"   // They live in the same Country and City, but different streets.
	jane.Friends = append(jane.Friends, "Angela") // Jane has the same friends as John, but is algo friend of Angela.

	fmt.Println(john, john.Address)
	fmt.Println(jane, jane.Address)
}
