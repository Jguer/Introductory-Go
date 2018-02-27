// +build OMIT

package main

import "fmt"

func main() {
	// START1 OMIT
	s := make([]string, 3) // HL
	fmt.Println("emp:", s)

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"

	fmt.Println("set:", s)
	fmt.Println("get:", s[2])
	fmt.Println("len:", len(s))
	fmt.Println("cap1:", cap(s))

	s = append(s, "d") // HL
	fmt.Println("cap2:", cap(s))
	s = append(s, "e", "f") // HL
	fmt.Println("apd:", s)

	c := make([]string, len(s))
	copy(c, s) // HL
	fmt.Println("cpy:", c)
	// STOP1 OMIT

	fmt.Print("\n\n\n\n\n\n")
	// START2 OMIT
	fmt.Println("s:", s)
	l := s[2:5] // HL
	fmt.Println("sl1:", l)
	l = s[:5] // HL
	fmt.Println("sl2:", l)
	l = s[2:] // HL
	fmt.Println("sl3:", l)

	t := []string{"g", "h", "i"}
	fmt.Println("dcl:", t)

	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
	// STOP2 OMIT

}
