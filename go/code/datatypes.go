// +build OMIT

// Single line comment
/* Multi-
line comment */

package main

import (
	"fmt"
	"math/cmplx"
)

var (
	maxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

func main() {
	toBe := true
	fmt.Printf("%T(%v)\n%T(%v)\n%T(%v)\n", toBe, toBe, maxInt, maxInt, z, z)
}
