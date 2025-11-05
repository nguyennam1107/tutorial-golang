package main

import (
	"fmt"
	"time"
)

// ============= CHALLENGE 2: RATE LIMITER =============

/*
YÊU CẦU:
1. Implement rate limiter: 5 requests per second
2. Queue requests khi vượt limit
3. Process requests theo rate limit
4. Test với 20 requests
*/

// TODO: Rate limiter với channel
type RateLimiter struct {
	// Implement here
	// Sử dụng ticker channel
	ticker *time.Ticker
	requests chan struct{}

}

func NewRateLimiter(requestsPerSecond int) *RateLimiter {
	ticker := time.NewTicker(time.Second / time.Duration(requestsPerSecond))
	requests := make(chan struct{}, requestsPerSecond)
	go func() {
		for range ticker.C {
			select {
			case requests <- struct{}{}:
			default:
			}
		}
	}()
	return &RateLimiter{
		ticker:  ticker,
		requests: requests,
	}
}

func (rl *RateLimiter) Allow() {
	// Implement here
	// Block cho đến khi có slot available
	<-rl.requests
}

// TODO: Test function
func makeRequest(id int) {
	fmt.Printf("Request %d at %v\n", id, time.Now().Format("15:04:05.000"))
	time.Sleep(100 * time.Millisecond) // Simulate work
}

func challenge2() {
	fmt.Println("\n=== CHALLENGE 2: RATE LIMITER ===")

	// TODO: Create rate limiter (5 req/sec)
	// limiter := NewRateLimiter(5)
	limiter := NewRateLimiter(5)

	// TODO: Make 20 requests
	// Observe rate limiting in action
	for i := 1; i <= 20; i++ {
		limiter.Allow()
		go makeRequest(i)
	}

	// Wait to allow all requests to complete
	time.Sleep(5 * time.Second)
}

func main() {
	challenge2()
}
