package main

import "fmt"

// Position represents all role position types.
type Position int

const (
	DEVELOPER Position = iota
	MANAGER
)

// Employer is an interface to handle employees.
type Employer interface {
	// Work is a function that makes the employee works.
	Work()
}

// Employee is a base model with basic details of an employee.
type Employee struct {
	Name     string
	Position Position
}

// Developer is an employee that works as a developer.
type Developer struct {
	Employee
}

// Work is a function that makes the developer works.
func (d Developer) Work() {
	fmt.Printf("I'm %s and working as a Developer\n", d.Name)
}

// Manager is an employee that works as a manager.
type Manager struct {
	Employee
}

// Work is a function that makes the manager works.
func (m Manager) Work() {
	fmt.Printf("I'm %s and working as a Manager\n", m.Name)
}

// NewEmployee is a function that creates a new Employer based on roles.
func NewEmployee(name string, role Position) Employer {
	switch role {
	case DEVELOPER:
		return Developer{Employee: Employee{Name: name, Position: DEVELOPER}}
	case MANAGER:
		return Manager{Employee: Employee{Name: name, Position: MANAGER}}
	default:
		panic("unsupported role")
	}
}

func main() {
	developer := NewEmployee("Jhon", DEVELOPER)
	developer.Work()

	manager := NewEmployee("Maria", MANAGER)
	manager.Work()
}
