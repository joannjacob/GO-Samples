package main

import (
	"bufio"
	"fmt"
	"strings"
)

func main() {
	s := strings.NewReader(strings.Repeat("a", 40) + "|")
	r := bufio.NewReaderSize(s, 16)
	token, err := r.ReadBytes('|')
	if err != nil {
		panic(err)
	}
	fmt.Printf("Token: %q\n", token)
}
