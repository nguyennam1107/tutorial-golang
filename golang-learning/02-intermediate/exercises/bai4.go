package main

import (
	"fmt"
	"time"
)

// ============= BÀI 4: SELECT STATEMENT =============

// TODO: Select từ nhiều channels
func multipleChannelsSelect() {
	// Implement here
	c1 := make(chan string)
	c2 := make(chan string)

	// Goroutine 1: gửi "message 1" sau 1 giây
	go func() {
		time.Sleep(1 * time.Second)
		c1 <- "message 1"
	}()
	
	// Goroutine 2: gửi "message 2" sau 2 giây
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "message 2"
	}()
	
	// Sử dụng select để nhận từ c1 hoặc c2
	// Nhận 2 messages
	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("Received", msg1)
		case msg2 := <-c2:
			fmt.Println("Received", msg2)
		}
	}
}

// TODO: Timeout với select
func selectWithTimeout() {
	// Implement here
	ch := make(chan string)

	// Goroutine gửi message sau 2 giây
	go func() {
		time.Sleep(2 * time.Second)
		ch <- "result"
	}()
	
	// Select với timeout 1 giây
	// In "Timeout!" nếu timeout
	// In message nếu nhận được
	select {
	case res := <-ch:
		fmt.Println(res)
	case <-time.After(1 * time.Second):
		fmt.Println("Timeout!")
	}
}

// TODO: Non-blocking select
func nonBlockingSelect() {
	// Implement here
	messages := make(chan string)
	signals := make(chan bool)

	// Try receive từ messages (không block)
	// Try send vào signals (không block)
	// Sử dụng default case
	select {
	case msg := <-messages:
		fmt.Println("Received:", msg)
	default:
		fmt.Println("No message received")
	}

	select {
	case signals <- true:
		fmt.Println("Sent signal")
	default:
		fmt.Println("No signal sent")
	}
}

func exercise4() {
	fmt.Println("\n=== BÀI 4: SELECT STATEMENT ===")

	fmt.Println("Multiple channels:")
	multipleChannelsSelect()

	fmt.Println("\nTimeout:")
	selectWithTimeout()

	fmt.Println("\nNon-blocking:")
	nonBlockingSelect()
}

func main() {
	exercise4()
}
