// +build OMIT

package main

import (
	"fmt"
)

// START1 OMIT
func valchanger(in string) (out string) {
	in = in[0:46] + in[50:55]
	fmt.Println("Inside func:", in)
	out = in
	return
}

func main() {
	in := "What you're referring to as Linux,is in fact, GNU/Linux"
	valchanger(in)
	fmt.Println("Outside func:", in)
	in = valchanger(in)
	fmt.Println("Outside func:", in)

}

// STOP1 OMIT
