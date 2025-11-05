package main
import (
	"fmt"
)
// ============= BÀI 3: CONTROL FLOW =============

// TODO: In bảng cửu chương từ 1-10
func multiplicationTable() {
	// Implement here
	for i := 1; i <= 10; i++ {
		for j := 1; j <= 10; j++ {
			fmt.Printf("%d×%d=%2d  ", i, j, i*j)
		}
		fmt.Println()
	}
}

// TODO: FizzBuzz từ 1-100
func fizzBuzz() {
	// Implement here
	for i := 1; i <= 100; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("FizzBuzz")
		} else if i%3 == 0 {
			fmt.Println("Fizz")
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}
	}
}

// TODO: Tìm số hoàn hảo < 10000
// Số hoàn hảo: tổng các ước (không kể chính nó) = chính nó
// Ví dụ: 6 = 1 + 2 + 3
func findPerfectNumbers() {
	for n := 1; n < 10000; n++ {
		sum := 0
		for i := 1; i < n; i++ {
			if n%i == 0 {
				sum += i
			}
		}
		if sum == n {
			fmt.Println(n)
		}
	}
}

func exercise3() {
	fmt.Println("\n=== BÀI 3: CONTROL FLOW ===")

	fmt.Println("Bảng cửu chương:")
	multiplicationTable()
	
	fmt.Println("\nFizzBuzz:")
	fizzBuzz()

	fmt.Println("\nPerfect numbers:")
	findPerfectNumbers()
}
func main () {
	exercise3()
}