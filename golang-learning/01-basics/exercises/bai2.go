package main
import (
	"fmt"
)

// ============= BÀI 2: FUNCTIONS =============

// TODO: Viết hàm tính giai thừa
func factorial(n int) int {
	// Implement here
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

// TODO: Viết hàm kiểm tra số nguyên tố
func isPrime(n int) bool {
	// Implement here
	if n < 2 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

// TODO: Viết hàm Fibonacci
func fibonacci(n int) int {
	// Implement here
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

// TODO: Viết hàm tìm min trong variadic parameters
func min(numbers ...int) int {
	// Implement here
	if len(numbers) == 0 {
		return 0 // Or handle error
	}
	minVal := numbers[0]
	for _, num := range numbers {
		if num < minVal {
			minVal = num
		}
	}
	return minVal
}

func exercise2() {
	fmt.Println("\n=== BÀI 2: FUNCTIONS ===")

	// Test factorial
	fmt.Printf("factorial(5) = %d (expected: 120)\n", factorial(5))

	// Test isPrime
	fmt.Printf("isPrime(17) = %t (expected: true)\n", isPrime(17))
	fmt.Printf("isPrime(20) = %t (expected: false)\n", isPrime(20))

	// Test fibonacci
	fmt.Printf("fibonacci(10) = %d (expected: 55)\n", fibonacci(10))

	// Test min
	fmt.Printf("min(5,2,8,1,9) = %d (expected: 1)\n", min(5, 2, 8, 1, 9))
}
func main() {
	exercise2()
}