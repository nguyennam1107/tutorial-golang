package main

import (
	"fmt"
	"time"
)

/*
GOROUTINES - CONCURRENT PROGRAMMING

Goroutine là lightweight thread được quản lý bởi Go runtime
- Rất nhẹ (~2KB stack)
- Có thể chạy hàng ngàn goroutines
- Sử dụng keyword "go" để start
*/

func main() {
	// 1. BASIC GOROUTINE
	fmt.Println("=== 1. BASIC GOROUTINE ===")

	// Synchronous execution
	fmt.Println("Before goroutine")
	sayHello("Alice")
	sayHello("Bob")
	fmt.Println("After synchronous calls")

	fmt.Println("\n--- With Goroutines ---")
	fmt.Println("Before goroutine")
	go sayHello("Charlie") // Run in goroutine
	go sayHello("David")   // Run in goroutine
	time.Sleep(100 * time.Millisecond) // Wait for goroutines
	fmt.Println("After goroutines")

	// 2. ANONYMOUS GOROUTINE
	fmt.Println("\n=== 2. ANONYMOUS GOROUTINE ===")

	go func() {
		fmt.Println("Anonymous goroutine 1")
	}()

	go func(name string) {
		fmt.Printf("Anonymous goroutine with param: %s\n", name)
	}("Eve")

	time.Sleep(100 * time.Millisecond)

	// 3. MULTIPLE GOROUTINES
	fmt.Println("\n=== 3. MULTIPLE GOROUTINES ===")

	for i := 0; i < 5; i++ {
		go func(n int) {
			fmt.Printf("Goroutine %d\n", n)
		}(i)
	}

	time.Sleep(100 * time.Millisecond)

	// 4. GOROUTINE VỚI LOOP
	fmt.Println("\n=== 4. GOROUTINE VỚI LOOP ===")

	// Common mistake - closure problem
	fmt.Println("Wrong way (closure problem):")
	for i := 0; i < 3; i++ {
		go func() {
			fmt.Printf("Value: %d\n", i) // i có thể đã thay đổi
		}()
	}
	time.Sleep(100 * time.Millisecond)

	// Correct way - pass as parameter
	fmt.Println("\nCorrect way (pass as param):")
	for i := 0; i < 3; i++ {
		go func(n int) {
			fmt.Printf("Value: %d\n", n)
		}(i)
	}
	time.Sleep(100 * time.Millisecond)

	// 5. COUNTING GOROUTINES
	fmt.Println("\n=== 5. COUNTING GOROUTINES ===")

	count := 0
	done := make(chan bool)

	// Concurrent counter (NOT SAFE - race condition)
	for i := 0; i < 1000; i++ {
		go func() {
			count++
			done <- true
		}()
	}

	// Wait for all goroutines
	for i := 0; i < 1000; i++ {
		<-done
	}
	fmt.Printf("Count: %d (might not be 1000 due to race condition)\n", count)

	// 6. PRACTICAL EXAMPLE - DOWNLOAD FILES
	fmt.Println("\n=== 6. DOWNLOAD SIMULATION ===")

	downloadFile := func(filename string) {
		fmt.Printf("Starting download: %s\n", filename)
		time.Sleep(time.Duration(100+len(filename)*10) * time.Millisecond)
		fmt.Printf("Completed download: %s\n", filename)
	}

	files := []string{"file1.txt", "file2.txt", "file3.txt"}

	// Sequential
	fmt.Println("Sequential download:")
	start := time.Now()
	for _, file := range files {
		downloadFile(file)
	}
	fmt.Printf("Sequential time: %v\n", time.Since(start))

	// Concurrent
	fmt.Println("\nConcurrent download:")
	start = time.Now()
	done2 := make(chan bool)
	for _, file := range files {
		go func(f string) {
			downloadFile(f)
			done2 <- true
		}(file)
	}
	for range files {
		<-done2
	}
	fmt.Printf("Concurrent time: %v\n", time.Since(start))

	// 7. GOROUTINE LEAK PREVENTION
	fmt.Println("\n=== 7. GOROUTINE LEAK ===")

	// Bad - goroutine leak (goroutine chạy mãi)
	leaky := func() {
		go func() {
			// Chờ mãi mãi
			select {}
		}()
	}

	_ = leaky // Don't actually call it

	// Good - with done channel
	goodGoroutine := func() {
		done := make(chan bool)
		go func() {
			select {
			case <-done:
				fmt.Println("Goroutine stopped cleanly")
				return
			}
		}()
		time.Sleep(50 * time.Millisecond)
		done <- true
	}

	goodGoroutine()

	// 8. WORKER POOL PATTERN
	fmt.Println("\n=== 8. WORKER POOL ===")

	jobs := make(chan int, 10)
	results := make(chan int, 10)

	// Start 3 workers
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	// Send 5 jobs
	for j := 1; j <= 5; j++ {
		jobs <- j
	}
	close(jobs)

	// Collect results
	for a := 1; a <= 5; a++ {
		<-results
	}

	// 9. FAN-OUT FAN-IN PATTERN
	fmt.Println("\n=== 9. FAN-OUT FAN-IN ===")

	// Generate numbers
	gen := func(nums ...int) <-chan int {
		out := make(chan int)
		go func() {
			for _, n := range nums {
				out <- n
			}
			close(out)
		}()
		return out
	}

	// Square numbers
	sq := func(in <-chan int) <-chan int {
		out := make(chan int)
		go func() {
			for n := range in {
				out <- n * n
			}
			close(out)
		}()
		return out
	}

	// Merge channels
	merge := func(cs ...<-chan int) <-chan int {
		out := make(chan int)
		for _, c := range cs {
			go func(ch <-chan int) {
				for n := range ch {
					out <- n
				}
			}(c)
		}
		return out
	}

	in := gen(2, 3, 4, 5)

	// Fan-out
	c1 := sq(in)
	c2 := sq(in)

	// Fan-in
	for n := range merge(c1, c2) {
		fmt.Printf("%d ", n)
	}
	fmt.Println()

	// 10. RACE CONDITION EXAMPLE
	fmt.Println("\n=== 10. RACE CONDITION ===")
	fmt.Println("Run with: go run -race 01_goroutines.go")
	fmt.Println("to detect race conditions")

	counter := 0
	done3 := make(chan bool)

	for i := 0; i < 100; i++ {
		go func() {
			counter++ // RACE!
			done3 <- true
		}()
	}

	for i := 0; i < 100; i++ {
		<-done3
	}

	fmt.Printf("Counter: %d (có thể < 100 do race condition)\n", counter)
}

// Helper function
func sayHello(name string) {
	time.Sleep(50 * time.Millisecond)
	fmt.Printf("Hello, %s!\n", name)
}

// Worker function
func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Printf("Worker %d processing job %d\n", id, j)
		time.Sleep(100 * time.Millisecond)
		results <- j * 2
	}
}
