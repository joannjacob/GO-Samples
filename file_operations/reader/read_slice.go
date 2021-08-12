package main

import (
	"bufio"
	"fmt"
	"strings"
)

func main() {
	s := strings.NewReader("abcdefghij")
	// s := strings.NewReader("abcdef|ghij")
	r := bufio.NewReader(s)
	token, err := r.ReadSlice('|')
	if err != nil {
		panic(err)
	}
	fmt.Printf("Token: %q\n", token)

}
