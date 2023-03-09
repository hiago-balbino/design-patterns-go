package main

import "fmt"

// Component is an interface to represent a file/folder component on a computer.
type Component interface {
	search(keyword string)
}

// File is a model to represent a file on a computer.
type File struct {
	name string
}

func (f *File) search(keyword string) {
	fmt.Printf("Searching for keyword %s in file %s\n", keyword, f.getName())
}

func (f *File) getName() string {
	return f.name
}

// Folder is a model to represent a folder on a computer.
type Folder struct {
	name       string
	components []Component // It can be a file or a folder.
}

func (f *Folder) search(keyword string) {
	fmt.Printf("Searching recursively for keyword %s in folder %s\n", keyword, f.name)
	for _, composite := range f.components {
		composite.search(keyword)
	}
}

func (f *Folder) add(component Component) {
	f.components = append(f.components, component)
}

func main() {
	file1 := &File{name: "File1"}
	file2 := &File{name: "File2"}
	file3 := &File{name: "File3"}

	folder1 := &Folder{name: "Folder1"}
	folder1.add(file1)

	folder2 := &Folder{name: "Folder2"}
	folder2.add(file2)
	folder2.add(file3)
	folder2.add(folder1)

	folder2.search("rose")

	// Output
	// Searching recursively for keyword rose in folder Folder2
	// Searching for keyword rose in file File2
	// Searching for keyword rose in file File3
	// Searching recursively for keyword rose in folder Folder1
	// Searching for keyword rose in file File1
}
