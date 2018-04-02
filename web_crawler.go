package main

import (
	"fmt"
	"sync"
	//"time"
)

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

type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

type Crawled struct {
	crawleds map[string]int
	mux      sync.Mutex
}

type Fetcher interface {
	// Fetch 返回 URL 的 body 内容，并且将在这个页面上找到的 URL 放到一个 slice 中。
	Fetch(url string) (body string, urls []string, err error)
}

func main() {
	crawled := Crawled{make(map[string]int), sync.Mutex{}}
	out := make(chan string)
	end := make(chan bool)
	go Crawl("https://golang.org/", 4, fetcher, crawled, out, end)
	//time.Sleep(5 * time.Second)

	for {
		select {
		case url := <-out:
			fmt.Println(url)
		case <-end:
			return
		}
	}
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

// Crawl 使用 fetcher 从某个 URL 开始递归的爬取页面，直到达到最大深度。
func Crawl(url string, depth int, fetcher Fetcher, crawled Crawled, out chan string, end chan bool) {
	if depth <= 0 {
		end <- true
		return
	}

	crawled.mux.Lock()

	if _, ok := crawled.crawleds[url]; ok {
		crawled.mux.Unlock()
		end <- true
		return
	}

	crawled.crawleds[url] = 1
	crawled.mux.Unlock()

	_, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
		end <- true
		return
	}

	//fmt.Printf("found: %s %q\n", url, body)

	out <- url

	for _, u := range urls {
		go Crawl(u, depth-1, fetcher, crawled, out, end)
	}

	for i := 0; i < len(urls); i++ {
		<-end
	}

	end <- true
	return
}
