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
	// Tạo context với timeout 1 giây
	
	// Goroutine thực hiện task mất 2 giây
	
	// Select giữa done channel và context.Done()
}

// TODO: Context với cancel
func contextWithCancel() {
	// Implement here
	// Tạo context với cancel
	
	// Goroutine chạy vô hạn
	// Lắng nghe ctx.Done()
	
	// Cancel sau 2 giây
}

// TODO: Context với values
func contextWithValues() {
	// Implement here
	// Tạo context với values
	ctx := context.WithValue(context.Background(), "userID", 123)
	ctx = context.WithValue(ctx, "requestID", "abc-123")
	
	// Function nhận context và in values
	processRequest(ctx)
}

func processRequest(ctx context.Context) {
	// TODO: Lấy values từ context
	// userID := ctx.Value("userID")
	// requestID := ctx.Value("requestID")
	// In ra values
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
