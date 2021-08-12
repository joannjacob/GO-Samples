package main

import (
	"bufio"
	"fmt"
	"strings"
)

func main() {
	s1 := strings.NewReader(strings.Repeat("a", 20))
	r := bufio.NewReaderSize(s1, 16)
	b, err := r.Peek(3)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%q\n", b)
	b, err = r.Peek(17)
	if err != nil {
		fmt.Println(err)
	}
	s2 := strings.NewReader("aaa")
	r.Reset(s2)
	b, err = r.Peek(10)
	if err != nil {
		fmt.Println(err)
	}

	// s1 := strings.NewReader(strings.Repeat("a", 16) + strings.Repeat("b", 16))
	// r := bufio.NewReaderSize(s1, 16)
	// b, _ := r.Peek(3)
	// fmt.Printf("%q\n", b)
	// r.Read(make([]byte, 16))
	// r.Read(make([]byte, 15))
	// fmt.Printf("%q\n", b)
}
