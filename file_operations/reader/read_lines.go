package main

import (
	"bufio"
	"fmt"
	"strings"
)

func main() {
	s := strings.NewReader(strings.Repeat("a", 20) + "\n" + "b")
	r := bufio.NewReaderSize(s, 16)
	token, isPrefix, err := r.ReadLine()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Token: %q, prefix: %t\n", token, isPrefix)
	token, isPrefix, err = r.ReadLine()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Token: %q, prefix: %t\n", token, isPrefix)
	// token, isPrefix, err = r.ReadLine()
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Printf("Token: %q, prefix: %t\n", token, isPrefix)
	// token, isPrefix, err = r.ReadLine()
	// if err != nil {
	// 	panic(err)
	// }

	s = strings.NewReader("abc")
	token, isPrefix, err = r.ReadLine()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Token: %q, prefix: %t\n", token, isPrefix)
	s = strings.NewReader("abc\n")
	r.Reset(s)
	token, isPrefix, err = r.ReadLine()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Token: %q, prefix: %t\n", token, isPrefix)
}
