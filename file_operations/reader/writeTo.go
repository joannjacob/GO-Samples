package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
)

type R struct {
	n int
}

func (r *R) Read(p []byte) (n int, err error) {
	fmt.Printf("Read #%d\n", r.n)
	if r.n >= 10 {
		return 0, io.EOF
	}
	copy(p, "abcd")
	r.n += 1
	return 4, nil
}
func main() {
	r := bufio.NewReaderSize(new(R), 16)
	n, err := r.WriteTo(ioutil.Discard)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Written bytes: %d\n", n)
}
