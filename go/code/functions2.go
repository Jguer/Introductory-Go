// +build OMIT

package main

import (
	"fmt"
)

func vals() (int, int) {
	return 3, 7
}

func main() {
	a, b := vals()
	fmt.Println("Double return:", a, b)

	_, c := vals()
	fmt.Println("One return:", c)
}
