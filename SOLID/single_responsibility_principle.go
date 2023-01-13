// single_responsibility_principle is a file to demonstrate what is the idea behind the Single Responsibility Principle
// including good and bad examples.
// The file simulates a Journal that Adds and Prints the entries(news).
package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

// Journal is a struct to represent the model with a list of entries.
type Journal struct {
	entries []string
}

// String is a override implementation of Stringer to facilitate the output printer.
func (j Journal) String() string {
	return strings.Join(j.entries, "\n")
}

// AddEntry put a new text entry in Journal.
func (j Journal) AddEntry(text string) {
	j.entries = append(j.entries, text)
}

// SaveFile is a function that breaks the Single Responsibility Principle given that it puts a new
// responsibility for the Journal model(Add and Persist file).
// The persistence responsibility can be implemented in a new file/pkg persistence that is responsible for saving the
// Journal anywhere you need(filepath, database, memory).
func (j Journal) SaveFile(filename string) {
	var lineSeparator = "\n"
	_ = ioutil.WriteFile(filename, []byte(strings.Join(j.entries, lineSeparator)), 0644)
}

func main() {
	journal := Journal{
		entries: []string{
			"I had a awesome day.",
		},
	}
	journal.entries = append(journal.entries, "I deserve a cold beer.")

	fmt.Println(journal.String())

	// This calls the function that broke the principle.
	journal.SaveFile("../journal.txt")
}
