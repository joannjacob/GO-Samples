package main

import (
	"bufio"
	"fmt"
)

type R int

func (r *R) Read(p []byte) (n int, err error) {
	fmt.Println("Read")
	copy(p, "abcd")
	return 4, nil
}
func main() {
	r := new(R)
	br := bufio.NewReader(r)
	buf := make([]byte, 2)
	n, err := br.Read(buf)
	if err != nil {
		panic(err)
	}
	buf = make([]byte, 4)
	n, err = br.Read(buf)
	if err != nil {
		panic(err)
	}
	fmt.Printf("read = %q, n = %d\n", buf[:n], n)
}
