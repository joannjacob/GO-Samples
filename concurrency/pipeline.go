package main

import (
	"fmt"
	"sync"
)

// func gen(nums ...int) <-chan int {
// 	out := make(chan int)
// 	go func() {
// 		for _, n := range nums {
// 			out <- n
// 		}
// 		close(out)
// 	}()
// 	return out
// }

// BUFFER IMPLEMENTATION
func gen(nums ...int) <-chan int {
	out := make(chan int, len(nums))
	for _, n := range nums {
		out <- n
	}
	close(out)
	return out
}

func sq(done <-chan struct{}, in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for n := range in {
			select {
				case out <- n * n
			case <- done :
				return
			}
		}		
		// close(out)
	}()
	return out
}

func merge(done <-chan struct{}, cs ...<-chan int) <-chan int {
	var wg sync.WaitGroup
	out := make(chan int)
	// out := make(chan int, 1) // enough space for the unread inputs

	// Start an output goroutine for each input channel in cs.  output
	// copies values from c to out until c is closed or it receives a value
	// from done, then output calls wg.Done.
	output := func(c <-chan int) {
		defer wg.Done()
		for n := range c {
			select {
			case out <- n:
			case <-done:
				return
			}
		}
		wg.Done()
	}
	// // Start an output goroutine for each input channel in cs.  output
	// // copies values from c to out until c is closed, then calls wg.Done.
	// output := func(c <-chan int) {
	// 	for n := range c {
	// 		out <- n
	// 	}
	// 	wg.Done()
	// }
	wg.Add(len(cs))
	for _, c := range cs {
		go output(c)
	}

	// Start a goroutine to close out once all the output goroutines are
	// done.  This must start after the wg.Add call.
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

// func main() {
// 	// Set up the pipeline.
// 	c := gen(2, 3)
// 	out := sq(c)

// 	// Consume the output.
// 	fmt.Println(<-out) // 4
// 	fmt.Println(<-out) // 9
// }

//LOOPING
// func main() {
// 	// Set up the pipeline and consume the output.
// 	for n := range sq(sq(gen(2, 3))) {
// 		fmt.Println(n) // 16 then 81
// 	}
// }

// FAN IN - FAN OUT
// func main() {
// 	in := gen(2, 3)

// 	// Distribute the sq work across two goroutines that both read from in.
// 	c1 := sq(in)
// 	c2 := sq(in)

// 	// Consume the merged output from c1 and c2.
// 	for n := range merge(c1, c2) {
// 		fmt.Println(n) // 4 then 9, or 9 then 4
// 	}
// }

// EXPLICIT CANCELLATION
// func main() {
// 	in := gen(2, 3)

// 	// Distribute the sq work across two goroutines that both read from in.
// 	c1 := sq(in)
// 	c2 := sq(in)

// 	// Consume the first value from output.
// 	done := make(chan struct{}, 2)
// 	out := merge(done, c1, c2)
// 	fmt.Println(<-out) // 4 or 9

// 	// Tell the remaining senders we're leaving.
// 	done <- struct{}{}
// 	done <- struct{}{}
// }

//PROPER IMPLEMENTATION

func main() {
	// Set up a done channel that's shared by the whole pipeline,
	// and close that channel when this pipeline exits, as a signal
	// for all the goroutines we started to exit.
	done := make(chan struct{})
	defer close(done)

	in := gen(done, 2, 3)

	// Distribute the sq work across two goroutines that both read from in.
	c1 := sq(done, in)
	c2 := sq(done, in)

	// Consume the first value from output.
	out := merge(done, c1, c2)
	fmt.Println(<-out) // 4 or 9

	// done will be closed by the deferred call.
}
