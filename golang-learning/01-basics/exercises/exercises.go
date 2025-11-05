package main

import (
	"fmt"
	"math"
	"sort"
	"strings"
)

// ============= BÀI 1: VARIABLES =============

func exercise1() {
	fmt.Println("=== BÀI 1: VARIABLES ===")

	// TODO: Khai báo biến name (string), age (int), height (float64)
	// TODO: In ra màn hình
	// TODO: Convert age sang float64 và nhân với 1.5
}

// ============= BÀI 2: FUNCTIONS =============

// TODO: Viết hàm tính giai thừa
func factorial(n int) int {
	// Implement here
	return 0
}

// TODO: Viết hàm kiểm tra số nguyên tố
func isPrime(n int) bool {
	// Implement here
	return false
}

// TODO: Viết hàm Fibonacci
func fibonacci(n int) int {
	// Implement here
	return 0
}

// TODO: Viết hàm tìm min trong variadic parameters
func min(numbers ...int) int {
	// Implement here
	return 0
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

// ============= BÀI 3: CONTROL FLOW =============

// TODO: In bảng cửu chương từ 1-10
func multiplicationTable() {
	// Implement here
}

// TODO: FizzBuzz từ 1-100
func fizzBuzz() {
	// Implement here
}

// TODO: Tìm số hoàn hảo < 10000
// Số hoàn hảo: tổng các ước (không kể chính nó) = chính nó
// Ví dụ: 6 = 1 + 2 + 3
func findPerfectNumbers() {
	// Implement here
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

// ============= BÀI 4: ARRAYS & SLICES =============

// TODO: Reverse slice
func reverseSlice(slice []int) []int {
	// Implement here
	return slice
}

// TODO: Remove element tại index
func removeAt(slice []int, index int) []int {
	// Implement here
	return slice
}

// TODO: Insert element tại index
func insertAt(slice []int, index, value int) []int {
	// Implement here
	return slice
}

// TODO: Merge 2 sorted slices
func mergeSorted(a, b []int) []int {
	// Implement here
	return []int{}
}

// TODO: Tìm phần tử xuất hiện nhiều nhất
func mostFrequent(slice []int) int {
	// Implement here
	return 0
}

func exercise4() {
	fmt.Println("\n=== BÀI 4: ARRAYS & SLICES ===")

	// Test reverse
	nums := []int{1, 2, 3, 4, 5}
	fmt.Printf("Original: %v\n", nums)
	fmt.Printf("Reversed: %v\n", reverseSlice(append([]int{}, nums...)))

	// Test remove
	fmt.Printf("Remove index 2: %v\n", removeAt(append([]int{}, nums...), 2))

	// Test insert
	fmt.Printf("Insert 99 at index 2: %v\n", insertAt(append([]int{}, nums...), 2, 99))

	// Test merge
	a := []int{1, 3, 5, 7}
	b := []int{2, 4, 6, 8}
	fmt.Printf("Merge %v and %v: %v\n", a, b, mergeSorted(a, b))

	// Test most frequent
	data := []int{1, 2, 2, 3, 3, 3, 4}
	fmt.Printf("Most frequent in %v: %d\n", data, mostFrequent(data))
}

// ============= BÀI 5: MAPS =============

// TODO: Đếm tần suất ký tự
func charFrequency(s string) map[rune]int {
	// Implement here
	return map[rune]int{}
}

// TODO: Group students theo class
type Student struct {
	Name  string
	Class string
	Score int
}

func groupByClass(students []Student) map[string][]Student {
	// Implement here
	return map[string][]Student{}
}

// TODO: Fibonacci với cache
func fibonacciWithCache() func(int) int {
	// Implement here using closure and map
	return func(n int) int {
		return 0
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

// ============= BÀI 6: STRUCTS =============

// TODO: Implement Person struct
type Person struct {
	Name  string
	Age   int
	Email string
}

// TODO: Implement String() method
func (p Person) String() string {
	// Implement here
	return ""
}

// TODO: Sort persons by age
func sortByAge(people []Person) {
	// Implement here using sort.Slice
}

func exercise6() {
	fmt.Println("\n=== BÀI 6: STRUCTS ===")

	people := []Person{
		{"Alice", 25, "alice@example.com"},
		{"Bob", 30, "bob@example.com"},
		{"Charlie", 20, "charlie@example.com"},
	}

	fmt.Println("Before sort:")
	for _, p := range people {
		fmt.Println(p)
	}

	sortByAge(people)

	fmt.Println("\nAfter sort by age:")
	for _, p := range people {
		fmt.Println(p)
	}
}

// ============= BÀI 7: METHODS & INTERFACES =============

// TODO: Define Shape interface
type Shape interface {
	// Add methods here
}

// TODO: Implement Circle
type Circle struct {
	Radius float64
}

// TODO: Implement Rectangle
type Rectangle struct {
	Width  float64
	Height float64
}

// TODO: Implement Triangle
type Triangle struct {
	Base   float64
	Height float64
}

// TODO: Tính tổng diện tích
func totalArea(shapes []Shape) float64 {
	// Implement here
	return 0
}

func exercise7() {
	fmt.Println("\n=== BÀI 7: METHODS & INTERFACES ===")

	shapes := []Shape{
		Circle{Radius: 5},
		Rectangle{Width: 10, Height: 5},
		Triangle{Base: 6, Height: 4},
	}

	fmt.Printf("Total area: %.2f\n", totalArea(shapes))
}

// ============= BÀI 8: ERROR HANDLING =============

// TODO: Validate email
func validateEmail(email string) error {
	// Check: không rỗng, có @, có .
	// Implement here
	return nil
}

// TODO: Custom error type
type ValidationError struct {
	Field   string
	Message string
}

func (e ValidationError) Error() string {
	// Implement here
	return ""
}

// TODO: Validate user với wrapped errors
func validateUser(name, email string, age int) error {
	// Implement here
	return nil
}

func exercise8() {
	fmt.Println("\n=== BÀI 8: ERROR HANDLING ===")

	emails := []string{
		"valid@example.com",
		"invalid",
		"",
		"no-at-sign.com",
	}

	for _, email := range emails {
		if err := validateEmail(email); err != nil {
			fmt.Printf("'%s': %v\n", email, err)
		} else {
			fmt.Printf("'%s': Valid ✓\n", email)
		}
	}
}
