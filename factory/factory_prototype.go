package main

import "fmt"

type Position int

const (
	DEVELOPER Position = iota
	MANAGER
)

type Employer interface {
	Work()
}

type Employee struct {
	Name     string
	Position Position
}

type Developer struct {
	Employee
}

func (d Developer) Work() {
	fmt.Printf("I'm %s and working as a Developer\n", d.Name)
}

type Manager struct {
	Employee
}

func (m Manager) Work() {
	fmt.Printf("I'm %s and working as a Manager\n", m.Name)
}

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
