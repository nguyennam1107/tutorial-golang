package main

import (
	"context"
	"fmt"
	"time"
)

// ============= BÀI 8: CONTEXT =============

// TODO: Context với timeout
func contextWithTimeout() {
	// Implement here
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	done := make(chan struct{})
	go func() {
		time.Sleep(2 * time.Second)
		done <- struct{}{}
	}()

	select {
	case <-done:
		fmt.Println("Task completed")
	case <-ctx.Done():
		fmt.Println("Task cancelled:", ctx.Err())
	}
}

// TODO: Context với cancel
func contextWithCancel() {
	// Implement here
	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("Goroutine stopped:", ctx.Err())
				return
			default:
				fmt.Println("Goroutine is running")
				time.Sleep(500 * time.Millisecond)
			}
		}
	}()

	time.Sleep(2 * time.Second)
	cancel()
	time.Sleep(1 * time.Second) 
}

// TODO: Context với values
func contextWithValues() {
	// Implement here
	ctx := context.WithValue(context.Background(), "userID", 123)
	ctx = context.WithValue(ctx, "requestID", "abc-123")
	processRequest(ctx)
}

func processRequest(ctx context.Context) {
	// TODO: Lấy values từ context
	// userID := ctx.Value("userID")
	// requestID := ctx.Value("requestID")
	// In ra values
	userID := ctx.Value("userID")
	requestID := ctx.Value("requestID")
	fmt.Printf("Processing request %s for user %d\n", requestID, userID)
}

func exercise8() {
	fmt.Println("\n=== BÀI 8: CONTEXT ===")

	fmt.Println("Context with timeout:")
	contextWithTimeout()

	fmt.Println("\nContext with cancel:")
	contextWithCancel()

	fmt.Println("\nContext with values:")
	contextWithValues()
}

func main() {
	exercise8()
}
