package main

import (
	"fmt"
	"time"
)

// Fetcher is an interface.
type Fetcher interface {
	//Fetch 返回 URL 的 body 内容，并且将在这个页面上找到的 URL 放到一个 slice 中.
	Fetch(url string) (body string, urls []string, err error)
}

//ParaCrawl ...
func ParaCrawl(url string, fetcher Fetcher) []string {
	fmt.Println("Crawl link:", url)

	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		//fmt.Println(err)
		return nil
	}

	fmt.Println("	return body:", body)
	fmt.Println("	return urls:", urls)

	return urls
}

//开启多个协程并行处理，无深度限制; 指定超时时间到时，自动退出;
func main() {
	worklist := make(chan []string)  // lists of URLs, may have duplicates
	unseenLinks := make(chan string) // de-duplicated URLs
	done := make(chan struct{})

	var searchLink []string
	searchLink = append(searchLink, "https://golang.org/")
	// Add command-line arguments to worklist.
	go func() { worklist <- searchLink }()

	// Create 20 crawler goroutines to fetch each unseen link.
	for i := 0; i < 20; i++ {
		go func() {
			for {
				select {
				case link := <-unseenLinks:
					foundLinks := ParaCrawl(link, fetcher)
					if foundLinks != nil {
						go func() { worklist <- foundLinks }()
					}
				case <-done:
					return
				}
			}
		}()
	}

	// The main goroutine de-duplicates worklist items
	// and sends the unseen ones to the crawlers.
	seen := make(map[string]bool)

	out := false
	for !out {
		select {
		case lists := <-worklist:
			for _, link := range lists {
				if !seen[link] {
					seen[link] = true
					unseenLinks <- link
				}
			}

		case <-time.After(time.Second * 5):
			close(done)
			out = true
		}
	}

	fmt.Println("Timeout, go out!")
	return
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

