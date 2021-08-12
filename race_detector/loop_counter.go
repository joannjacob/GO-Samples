package main

import (
	"fmt"
	"sync"
)

// func main() {
// 	var wg sync.WaitGroup
// 	wg.Add(5)
// 	for i := 0; i < 5; i++ {
// 		go func() {
// 			fmt.Println(i) // Not the 'i' you are looking for.
// 			wg.Done()
// 		}()
// 	}
// 	wg.Wait()
// }

//FIX

func main() {
	var wg sync.WaitGroup
	wg.Add(5)
	for i := 0; i < 5; i++ {
		go func(j int) {
			fmt.Println(j) // Good. Read local copy of the loop counter.
			wg.Done()
		}(i)
	}
	wg.Wait()
}
