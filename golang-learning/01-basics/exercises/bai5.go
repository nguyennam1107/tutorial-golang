package main

import (
	"fmt"
)
// ============= BÀI 5: MAPS =============

// TODO: Đếm tần suất ký tự
func charFrequency(s string) map[rune]int {
	// Implement here
	frequency := make(map[rune]int)
	for _, char := range s {
		frequency[char]++
	}
	return frequency
}

// TODO: Group students theo class
type Student struct {
	Name  string
	Class string
	Score int
}
func groupByClass(students []Student) map[string][]Student {
	// Implement here
	grouped := make(map[string][]Student)
	for _, student := range students {
		grouped[student.Class] = append(grouped[student.Class], student)
	}
	return grouped
}

// TODO: Fibonacci với cache
func fibonacciWithCache() func(int) int {
	cache := make(map[int]int)
	return func(n int) int {
		if n <= 1 {
			return n
		}
		if val, found := cache[n]; found {
			return val
		}
		result := fibonacciWithCache()(n - 1) + fibonacciWithCache()(n - 2)
		cache[n] = result
		return result
	}
}

func exercise5() {
	fmt.Println("\n=== BÀI 5: MAPS ===")

	// Test char frequency
	text := "hello world"
	freq := charFrequency(text)
	fmt.Printf("Char frequency of '%s': %v\n", text, freq)

	// Test group by class
	students := []Student{
		{"Alice", "A", 90},
		{"Bob", "B", 85},
		{"Charlie", "A", 88},
	}
	grouped := groupByClass(students)
	fmt.Printf("Grouped: %v\n", grouped)

	// Test fibonacci with cache
	fib := fibonacciWithCache()
	fmt.Printf("Fib(10) = %d\n", fib(10))
	fmt.Printf("Fib(20) = %d\n", fib(20))
}

func main() {
	exercise5()
}