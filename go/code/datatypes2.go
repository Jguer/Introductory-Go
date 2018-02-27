// +build OMIT

package main

import (
	"fmt"
)

var (
	i = 10
)

func main() {
	i := 42
	fmt.Printf("\tValue of i: %d\n", i)

	fmt.Println("1+1 =", 1+1)
	fmt.Println("7/3 =", 7/3)
	fmt.Println("7%3 =", 7%3)
	fmt.Println("7*3 =", 7*3)
	fmt.Println("7.0/3.0 =", 7.0/3.0)

	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)
}
