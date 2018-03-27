package main

import (
	"fmt"
	"sync"
)

type Fetcher interface {
	Fetch(url string) (body string, urls []string, err error)
}

type SafeMap struct {
	h   map[string]bool
	mux sync.Mutex
}

func (mp *SafeMap) Contains(key string) bool {
	mp.mux.Lock()
	_, in := mp.h[key]
	defer mp.mux.Unlock()
	return in
}

func (mp *SafeMap) Insert(key string) {
	mp.mux.Lock()
	mp.h[key] = true
	mp.mux.Unlock()
}

var mp = SafeMap{h: make(map[string]bool)}

func Crawl(url string, depth int, fetcher Fetcher, ori *sync.WaitGroup) {
	if depth <= 0 {
		ori.Done()
		return
	}
	if mp.Contains(url) {
		ori.Done()
		return
	}
	mp.Insert(url)
	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
		ori.Done()
		return
	}
	fmt.Printf("found: %s %q\n", url, body)
	var queue sync.WaitGroup
	for _, u := range urls {
		queue.Add(1)
		go Crawl(u, depth-1, fetcher, &queue)
	}
	queue.Wait()
	ori.Done()
}

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go Crawl("https://golang.org/", 4, fetcher, &wg)
	wg.Wait()
}

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
