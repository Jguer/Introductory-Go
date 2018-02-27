// +build OMIT

package main

import (
	"fmt"
)

func main() {
	// START1 OMIT
    package main

    import "fmt"

    func main() {
        fmt.Println("Welcome to Go 101 by HS")
    }
	// STOP1 OMIT

	// START2 OMIT
    bool
    string
    Numeric types:
    uint        either 32 or 64 bits
    int         same size as uint
    uintptr     an unsigned integer large enough to store the uninterpreted bits of
    a pointer value
    uint8       the set of all unsigned  8-bit integers (0 to 255)
    uint16      the set of all unsigned 16-bit integers (0 to 65535)
    uint32      the set of all unsigned 32-bit integers (0 to 4294967295)
    uint64      the set of all unsigned 64-bit integers (0 to 18446744073709551615)
    int8        the set of all signed  8-bit integers (-128 to 127)
    int16       the set of all signed 16-bit integers (-32768 to 32767)
    int32       the set of all signed 32-bit integers (-2147483648 to 2147483647)
    int64       the set of all signed 64-bit integers
    (-9223372036854775808 to 9223372036854775807)
    float32     the set of all IEEE-754 32-bit floating-point numbers
    float64     the set of all IEEE-754 64-bit floating-point numbers
    complex64   the set of all complex numbers with float32 real and imaginary parts
    complex128  the set of all complex numbers with float64 real and imaginary parts
    byte        alias for uint8
    rune        alias for int32 (represents a Unicode code point)
	// STOP2 OMIT

    #arrays
	// START3 OMIT
	var a [5]int
    b := [...]int{5, 10, 15}
    c := [5]int{5, 10, 15}

	// STOP3 OMIT

    #slices
	// START4 OMIT
    var a []int
    b := make([]int, 0)
    c := []string{"a", "b", "c", "d"}

	a = append(a, 3)
    len(a)
    cap(a)
	// STOP4 OMIT

	// START5 OMIT
    if
    for
    while
    switch
	// STOP5 OMIT
}
