// +build OMIT

package main

// START0 OMIT
import (
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
	fmt.Println("Test", rand.Intn(999)) // 0 < i <= 999
}

// STOP0 OMIT

func main() {
	// START1 OMIT
	numap := map[int]string{
		0: "zero", 1: "one", 2: "two", 3: "three", 4: "four",
		5: "five", 6: "six", 7: "seven", 8: "eight", 9: "nine",
	}
	randomNum := rand.Intn(999)
	fmt.Println("Random Number", randomNum)
	fmt.Print("Converted: ")

	if randomNum/100 != 0 {
		fmt.Print(numap[randomNum/100], " ")
	}
	if n := randomNum % 100; n%10 != 0 {
		fmt.Print(numap[n/10], " ")
		fmt.Print(numap[n%10])

	}
	fmt.Print("\n")
	// STOP1 OMIT
	alternativeone(numap)

}

func alternativeone(numap map[int]string) {
	// START2 OMIT
	randomNum := rand.Intn(999)
	fmt.Println("Random Number", randomNum)

	temp := randomNum
	factor := 1
	for temp > 0 {
		temp = temp / 10
		factor = factor * 10
	}

	fmt.Print("Converted: ")
	for factor > 1 {
		factor = factor / 10
		fmt.Printf("%s ", numap[randomNum/factor])
		randomNum = randomNum % factor
	}
	fmt.Print("\n")

	// STOP2 OMIT
	return
}
