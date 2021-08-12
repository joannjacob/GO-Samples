package main

import "fmt"

func Sum(x int, y int) int {
	return x + y
}

func IntMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	fmt.Println(Sum(5, 5))
	fmt.Println(IntMin(2, 3))
}
