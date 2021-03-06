package main

import (
	"bufio"
	"fmt"
)

type Writer int

func (*Writer) Write(p []byte) (n int, err error) {
	return len(p), nil
}
func main() {
	w := new(Writer)
	bw := bufio.NewWriterSize(w, 3)
	fmt.Println(bw.Buffered())
	bw.Write([]byte{'a'})
	fmt.Println(bw.Buffered())
	bw.Write([]byte{'b'})
	fmt.Println(bw.Buffered())
	bw.Write([]byte{'c'})
	fmt.Println(bw.Buffered())
	bw.Write([]byte{'d'})
	fmt.Println(bw.Buffered())
	bw.Flush()
	fmt.Println(bw.Buffered())
}
