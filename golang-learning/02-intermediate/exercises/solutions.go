package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// ĐÁP ÁN CHO TẤT CẢ BÀI TẬP INTERMEDIATE

// ============= BÀI 1: GOROUTINES =============

func printNumber1(n int) {
	fmt.Printf("%d ", n)
	time.Sleep(100 * time.Millisecond)
}

func sequentialExecution1() {
	for i := 1; i <= 10; i++ {
		printNumber1(i)
	}
	fmt.Println()
}

func concurrentExecution1() {
	var wg sync.WaitGroup
	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			printNumber1(n)
		}(i)
	}
	wg.Wait()
	fmt.Println()
}

// ============= BÀI 2: CHANNELS =============

func producer2(ch chan<- int, count int) {
	for i := 1; i <= count; i++ {
		ch <- i
	}
	close(ch)
}

func consumer2(ch <-chan int) {
	for num := range ch {
		fmt.Printf("%d ", num)
	}
	fmt.Println()
}

func bufferedChannelDemo2() {
	ch := make(chan string, 5)
	
	// Send without blocking
	ch <- "one"
	ch <- "two"
	ch <- "three"
	ch <- "four"
	ch <- "five"
	close(ch)
	
	// Receive
	for msg := range ch {
		fmt.Println(msg)
	}
}

// ============= BÀI 3: WORKER POOL =============

func worker3(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Printf("Worker %d processing job %d\n", id, j)
		time.Sleep(100 * time.Millisecond)
		results <- j * j
	}
}

func createWorkerPool3(numWorkers int, jobs <-chan int, results chan<- int) {
	for w := 1; w <= numWorkers; w++ {
		go worker3(w, jobs, results)
	}
}

// ============= BÀI 4: SELECT =============

func multipleChannelsSelect4() {
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		c1 <- "message 1"
	}()

	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "message 2"
	}()

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("Received:", msg1)
		case msg2 := <-c2:
			fmt.Println("Received:", msg2)
		}
	}
}

func selectWithTimeout4() {
	ch := make(chan string)

	go func() {
		time.Sleep(2 * time.Second)
		ch <- "result"
	}()

	select {
	case res := <-ch:
		fmt.Println("Received:", res)
	case <-time.After(1 * time.Second):
		fmt.Println("Timeout!")
	}
}

// ============= BÀI 5: SYNC =============

func waitGroupExample5() {
	var wg sync.WaitGroup
	
	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			fmt.Printf("Goroutine %d\n", n)
		}(i)
	}
	
	wg.Wait()
	fmt.Println("All goroutines done")
}

var (
	counter5 int
	mu5      sync.Mutex
)

func incrementWithMutex5(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 1000; i++ {
		mu5.Lock()
		counter5++
		mu5.Unlock()
	}
}

// ============= BÀI 6: PIPELINE =============

func generate6(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	return out
}

func square6(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()
	return out
}

func filterEven6(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			if n%2 == 0 {
				out <- n
			}
		}
		close(out)
	}()
	return out
}

// ============= BÀI 7: FAN-OUT FAN-IN =============

func merge7(channels ...<-chan int) <-chan int {
	var wg sync.WaitGroup
	out := make(chan int)

	output := func(c <-chan int) {
		defer wg.Done()
		for n := range c {
			out <- n
		}
	}

	wg.Add(len(channels))
	for _, c := range channels {
		go output(c)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

// ============= BÀI 8: CONTEXT =============

func contextWithTimeout8() {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	done := make(chan bool)
	go func() {
		time.Sleep(2 * time.Second)
		done <- true
	}()

	select {
	case <-done:
		fmt.Println("Task completed")
	case <-ctx.Done():
		fmt.Println("Timeout:", ctx.Err())
	}
}

func contextWithCancel8() {
	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("Goroutine stopped:", ctx.Err())
				return
			default:
				fmt.Println("Working...")
				time.Sleep(500 * time.Millisecond)
			}
		}
	}()

	time.Sleep(2 * time.Second)
	cancel()
	time.Sleep(500 * time.Millisecond) // Wait for cleanup
}

// ============= MAIN =============

func main() {
	fmt.Println("╔════════════════════════════════════════════╗")
	fmt.Println("║   INTERMEDIATE SOLUTIONS - PREVIEW        ║")
	fmt.Println("╚════════════════════════════════════════════╝\n")

	fmt.Println("=== BÀI 1: GOROUTINES ===")
	start := time.Now()
	sequentialExecution1()
	fmt.Printf("Sequential: %v\n", time.Since(start))
	
	start = time.Now()
	concurrentExecution1()
	fmt.Printf("Concurrent: %v\n", time.Since(start))

	fmt.Println("\n=== BÀI 2: CHANNELS ===")
	ch := make(chan int)
	go producer2(ch, 5)
	consumer2(ch)

	fmt.Println("\n=== BÀI 4: SELECT ===")
	multipleChannelsSelect4()

	fmt.Println("\n=== BÀI 5: SYNC ===")
	waitGroupExample5()

	fmt.Println("\n=== BÀI 6: PIPELINE ===")
	numbers := generate6(1, 2, 3, 4, 5)
	squared := square6(numbers)
	evens := filterEven6(squared)
	for n := range evens {
		fmt.Printf("%d ", n)
	}
	fmt.Println()

	fmt.Println("\n=== BÀI 8: CONTEXT ===")
	contextWithTimeout8()

	fmt.Println("\n╔════════════════════════════════════════════╗")
	fmt.Println("║          ✅ SOLUTIONS COMPLETED!          ║")
	fmt.Println("╚════════════════════════════════════════════╝")
}
