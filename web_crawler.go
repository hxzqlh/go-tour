package main

import (
	"fmt"
	"sync"
)

type Fetcher interface {
	// Fetch 返回 URL 的 body 内容，并且将在这个页面上找到的 URL 放到一个 slice 中。
	Fetch(url string) (body string, urls []string, err error)
}

type SafeCenter struct {
	urlsMap map[string]string
	mux     sync.Mutex
	wg      sync.WaitGroup
}

// Crawl 使用 fetcher 从某个 URL 开始递归的爬取页面，直到达到最大深度。
func Crawl(url string, depth int, fetcher Fetcher, sf *SafeCenter) {
	// TODO: 并行的抓取 URL。
	// TODO: 不重复抓取页面。
	// 下面并没有实现上面两种情况：

	defer sf.wg.Done()
	if depth <= 0 {
		return
	}
	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		//	fmt.Println(err)
		return
	}
	(*sf).mux.Lock()
	(*sf).urlsMap[body] = url
	(*sf).mux.Unlock()
	//fmt.Printf("found: %s %q\n", url, body)
	for _, u := range urls {
		sf.wg.Add(1)
		go Crawl(u, depth-1, fetcher, sf)
	}
	return
}

func main() {
	var sf SafeCenter
	sf.urlsMap = make(map[string]string)
	sf.wg.Add(1)
	Crawl("https://golang.org/", 4, fetcher, &sf)
	sf.wg.Wait()
	for k, v := range sf.urlsMap {
		fmt.Printf("found: %s %q\n", v, k)
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
