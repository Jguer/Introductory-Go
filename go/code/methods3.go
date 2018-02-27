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

func AbsFunc(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())
	fmt.Println(AbsFunc(v))

	vp := &Vertex{4, 3}
	fmt.Println(vp.Abs())
	fmt.Println(AbsFunc(*vp))
}

// STOP1 OMIT
