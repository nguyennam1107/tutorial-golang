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
}

func NewRateLimiter(requestsPerSecond int) *RateLimiter {
	// Implement here
	return nil
}

func (rl *RateLimiter) Allow() {
	// Implement here
	// Block cho đến khi có slot available
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

	// TODO: Make 20 requests
	// Observe rate limiting in action
}

func main() {
	challenge2()
}
