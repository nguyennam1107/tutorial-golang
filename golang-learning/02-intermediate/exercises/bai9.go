package main

import (
	"fmt"
	"time"
)

// ============= CHALLENGE 1: CONCURRENT WEB SCRAPER =============

/*
YÊU CẦU:
1. Scrape 10 URLs concurrently
2. Sử dụng worker pool pattern (5 workers)
3. Handle timeout (3 seconds per URL)
4. Collect và in results
5. Handle errors gracefully
*/

type ScrapeResult struct {
	URL     string
	Title   string
	Content string
	Error   error
}

// TODO: Simulate fetch URL
func fetchURL(url string) (string, error) {
	// Implement here
	// Simulate HTTP request với time.Sleep
	// Return fake content hoặc error
	return "", nil
}

// TODO: Worker function
func scrapeWorker(id int, urls <-chan string, results chan<- ScrapeResult) {
	// Implement here
	// Nhận URL từ channel
	// Fetch với timeout
	// Gửi result vào results channel
}

func challenge1() {
	fmt.Println("\n=== CHALLENGE 1: WEB SCRAPER ===")

	urls := []string{
		"https://golang.org",
		"https://github.com",
		"https://stackoverflow.com",
		"https://reddit.com",
		"https://medium.com",
		"https://dev.to",
		"https://news.ycombinator.com",
		"https://lobste.rs",
		"https://twitter.com",
		"https://facebook.com",
	}

	// TODO: Setup channels
	// urlChan := make(chan string, len(urls))
	// results := make(chan ScrapeResult, len(urls))

	// TODO: Start worker pool (5 workers)

	// TODO: Send URLs

	// TODO: Collect results

	// TODO: Print summary
}

func main() {
	challenge1()
}
