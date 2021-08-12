package main

import (
	"fmt"
	"time"
)

// func Query(conns []Conn, query string) Result {
// 	ch := make(chan Result)
// 	for _, conn := range conns {
// 		go func(c Conn) {
// 			select {
// 			case ch <- c.DoQuery(query):
// 			default:
// 			}
// 		}(conn)
// 	}
// 	return <-ch
// }

func main() {
	timeout := make(chan bool, 1)
	go func() {
		time.Sleep(1 * time.Second)
		timeout <- true
	}()

	ch1 := make(chan string)

	select {
	case <-ch1:
		// a read from ch has occurred
		fmt.Println("channel")
	case <-timeout:
		// the read from ch has timed out
		fmt.Println("Timeout")
	}
}
