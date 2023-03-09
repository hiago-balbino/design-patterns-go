package main

import "fmt"

// Printer is an interface to handle print files in different printers.
type Printer interface {
	// PrintFile prints the file given the printer model.
	PrintFile()
}

// Epson is a model of printer.
type Epson struct{}

// PrintFile prints the file given the Espon printer model.
func (p *Epson) PrintFile() {
	fmt.Println("Printing by an Epson Printer")
}

// Hp is a model of printer.
type Hp struct{}

// PrintFile prints the file given the Hp printer model.
func (p *Hp) PrintFile() {
	fmt.Println("Printing by a HP Printer")
}

// Computer is an interface representing a computer that prints some output using any printer model.
type Computer interface {
	// SetPrinter sets the model of the printer for the computer.
	SetPrinter(printer Printer)
	// Print prints any file using any printer connected to the computer.
	Print()
}

// Mac is a model of computer.
type Mac struct {
	printer Printer
}

// SetPrinter sets the model of the printer for the Mac computer.
func (m *Mac) SetPrinter(printer Printer) {
	m.printer = printer
}

// Print prints any file using any printer connected to the Mac computer.
func (m *Mac) Print() {
	fmt.Println("Print request for mac")
	m.printer.PrintFile()
}

// Windows is a model of computer.
type Windows struct {
	printer Printer
}

// SetPrinter sets the model of the printer for the Windows computer.
func (w *Windows) SetPrinter(printer Printer) {
	w.printer = printer
}

// Print prints any file using any printer connected to the Windows computer.
func (w *Windows) Print() {
	fmt.Println("Print request for windows")
	w.printer.PrintFile()
}

func main() {
	hpPrinter := &Hp{}
	epsonPrinter := &Epson{}

	macComputer := &Mac{}
	macComputer.SetPrinter(hpPrinter)
	macComputer.Print()
	fmt.Println()
	macComputer.SetPrinter(epsonPrinter)
	macComputer.Print()
	fmt.Println()

	windowsComputer := &Windows{}
	windowsComputer.SetPrinter(hpPrinter)
	windowsComputer.Print()
	fmt.Println()
	windowsComputer.SetPrinter(epsonPrinter)
	windowsComputer.Print()
}
