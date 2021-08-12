package main

import (
	"bufio"
	"fmt"
)

type Writer int

func (*Writer) Write(p []byte) (int, error) {

	return len(p), nil
}

func main() {

	w := new(Writer)
	bw := bufio.NewWriterSize(w, 10)
	fmt.Println(bw.Buffered())
	bw.WriteByte('a')
	fmt.Println(bw.Buffered())
	bw.WriteRune('ł') // 'ł' occupies 2 bytes
	fmt.Println(bw.Buffered())
	bw.WriteString("aa")
	fmt.Println(bw.Buffered())
}
