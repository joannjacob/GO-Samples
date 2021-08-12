package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	c := fanIn(boring("Joe"), boring("Ann"))

	for i := 0; i < 10; i++ {
		fmt.Println(<-c) // Display any message received on the FanIn channel.
	}

	fmt.Println("You're boring: I'm leaving.")
}

func fanIn(input1, input2 <-chan string) <-chan string {
	c := make(chan string) // The FanIn channel

	go func() { // This Goroutine will receive messages from Joe.
		for {
			c <- <-input1 // Write the message to the FanIn channel, Blocking Call.
		}
	}()

	go func() { // This Goroutine will receive messages from Ann
		for {
			c <- <-input2 // Write the message to the FanIn channel, Blocking Call.
		}
	}()

	return c
}

func boring(msg string) <-chan string { // Returns receive-only (<-) channel of strings.
	c := make(chan string)

	go func() { // Launch the goroutine from inside the function. Function Literal.
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}()

	return c // Return the channel to the caller.
}
