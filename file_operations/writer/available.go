package main

import (
	"bufio"
	"fmt"
)

type Writer int

func (*Writer) Write(p []byte) (n int, err error) {
	// fmt.Println(len(p))
	return len(p), nil
}

func main() {
	w := new(Writer)
	bw := bufio.NewWriterSize(w, 2)
	fmt.Println(bw.Available())
	bw.Write([]byte{'a'})
	fmt.Println(bw.Available())
	bw.Write([]byte{'b'})
	fmt.Println(bw.Available())
	bw.Write([]byte{'c'})
	fmt.Println(bw.Available())
}
