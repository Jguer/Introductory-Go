// +build OMIT

package main

import (
	"fmt"
)

func main() {
	// START1 OMIT
	nums := []int{2, 3, 4}
	sum := 0
	for i, num := range nums {
		fmt.Printf("index: %d, value %d\n", i, num)
		sum += num
	}
	fmt.Println("sum:", sum)
	// STOP1 OMIT

	fmt.Print("\n\n\n\n\n\n\n")

	// START2 OMIT
	fmt.Printf("%+v\n", nums)
	for i, num := range nums {
		if num == 3 {
			fmt.Println("index:", i)
		}
	}
	// STOP2 OMIT
	fmt.Print("\n\n\n\n\n\n\n")

	// START3 OMIT
	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}
	// STOP3 OMIT
	fmt.Print("\n\n\n\n\n\n\n")

	// START4 OMIT
	for i, c := range "go" {
		fmt.Println(i, c)
		fmt.Printf("%c\n", c) //defer
	}
	// STOP4 OMIT
}
