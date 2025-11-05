package main

import (
	"fmt"
	"time"
)

/*
CHANNELS - COMMUNICATION GIỮA GOROUTINES

Channel là typed conduit để gửi/nhận giá trị giữa các goroutines
- ch <- value  : gửi value vào channel
- value := <-ch : nhận value từ channel
*/

func main() {
	// 1. UNBUFFERED CHANNEL
	fmt.Println("=== 1. UNBUFFERED CHANNEL ===")

	ch := make(chan string)

	go func() {
		ch <- "Hello from goroutine"
	}()

	msg := <-ch
	fmt.Println("Received:", msg)

	// 2. BUFFERED CHANNEL
	fmt.Println("\n=== 2. BUFFERED CHANNEL ===")

	buffered := make(chan int, 2) // capacity = 2

	buffered <- 1
	buffered <- 2
	// buffered <- 3 // Would block without goroutine

	fmt.Println("Received:", <-buffered)
	fmt.Println("Received:", <-buffered)

	// 3. CHANNEL DIRECTIONS
	fmt.Println("\n=== 3. CHANNEL DIRECTIONS ===")

	// Send-only channel
	sendOnly := func(ch chan<- string) {
		ch <- "message"
	}

	// Receive-only channel
	receiveOnly := func(ch <-chan string) {
		msg := <-ch
		fmt.Println("Received:", msg)
	}

	ch2 := make(chan string)
	go sendOnly(ch2)
	receiveOnly(ch2)

	// 4. CLOSING CHANNELS
	fmt.Println("\n=== 4. CLOSING CHANNELS ===")

	jobs := make(chan int, 5)
	done := make(chan bool)

	go func() {
		for {
			j, more := <-jobs
			if more {
				fmt.Println("Received job:", j)
			} else {
				fmt.Println("All jobs received")
				done <- true
				return
			}
		}
	}()

	for j := 1; j <= 3; j++ {
		jobs <- j
	}
	close(jobs)
	<-done

	// 5. RANGE OVER CHANNEL
	fmt.Println("\n=== 5. RANGE OVER CHANNEL ===")

	queue := make(chan string, 3)
	queue <- "one"
	queue <- "two"
	queue <- "three"
	close(queue)

	for elem := range queue {
		fmt.Println(elem)
	}

	// 6. SELECT STATEMENT
	fmt.Println("\n=== 6. SELECT STATEMENT ===")

	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		c1 <- "one"
	}()

	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "two"
	}()

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("Received", msg1)
		case msg2 := <-c2:
			fmt.Println("Received", msg2)
		}
	}

	// 7. TIMEOUT WITH SELECT
	fmt.Println("\n=== 7. TIMEOUT ===")

	c3 := make(chan string)

	go func() {
		time.Sleep(2 * time.Second)
		c3 <- "result"
	}()

	select {
	case res := <-c3:
		fmt.Println(res)
	case <-time.After(1 * time.Second):
		fmt.Println("Timeout!")
	}

	// 8. NON-BLOCKING SELECT
	fmt.Println("\n=== 8. NON-BLOCKING SELECT ===")

	messages := make(chan string)
	signals := make(chan bool)

	select {
	case msg := <-messages:
		fmt.Println("Received message:", msg)
	default:
		fmt.Println("No message received")
	}

	msg2 := "hi"
	select {
	case messages <- msg2:
		fmt.Println("Sent message:", msg2)
	default:
		fmt.Println("No message sent")
	}

	select {
	case msg := <-messages:
		fmt.Println("Received:", msg)
	case sig := <-signals:
		fmt.Println("Received signal:", sig)
	default:
		fmt.Println("No activity")
	}

	// 9. PIPELINE PATTERN
	fmt.Println("\n=== 9. PIPELINE ===")

	// Stage 1: Generate numbers
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

	// Stage 2: Square numbers
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

	// Pipeline
	for n := range sq(gen(2, 3, 4)) {
		fmt.Println(n)
	}

	// 10. WORKER POOL
	fmt.Println("\n=== 10. WORKER POOL ===")

	numJobs := 5
	jobs2 := make(chan int, numJobs)
	results := make(chan int, numJobs)

	// Start workers
	for w := 1; w <= 3; w++ {
		go worker(w, jobs2, results)
	}

	// Send jobs
	for j := 1; j <= numJobs; j++ {
		jobs2 <- j
	}
	close(jobs2)

	// Collect results
	for a := 1; a <= numJobs; a++ {
		fmt.Println("Result:", <-results)
	}
}

func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Printf("Worker %d processing job %d\n", id, j)
		time.Sleep(time.Second)
		results <- j * 2
	}
}
