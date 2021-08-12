package main

import (
	"fmt"
	"math"
)

func Sqrt(n float64) float64 {
	ans, temp := float64(1), float64(0)
	for {
		ans, temp = ans-(ans*ans-n)/(2*ans), ans
		if math.Abs(temp-ans) < 1e-6 {
			break
		}
	}
	return ans
}
func main() {
	guess := Sqrt(2)
	actual := math.Sqrt(2)
	fmt.Printf("Guess: %v\n", guess)
	fmt.Printf("Expected: %v\n", actual)
	fmt.Printf("Error: %v\n", math.Abs(guess-actual))
}
