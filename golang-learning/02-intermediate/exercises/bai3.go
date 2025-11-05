package main

import (
	"fmt"
	"time"
)

// ============= BÀI 3: WORKER POOL =============

// TODO: Worker function
func worker(id int, jobs <-chan int, results chan<- int) {
	// Implement here
	// Nhận jobs từ channel
	// Process: square the number
	// Gửi result vào results channel
	// In ra: "Worker {id} processing job {job}"
	for j := range jobs {
		fmt.Printf("Worker %d processing job %d\n", id, j)
		time.Sleep(time.Second) 
		results <- j * j
	}
}

// TODO: Create worker pool
func createWorkerPool(numWorkers int, jobs <-chan int, results chan<- int) {
	// Implement here
	// Start numWorkers goroutines
	// Mỗi worker gọi worker()
	for i := 0; i < numWorkers; i++ {
		go worker(i, jobs, results)
	}
}

func exercise3() {
	fmt.Println("\n=== BÀI 3: WORKER POOL ===")

	// TODO: Setup
	numJobs := 20
	numWorkers := 5

	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	// TODO: Start worker pool
	createWorkerPool(numWorkers, jobs, results)

	// TODO: Send jobs
	// Gửi 20 jobs vào channel
	for i := 1; i <= numJobs; i++ {
		jobs <- i
	}
	// TODO: Close jobs channel
	close(jobs)
	// TODO: Collect results
	// Nhận và in ra 20 results
	for i := 1; i <= numJobs; i++ {
		result := <-results
		fmt.Printf("Result of job %d: %d\n", i, result)
	}
}

func main() {
	exercise3()
}
