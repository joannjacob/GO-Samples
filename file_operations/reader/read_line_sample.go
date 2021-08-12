package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

func main() {

	s := strings.NewReader(strings.Repeat("a", 20) + "\n")
	r := bufio.NewReaderSize(s, 16)
	token, _, _ := r.ReadLine()
	fmt.Printf("Token (ReadLine): \t%q\n", token)
	s.Seek(0, io.SeekStart)
	r.Reset(s)
	token, _ = r.ReadBytes('\n')
	fmt.Printf("Token (ReadBytes): \t%q\n", token)
	s.Seek(0, io.SeekStart)
	scanner := bufio.NewScanner(s)
	scanner.Scan()
	fmt.Printf("Token (Scanner): \t%q\n", scanner.Text())

}
