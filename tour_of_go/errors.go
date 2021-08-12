package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("Cannot Sqrt negative number:%v", float64(e))
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}
	ans, temp := float64(1), float64(0)
	for {
		ans, temp = ans-(ans*ans-x)/(2*ans), ans
		if math.Abs(temp-ans) < 1e-6 {
			break
		}
	}
	return ans, nil

}

func main() {
	fmt.Println(Sqrt(4))
	fmt.Println(Sqrt(-2))
}
