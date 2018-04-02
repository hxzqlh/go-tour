package main

import (
	"fmt"
	"sync"
)

type Safe struct {
	aurls []string
	mux   sync.Mutex
}

var num int = 0

type Fetcher interface {
	// Fetch 返回 URL 的 body 内容，并且将在这个页面上找到的 URL 放到一个 slice 中。
	Fetch(url string) (body string, urls []string, err error)
}

func gothrough(url string, aurls []string) bool {
	for _, v := range aurls {
		//fmt.Println(v + "hshsaihihs")
		if url == v {
			return false
		}
	}
	return true
}

var crawl = Safe{make([]string, 10), sync.Mutex{}}

//全局变量
var c1 chan string

var quit chan bool

// Crawl 使用 fetcher 从某个 URL 开始递归的爬取页面，直到达到最大深度。
func Crawl(url string, depth int, fetcher Fetcher) {
	//fmt.Println(url)
	// TODO: 并行的抓取 URL。
	// TODO: 不重复抓取页面。
	// 下面并没有实现上面两种情况：
	if depth <= 0 {
		quit <- true
		return
	}
	crawl.mux.Lock()
	if gothrough(url, crawl.aurls) == false {
		crawl.mux.Unlock()
		quit <- true
		return
	}
	crawl.aurls = append(crawl.aurls, url)
	crawl.mux.Unlock()

	_, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
		quit <- true
		return
	}

	c1 <- url

	for _, u := range urls {
		go Crawl(u, depth-1, fetcher)
	}
	for i := 0; i < len(urls); i++ {
		num++
		<-quit
	}
	quit <- true
	return

}

func main() {
	c1 = make(chan string)
	quit = make(chan bool)

	go Crawl("https://golang.org/", 4, fetcher)
	for {
		select {
		case url := <-c1:
			fmt.Println("found: ", url)
		case <-quit:
			fmt.Println(num)
			return
		}
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
