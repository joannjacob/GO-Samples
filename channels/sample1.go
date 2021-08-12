package main

import (
	"fmt"
	"time"
)

func makeMessage(s string, channel chan string) {
	for i := 0; i < 2; i++ {
		channel <- fmt.Sprintf("I wanted to say '%s' for the %dth time", s, i) // write the message to the channel
		time.Sleep(500 * time.Millisecond)                                     // to show that the main function will wait for the channel to be filled again
	}
}
func main() {
	channel := make(chan string)      // create channel
	go makeMessage("Hello!", channel) // run this in a separate goroutine
	for i := 0; i < 2; i++ {
		fmt.Printf("Hey there, %s\n", <-channel) // read from channel and print out message
	}
	fmt.Println("Cool, that's all I wanted to say")
}
