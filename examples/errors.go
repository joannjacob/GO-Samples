package main

import (
	"errors"
	"fmt"
	"math"
)

// errorString is a trivial implementation of error.
type errorString struct {
	s string
}

func (e *errorString) Error() string {
	return e.s
}

// New returns an error that formats as the given text.
func New(text string) error {
	return &errorString{text}
}

func Sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, errors.New("math: square root of negative number")
	}
	// implementation
	return math.Sqrt(f), nil
}

func main() {
	// f, err := os.Open("file.txt")
	// if err != nil {
	// 	log.Fatal(err)
	// 	fmt.Printf("%T", err)
	// }
	// fmt.Println(f)

	val, err1 := Sqrt(-4)
	if err1 != nil {
		fmt.Println(err1)
	}
	print(val)
}
