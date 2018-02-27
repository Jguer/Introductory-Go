// +build OMIT

package main

import (
	"fmt"
)

// START1 OMIT
func main() {
	var x int
	funky := func(a, b int) int {
		return (a + b)
	}
	fmt.Println("Sum 3 and 5: ", funky(3, 5))

	fmt.Println("Sum 3 and 5: ",
		func(a, b int) int {
			return (a + b)
		}(3, 5)) // Called with args 3 and 5

	xBig := func() bool {
		return x > 10000 // References x declared in main start.
	}

	x = 99999
	fmt.Println("xBig:", xBig()) // true
	x = 1.3e3                    // This makes x == 13000
	fmt.Println("xBig:", xBig()) // false now.

}

// STOP1 OMIT
