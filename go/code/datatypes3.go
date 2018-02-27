// +build OMIT

package main

import (
	"fmt"
)

func main() {
	newString := "It's me HES"
	fmt.Print(newString, "\n")

	fmt.Println("go" + "lang")

	gOne := "go"
	gTwo := "lang"
	fmt.Println(gOne + gTwo)
}
