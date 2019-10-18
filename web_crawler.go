package main

import (
	"fmt"
	"runtime"
	"sync"
)

type Urls struct {
	done map[string]bool
	mux  sync.Mutex
	// rutCnt int
	// cmux   sync.Mutex
}

type Fetcher interface {
	// Fetch 返回 URL 的 body 内容，并且将在这个页面上找到的 URL 放到一个 slice 中。
	Fetch(url string) (body string, urls []string, err error)
}

// Crawl 使用 fetcher 从某个 URL 开始递归的爬取页面，直到达到最大深度。
func Crawl(url string, U *Urls, depth int, fetcher Fetcher, ch chan string) {
	// TODO: 并行的抓取 URL。
	// TODO: 不重复抓取页面。
	// 下面并没有实现上面两种情况：
	if depth <= 0 {
		return
	}
	// U.cmux.Lock()
	// U.rutCnt++
	// U.cmux.Unlock()
	body, urls, err := fetcher.Fetch(url)
	U.mux.Lock()
	U.done[url] = true
	U.mux.Unlock()
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}
	ch <- fmt.Sprintf("found: %s %q\n", url, body)
	for _, u := range urls {
		count := len(urls)
		U.mux.Lock()
		if !U.done[u] {
			count++
			go Crawl(u, U, depth-1, fetcher, ch)
		}
		U.mux.Unlock()
		if count == 0 {
			close(ch)
		}
	}
	return
}

func main() {
	ch := make(chan string)
	U := Urls{done: make(map[string]bool)}
	go Crawl("https://golang.org/", &U, 4, fetcher, ch)
	// var RC int
	for v := range ch {
		fmt.Println(v)
		// U.cmux.Lock()
		// rc := U.rutCnt
		// U.cmux.Unlock()
		if runtime.NumGoroutine() == 1 {
			return
		}
		// RC = rc
	}
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
