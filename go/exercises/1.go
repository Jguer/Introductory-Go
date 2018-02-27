// +build OMIT

package main

import "fmt"

func main() {
	fmt.Println("Method one")
	// START1 OMIT
	for i := 12; i >= 0; i-- {
		if i == 10 || i == 4 || i == 2 {
			continue
		}
		fmt.Println("i = ", i)
	}
	// STOP1 OMIT

	alternativeone()
	alternativetwo()
}

func alternativeone() {
	fmt.Println("Method two")
	// START2 OMIT
	for i := 12; i >= 0; i-- {
		if i != 10 && i != 4 && i != 2 {
			fmt.Printf("i = %d\n", i)
			continue
		}
	}
	// STOP2 OMIT
	return
}

func alternativetwo() {
	fmt.Println("Method three")
	// START3 OMIT
	for i := 12; i >= 0; i-- {
		switch i {
		case 10, 4, 2:
			continue
		default:
			fmt.Println("i = ", i)
		}
	}
	// STOP3 OMIT
	return
}
