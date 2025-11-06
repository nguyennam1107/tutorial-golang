package main

import (
	"fmt"
	"time"
)

// ============= BÀI 3: WORKER POOL =============

// TODO: Worker function
func worker(id int, jobs <-chan int, results chan<- int) {
	// Implement here
	for j := range jobs {
		fmt.Printf("Worker %d processing job %d\n", id, j)
		time.Sleep(time.Second) 
		results <- j * j
	}
}

// TODO: Create worker pool
func createWorkerPool(numWorkers int, jobs <-chan int, results chan<- int) {
	// Implement here
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
	for i := 1; i <= numJobs; i++ {
		jobs <- i
	}
	// TODO: Close jobs channel
	close(jobs)
	// TODO: Collect results
	for i := 1; i <= numJobs; i++ {
		result := <-results
		fmt.Printf("Result of job %d: %d\n", i, result)
	}
}

func main() {
	exercise3()
}
