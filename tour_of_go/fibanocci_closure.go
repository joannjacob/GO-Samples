package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	f1, f2 := 0, 1
	return func() int {
		f3 := f1
		f1, f2 = f2, f3+f2

		return f3
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Print(f(), " ")
	}
}