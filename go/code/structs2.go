// +build OMIT

package main

import "fmt"

// START1 OMIT

// This `person` struct type has `name` and `age` fields.
type person struct {
	name string
	age  int
}

func main() {
	// This syntax creates a new struct.
	fmt.Println(person{"Bob", 20})

	// You can name the fields when initializing a struct.
	fmt.Println(person{name: "Alice", age: 30})

	// Omitted fields will be zero-valued.
	fmt.Println(person{name: "Fred"})

	// An `&` prefix yields a pointer to the struct.
	fmt.Println(&person{name: "Ann", age: 40})

	// Access struct fields with a dot.
	s := person{name: "Sean", age: 50}
	fmt.Println(s.name)
	fmt.Println(s)

	// STOP1 OMIT
	fmt.Println("\n\n\n\n")
	// START2 OMIT

	// You can also use dots with struct pointers - the
	// pointers are automatically dereferenced.
	sp := &s
	fmt.Println(sp.age)

	// Structs are mutable.
	sp.age = 51
	fmt.Println(sp.age)

	// STOP2 OMIT
}
