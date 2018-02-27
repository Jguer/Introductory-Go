// +build OMIT

package main

import (
	"fmt"
	"math"
)

// START1 OMIT

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex{3, 4}
	v.Scale(5)
	fmt.Println(v.Abs())
}

// STOP1 OMIT
