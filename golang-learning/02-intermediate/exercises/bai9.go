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
	time.Sleep(2 * time.Second) // Simulate network delay
	// For simplicity, return fake content
	return fmt.Sprintf("Content of %s", url), nil
}

// TODO: Worker function
func scrapeWorker(id int, urls <-chan string, results chan<- ScrapeResult) {
	// Implement here
	// Nhận URL từ channel
	// Fetch với timeout
	// Gửi result vào results channel
	for url := range urls {
		content, err := fetchURL(url)
		result := ScrapeResult{
			URL:     url,
			Content: content,
			Error:   err,
		}
		results <- result
	}
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
	urlChan := make(chan string, len(urls))
	results := make(chan ScrapeResult, len(urls))

	// TODO: Start worker pool (5 workers)
	for i := 0; i < 5; i++ {
		go scrapeWorker(i, urlChan, results)
	}

	// TODO: Send URLs
	for _, url := range urls {
		urlChan <- url
	}
	close(urlChan)

	// TODO: Collect results
	for i := 0; i < len(urls); i++ {
		result := <-results
		if result.Error != nil {
			fmt.Printf("Error scraping %s: %v\n", result.URL, result.Error)
		} else {
			fmt.Printf("Scraped %s: %s\n", result.URL, result.Content)
		}
	}
	close(results)

	// TODO: Print summary
	fmt.Println("Scraping completed.")
}

func main() {
	challenge1()
}
