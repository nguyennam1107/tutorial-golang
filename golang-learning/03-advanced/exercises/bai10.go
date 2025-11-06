package main

import (
	"fmt"
	"context"
	"sync"
	"time"
)

// ============= BÀI 10: RATE LIMITER =============

// TODO: Token Bucket Rate Limiter
type RateLimiter struct {
	rate   int           // tokens per second
	burst  int           // max tokens
	tokens int           // current tokens
	mu     sync.Mutex    // protect tokens
	ticker *time.Ticker  // refill ticker
}

func NewRateLimiter(rate int, burst int) *RateLimiter {
	// Implement here
	rl := &RateLimiter{
		rate:   rate,
		burst:  burst,
		tokens: burst,
		ticker: time.NewTicker(time.Second / time.Duration(rate)),
	}
	go func() {
		for range rl.ticker.C {
			rl.mu.Lock()
			if rl.tokens < rl.burst {
				rl.tokens++
			}
			rl.mu.Unlock()
		}
	}()
	return rl
}

func (rl *RateLimiter) Allow() bool {
	// Implement here
	rl.mu.Lock()
	defer rl.mu.Unlock()
	if rl.tokens > 0 {
		rl.tokens--
		return true
	}
	return false
}

func (rl *RateLimiter) Wait(ctx context.Context) error {
	// Implement here
	for {
		rl.mu.Lock()
		if rl.tokens > 0 {
			rl.tokens--
			rl.mu.Unlock()
			return nil
		}
		rl.mu.Unlock()

		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-time.After(time.Second / time.Duration(rl.rate)):
		}
	}
	return nil
}

func exercise10() {
	fmt.Println("\n=== BÀI 10: RATE LIMITER ===")

	// TODO: Test rate limiter
	limiter := NewRateLimiter(5, 10)

	for i := 0; i < 20; i++ {
		if limiter.Allow() {
			fmt.Printf("Request %d: ALLOWED at %v\n", i, time.Now())
		} else {
			fmt.Printf("Request %d: DENIED\n", i)
		}
		time.Sleep(100 * time.Millisecond)
	}
}
func main() {
	exercise10()
}