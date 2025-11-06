package main

import (
	"fmt"
	"time"
)

// ============= BÀI 2: CHANNELS =============

// TODO: Producer - gửi numbers vào channel
func producer(ch chan<- int, count int) {
	// Implement here
	for i := 1; i <= count; i++ {
		ch <- i
		time.Sleep(100 * time.Millisecond) 
	}
	close(ch)
}

// TODO: Consumer - nhận numbers từ channel
func consumer(ch <-chan int) {
	// Implement here
	for num := range ch {
		fmt.Println("Received:", num)
	}
}

// TODO: Buffered channel example
func bufferedChannelDemo() {
	// Implement here
	buffered := make(chan string, 5)
	buffered <- "msg1"
	buffered <- "msg2"
	buffered <- "msg3"
	buffered <- "msg4"
	buffered <- "msg5"

	close(buffered)

	for msg := range buffered {
		fmt.Println("Buffered received:", msg)
	}
}

func exercise2() {
	fmt.Println("\n=== BÀI 2: CHANNELS ===")

	// Test producer-consumer
	ch := make(chan int)
	go producer(ch, 10)
	consumer(ch)

	// Test buffered channel
	bufferedChannelDemo()
}

func main() {
	exercise2()
}
