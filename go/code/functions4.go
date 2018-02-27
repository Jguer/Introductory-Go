// +build OMIT

package main

import (
	"fmt"
)

// START1 OMIT
func valchanger(a int, b int) int {
	a = a + b
	return a
}
func main() {
	a := 2
	b := 3

	a = valchanger(a, b)
	fmt.Println("Post-Function: ", a)
}

// STOP1 OMIT
