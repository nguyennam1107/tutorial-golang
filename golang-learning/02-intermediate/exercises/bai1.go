package main

import (
	"fmt"
	"time"
)

// ============= BÀI 1: GOROUTINES BASICS =============

// TODO: Tạo function in ra số với goroutine
func printNumber(n int) {
	// Implement here
	fmt.Println(n)
	time.Sleep(100 * time.Millisecond)
	// In ra số n và sleep 100ms
}

// TODO: So sánh sequential vs concurrent
func sequentialExecution() {
	// Implement here
	// Gọi printNumber tuần tự 10 lần
	for i := 1; i <= 10; i++ {
		printNumber(i)
	}
}

func concurrentExecution() {
	// Implement here
	// Gọi printNumber với goroutines 10 lần
	for i := 1; i <= 10; i++ {
		go printNumber(i)
	}
	// Nhớ wait để goroutines hoàn thành
	time.Sleep(2 * time.Second)
}

func exercise1() {
	fmt.Println("=== BÀI 1: GOROUTINES BASICS ===")

	// TODO: Đo thời gian sequential
	start := time.Now()
	sequentialExecution()
	seqTime := time.Since(start)
	fmt.Printf("Sequential time: %v\n", seqTime)

	// TODO: Đo thời gian concurrent
	start = time.Now()
	concurrentExecution()
	concTime := time.Since(start)
	fmt.Printf("Concurrent time: %v\n", concTime)
	fmt.Printf("Speedup: %.2fx\n", float64(seqTime)/float64(concTime))
}

func main() {
	exercise1()
}
