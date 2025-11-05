package main

import (
	"fmt"
	"sync"
)

// ============= BÀI 7: FAN-OUT FAN-IN =============

// TODO: Generate numbers
func generate7(nums ...int) <-chan int {
	// Implement here
	return nil
}

// TODO: Square numbers (worker)
func square7(in <-chan int) <-chan int {
	// Implement here
	// Mỗi worker process từ in channel
	return nil
}

// TODO: Merge nhiều channels thành 1 (fan-in)
func merge(channels ...<-chan int) <-chan int {
	// Implement here
	var wg sync.WaitGroup
	out := make(chan int)

	// TODO: Start goroutine cho mỗi channel
	// Đọc từ channel và gửi vào out
	
	// TODO: Close out khi tất cả channels done
	
	return out
}

func exercise7() {
	fmt.Println("\n=== BÀI 7: FAN-OUT FAN-IN ===")

	// TODO: Setup pipeline
	// in := generate7(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)

	// TODO: Fan-out - start 3 workers
	// c1 := square7(in)
	// c2 := square7(in)
	// c3 := square7(in)

	// TODO: Fan-in - merge results
	// results := merge(c1, c2, c3)

	// TODO: Consume results
	// for num := range results {
	//     fmt.Println(num)
	// }
}

func main() {
	exercise7()
}
