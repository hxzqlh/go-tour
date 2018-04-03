package main

import (
	"fmt"
	"sync"
)

type Fetcher interface {
	// Fetch 返回 URL 的 body 内容，并且将在这个页面上找到的 URL 放到一个 slice 中。
	Fetch(url string) (body string, urls []string, err error)
}

var mux = sync.Mutex{}
var urlmap = map[string]bool{}

// Crawl 使用 fetcher 从某个 URL 开始递归的爬取页面，直到达到最大深度。
func Crawl(url string, depth int, fetcher Fetcher, c chan bool) {
	// TODO: 并行的抓取 URL。
	// TODO: 不重复抓取页面。
	// 下面并没有实现上面两种情况：
	if depth <= 0 {
		c <- true
		return
	}

	mux.Lock()
	_, ok := urlmap[url]
	mux.Unlock()

	if ok {
		c <- true
		return
	}

	body, urls, err := fetcher.Fetch(url)

	mux.Lock()
	urlmap[url] = true
	mux.Unlock()

	if err != nil {
		fmt.Println(err)
		c <- true
		return
	}
	fmt.Printf("found: %s %q\n", url, body)

	subc := make(chan bool)
	for _, u := range urls {
		go Crawl(u, depth-1, fetcher, subc)
	}

	for i := 0; i < len(urls); i++ {
		<-subc
	}
	c <- true
	return
}

func main() {
	c := make(chan bool, 10)
	Crawl("https://golang.org/", 4, fetcher, c)
	<-c
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
