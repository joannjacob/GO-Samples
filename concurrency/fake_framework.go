package main

import (
	"fmt"
	"math/rand"
	"time"
)

// var (
// 	web   = fakeSearch("web")
// 	image = fakeSearch("image")
// 	video = fakeSearch("video")
// )

var (
	web1   = fakeSearch("web")
	web2   = fakeSearch("web")
	image1 = fakeSearch("image")
	image2 = fakeSearch("image")
	video1 = fakeSearch("video")
	video2 = fakeSearch("video")
)

type (
	result string
	search func(query string) result
)

func first(query string, replicas ...search) result {
	c := make(chan result)

	// Define a function that takes the index to the replica function to use.
	// Then it executes that function writing the results to the channel.
	searchReplica := func(i int) {
		c <- replicas[i](query)
	}

	// Run each replica function in its own Goroutine.
	for i := range replicas {
		go searchReplica(i)
	}

	// As soon as one of the replica functions write a result, return.
	return <-c
}

func main() {
	rand.Seed(time.Now().UnixNano())

	start := time.Now()
	// results := first("golang",
	// 	fakeSearch("replica 1"),
	// 	fakeSearch("replica 2"))
	results := google("golang")
	elapsed := time.Since(start)

	fmt.Println(results)
	fmt.Println(elapsed)
}

func fakeSearch(kind string) search {
	return func(query string) result {
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
		return result(fmt.Sprintf("%s result for %q\n", kind, query))
	}
}
func google(query string) (results []result) {
	c := make(chan result)

	go func() {
		c <- first(query, web1, web2)
	}()

	go func() {
		c <- first(query, image1, image2)
	}()

	go func() {
		c <- first(query, video1, video2)
	}()

	// MODIFY WITH ALL IMPROVEMENTS
	// go func() {
	// 	c <- web(query)
	// }()

	// go func() {
	// 	c <- image(query)
	// }()

	// go func() {
	// 	c <- video(query)
	// }()

	// SECOND METHOD
	// for i := 0; i < 3; i++ {
	// 	r := <-c
	// 	results = append(results, r)
	// }

	// TIMEOUT
	timeout := time.After(80 * time.Millisecond)

	for i := 0; i < 3; i++ {
		select {
		case r := <-c:
			results = append(results, r)
		case <-timeout:
			fmt.Println("timed out")
			return results
		}
	}

	return results
}
