package main

import (
	"fmt"
	"sync"
)

type Fetcher interface {
	// Fetch 返回 URL 的 body 内容，并且将在这个页面上找到的 URL 放到一个 slice 中。
	Fetch(url string) (body string, urls []string, err error)
}
type UrlCounter struct {
	v   map[string]bool
	mux sync.Mutex
}

func checkUrl(url string) bool {
	urlCounter.mux.Lock()
	defer urlCounter.mux.Unlock()
	if urlCounter.v[url] == true {
		return true
	}
	return false
}

func markUrlCount(url string) {
	urlCounter.mux.Lock()
	defer urlCounter.mux.Unlock()
	urlCounter.v[url] = true
}

// Crawl 使用 fetcher 从某个 URL 开始递归的爬取页面，直到达到最大深度。
func Crawl(url string, depth int, fetcher Fetcher) {
	// TODO: 并行的抓取 URL。
	// TODO: 不重复抓取页面。
	// 下面并没有实现上面两种情况：
	if depth <= 0 {
		return
	}
	if checkUrl(url) {
		return
	}
	body, urls, err := fetcher.Fetch(url)
	markUrlCount(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("found: %s %q\n", url, body)
	ch := make(chan int)
	for _, u := range urls {
		go func(url string) {
			Crawl(url, depth-1, fetcher)
			ch <- 1
		}(u)
	}
	for range urls {
		<-ch
	}
	return
}

var urlCounter UrlCounter

func main() {
	urlCounter = UrlCounter{v: make(map[string]bool)}
	Crawl("https://golang.org/", 4, fetcher)
}

// fakeFetcher 是返回若干结果的 Fetcher。
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

// fetcher 是填充后的 fakeFetcher。
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
