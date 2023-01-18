// interface_segregation_principle is a file to demonstrate what is the idea behind the Interface segregation principle including some examples.
// The file simulates a Printer.
package main

import "fmt"

// Document is a struct to represent the model that will be used to print some data.
type Document struct{}

// MachinePrinter is a "Bad Example" interface that contains output methods for documents.
// This approach has some problems, given that neither all printers have all output types (Print, Scan, etc)
// So some concrete struct will need to implement the function that is not necessary/supported.
type MachinePrinter interface {
	// Print is a function that prints a document.
	Print(document Document)
	// Scan is a function for scanning a document.
	Scan(document Document)
}

// OldFashionedPrinter is a struct that implements the MachinePrinter interface to print documents.
type OldFashionedPrinter struct{}

// Print is a function that prints a document.
func (o OldFashionedPrinter) Print(document Document) {
	fmt.Println("[OldFashionedPrinter-Print] do something with successful here")
}

// Scan is a function for scanning a document.
func (o OldFashionedPrinter) Scan(document Document) {
	panic("[OldFashionedPrinter-Scan] operation not supported")
}

// Printer is a "Good Example" interface that has a unique way of outputting a document.
// This is a better way given that just the specific machine that prints a document will implement the interface.
// If the machine need another way of output, it can implement more interfaces, this is the Interface Segregation.
type Printer interface {
	// Print is a function that prints a document.
	Print(document Document)
}

// Scanner is a "Good Example" interface that contains output methods for scanning a document.
type Scanner interface {
	// Scan is a function for scanning a document.
	Scan(document Document)
}

// MyPrinter is a struct to implement the Printer interface for printing.
type MyPrinter struct{}

// Print is a function that prints a document.
func (m MyPrinter) Print(document Document) {
	fmt.Println("[MyPrinter-Print] do something with successful here")
}

// Photocopier is a struct to implement the Printer and Scanner interface for printing and scanning.
type Photocopier struct{}

// Scan is a function for scanning a document.
func (p Photocopier) Scan(document Document) {
	fmt.Println("[Photocopier-Scan] do something with successful here")
}

// Print is a function that prints a document.
func (p Photocopier) Print(document Document) {
	fmt.Println("[Photocopier-Print] do something with successful here")
}

func main() {
	document := Document{}

	// Bad approach
	//	oldFashionedPrinter := OldFashionedPrinter{}
	//	oldFashionedPrinter.Print(document)
	//	oldFashionedPrinter.Scan(document)

	// Good approach
	myPrinter := MyPrinter{}
	myPrinter.Print(document)

	// Good approach
	photocopier := Photocopier{}
	photocopier.Print(document)
	photocopier.Scan(document)
}
