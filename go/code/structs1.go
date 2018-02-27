// +build OMIT

package main

import "fmt"

// START1 OMIT

type newstringtype string

// Vertex describes a point in space.
type Vertex struct {
	X int
	Y int
}

func main() {
	var myString newstringtype
	myString = "it works"

	fmt.Println(myString)

	myVertex := Vertex{1, 2} //Declare implicit
	fmt.Println(myVertex)

	myNewVertex := Vertex{X: 3, Y: 4} //Declare explicit
	fmt.Println(myNewVertex)
}

// STOP1 OMIT
