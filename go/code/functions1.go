// +build OMIT

package main

import (
	"fmt"
)

// START1 OMIT
func plus(a int, b int) int {
	return a + b
}

func morePlus(a, b, c int) int {
	return a + b + c
}
func main() {
	res := plus(1, 2)
	fmt.Println("1+2 =", res)
	fmt.Println("1+2 =", plus(1, 2))

	res = morePlus(1, 2, 3)
	fmt.Println("1+2+3 =", res)
}
// STOP1 OMIT
