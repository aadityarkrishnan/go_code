// Go program to illustrate how to
// create methods of the same name
package main

import "fmt"

// Creating structures
type student struct {
	name string
	branch string
}

type teacher struct {
	language string
	marks int
}

// Same name methods, but with
// different type of receivers
func (s student) show() {

	fmt.Println("Name of the Student:", s.name)
	fmt.Println("Branch: ", s.branch)
}

func (t teacher) show() {

	fmt.Println("Language:", t.language)
	fmt.Println("Student Marks: ", t.marks)
}

// Main function
func main() {

	// Initializing values
	// of the structures
	val1 := student{"Rohit", "EEE"}

	val2 := teacher{"Java", 50}

	// Calling the methods
	val1.show()
	val2.show()
}
