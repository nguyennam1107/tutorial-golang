package main

import "fmt"

// ============= BÀI 6: PIPELINE PATTERN =============

// TODO: Stage 1 - Generate numbers
func generate(nums ...int) <-chan int {
	// Implement here	
	// Tạo channel và goroutine
	// Gửi tất cả nums vào channel
	// Close channel khi done
	// Return channel
	out := make(chan int)
	
	go func() {
		defer close(out)
		for _, n := range nums {
			out <- n
		}
	}()
	return out			
}

// TODO: Stage 2 - Square numbers
func square(in <-chan int) <-chan int {
	// Implement here
	// Nhận từ in channel
	// Square mỗi số
	// Gửi vào out channel
	// Close out channel khi in closed
	// Return out channel
	out := make(chan int)
	go func() {
		defer close(out)
		for n := range in {
			out <- n * n
		}
	}()
	return out
}

// TODO: Stage 3 - Filter even numbers
func filterEven(in <-chan int) <-chan int {
	// Implement here
	// Nhận từ in channel
	// Chỉ gửi số chẵn vào out channel
	// Return out channel
	out := make(chan int)
	go func() {
		defer close(out)
		for n := range in {
			if n%2 == 0 {
				out <- n
			}
		}
	}()
	return out	
}

func exercise6() {
	fmt.Println("\n=== BÀI 6: PIPELINE PATTERN ===")

	// TODO: Build pipeline
	// numbers := generate(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	// squared := square(numbers)
	// evens := filterEven(squared)
	numbers := generate(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	squared := square(numbers)
	evens := filterEven(squared)
	// TODO: Consume results
	// for num := range evens {
	//     fmt.Println(num)
	// }
	for num := range evens {
		fmt.Println(num)
	}
}

func main() {
	exercise6()
}
