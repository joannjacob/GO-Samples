package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

func main() {

	s := strings.NewReader(strings.Repeat("a", 64*1024) + "\n")
	r := bufio.NewReader(s)
	token, _, err := r.ReadLine()
	fmt.Printf("Token (ReadLine): %d\n", len(token))
	fmt.Printf("Error (ReadLine): %v\n", err)
	s.Seek(0, io.SeekStart)
	r.Reset(s)
	token, err = r.ReadBytes('\n')
	fmt.Printf("Token (ReadBytes): %d\n", len(token))
	fmt.Printf("Error (ReadBytes): %v\n", err)
	s.Seek(0, io.SeekStart)
	scanner := bufio.NewScanner(s)
	scanner.Scan()
	fmt.Printf("Token (Scanner): %d\n", len(scanner.Text()))
	fmt.Printf("Error (Scanner): %v\n", scanner.Err())
}
