package main

import (
	"fmt"
	"sync"
)

// ============= BÀI 5: SYNC PACKAGE =============

// TODO: WaitGroup example
func waitGroupExample() {
	// Implement here
	var wg sync.WaitGroup
	numGoroutines := 5
	wg.Add(numGoroutines)

	for i := 1; i <= numGoroutines; i++ {
		go func(id int) {
			defer wg.Done()
			fmt.Printf("Goroutine %d is running\n", id)
		}(i)
	}

	wg.Wait()
	fmt.Println("All goroutines completed.")
}

// TODO: Mutex để fix race condition
var (
	counter int
	mu      sync.Mutex
)

func incrementWithoutMutex(wg *sync.WaitGroup) {
	// Implement here
	for i := 0; i < 1000; i++ {
		counter++
	}
	wg.Done()
}

func incrementWithMutex(wg *sync.WaitGroup) {
	// Implement here
	for i := 0; i < 1000; i++ {
		mu.Lock()
		counter++
		mu.Unlock()
	}
	wg.Done()
}

func exercise5() {
	fmt.Println("\n=== BÀI 5: SYNC PACKAGE ===")

	fmt.Println("WaitGroup example:")
	waitGroupExample()

	// Test WITHOUT mutex (race condition)
	fmt.Println("\nWithout Mutex (có race condition):")
	counter = 0
	var wg1 sync.WaitGroup
	// TODO: Start 10 goroutines gọi incrementWithoutMutex
	wg1.Wait()
	fmt.Printf("Final counter (without mutex): %d\n", counter)

	// Test WITH mutex (safe)
	fmt.Println("\nWith Mutex (safe):")
	counter = 0
	var wg2 sync.WaitGroup
	// TODO: Start 10 goroutines gọi incrementWithMutex
	for i := 1; i <= 10; i++ {
		wg2.Add(1)
		go incrementWithMutex(&wg2)
	}
	wg2.Wait()
	fmt.Printf("Final counter (with mutex): %d\n", counter)
}

func main() {
	exercise5()
}
