// +build OMIT

package main

import (
	"fmt"
	"strings"
)

func main() {
	// START1 OMIT
	fmt.Println("Method one")
	wordSlice := [...]string{"Gate", "Comet", "Pizza"} // Create Array
	fmt.Println(wordSlice)

	for _, w := range wordSlice {
		for _, c := range w {
			if c == 'e' {
				fmt.Println("Found one!", string(w[0]))
			}
		}
	}
	// STOP1 OMIT

	alternativeone()
	alternativetwo()
}

func alternativeone() {
	// START2 OMIT
	fmt.Println("Method two")
	wordSlice := []string{"Gate", "Comet", "Pizza"} // Create Slice
	fmt.Println(wordSlice)

	for _, w := range wordSlice {
		for _, c := range w {
			if c == 'e' {
				fmt.Printf("Found one! %c\n", w[0])
			}
		}
	}
	// STOP2 OMIT
	return
}

func alternativetwo() {
	// START3 OMIT
	fmt.Println("Method three")
	wordSlice := [3]string{"Gate", "Comet", "Pizza"} // Create Array
	fmt.Println(wordSlice)

	for _, w := range wordSlice {
		if strings.ContainsAny(w, "e") {
			fmt.Printf("Found one! %c\n", w[0])
		}
	}
	// STOP3 OMIT
	return
}
