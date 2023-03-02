package main

import "fmt"

type Printer interface {
	PrintFile()
}

type Epson struct{}

func (p *Epson) PrintFile() {
	fmt.Println("Printing by an Epson Printer")
}

type Hp struct{}

func (p *Hp) PrintFile() {
	fmt.Println("Printing by a HP Printer")
}

type Computer interface {
	SetPrinter(printer Printer)
	Print()
}

type Mac struct {
	printer Printer
}

func (m *Mac) SetPrinter(printer Printer) {
	m.printer = printer
}

func (m *Mac) Print() {
	fmt.Println("Print request for mac")
	m.printer.PrintFile()
}

type Windows struct {
	printer Printer
}

func (w *Windows) SetPrinter(printer Printer) {
	w.printer = printer
}

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
