package main

import (
	"fmt"
)

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

// type SafeCounter struct {
// 	v   map[string]bool
// 	mux sync.Mutex
// 	wg  sync.WaitGroup
// }

// var c SafeCounter = SafeCounter{v: make(map[string]bool)}

// func (s SafeCounter) checkvisited(url string) bool {
// 	s.mux.Lock()
// 	defer s.mux.Unlock()
// 	_, ok := s.v[url]
// 	if ok == false {
// 		s.v[url] = true
// 		return false
// 	}
// 	return true

// }

// // Crawl uses fetcher to recursively crawl
// // pages starting with url, to a maximum of depth.
// func Crawl(url string, depth int, fetcher Fetcher) {
// 	// TODO: Fetch URLs in parallel.
// 	// TODO: Don't fetch the same URL twice.
// 	// This implementation doesn't do either:
// 	defer c.wg.Done()
// 	if depth <= 0 {
// 		return
// 	}
// 	if c.checkvisited(url) {
// 		return
// 	}
// 	body, urls, err := fetcher.Fetch(url)
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	fmt.Printf("Found urls: %s %q\n", url, body)
// 	for _, u := range urls {
// 		c.wg.Add(1)
// 		go Crawl(u, depth-1, fetcher)
// 	}
// 	return
// }
func CrawlRecursive(url string, depth int, fetcher Fetcher, quit chan bool, visitedUrls map[string]bool) {
	if depth <= 0 {
		quit <- true
		return
	}

	didIt, hasIt := visitedUrls[url]

	// If we have already visited this link,stop here
	if didIt && hasIt {
		quit <- true
		return

	} else {
		// Mark it has visited
		visitedUrls[url] = true
	}

	// Fetch children URLs
	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
		quit <- true
		return
	}
	fmt.Printf("Found URL: %s  Body: %q\n", url, body)

	// Crawl children URLs
	childrenQuit := make(chan bool)
	for _, childrenUrl := range urls {
		go CrawlRecursive(childrenUrl, depth-1, fetcher, childrenQuit, visitedUrls)
		// To exit goroutines. This channel will always be filled
		<-childrenQuit
	}

	quit <- true
}

func Crawl(url string, depth int, fetcher Fetcher) {
	quit := make(chan bool)

	visitedUrls := map[string]bool{url: false}

	go CrawlRecursive(url, depth, fetcher, quit, visitedUrls)

	// We will not quit until we have something
	// in the "quit" channel
	<-quit
}

func main() {
	// c.wg.Add(1)
	Crawl("https://golang.org/", 4, fetcher)
	// c.wg.Wait()
}

// fakeFetcher is Fetcher that returns canned results.
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

// fetcher is a populated fakeFetcher.
var fetcher = fakeFetcher{
	"https://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"https://golang.org/pkg/",
			"https://golang.org/cmd/",
		},
	},
	"https://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"https://golang.org/",
			"https://golang.org/cmd/",
			"https://golang.org/pkg/fmt/",
			"https://golang.org/pkg/os/",
		},
	},
	"https://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
	"https://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
}
