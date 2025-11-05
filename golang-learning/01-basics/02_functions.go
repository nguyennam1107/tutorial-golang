package main

import (
	"fmt"
	"strings"
)

/*
FUNCTIONS TRONG GO

1. Cú pháp cơ bản:
   func functionName(param1 type1, param2 type2) returnType {
       // code
   }

2. Đặc điểm:
   - Multiple return values
   - Named return values
   - Variadic functions
   - Anonymous functions
   - Closures
*/

// 1. Function cơ bản
func greet(name string) string {
	return "Hello, " + name + "!"
}

// 2. Function với nhiều tham số
func add(a int, b int) int {
	return a + b
}

// 3. Nhiều tham số cùng kiểu (shorthand)
func multiply(a, b, c int) int {
	return a * b * c
}

// 4. Multiple return values
func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("cannot divide by zero")
	}
	return a / b, nil
}

// 5. Named return values
func rectangle(width, height float64) (area, perimeter float64) {
	area = width * height
	perimeter = 2 * (width + height)
	return // naked return
}

// 6. Variadic function (tham số không giới hạn)
func sum(numbers ...int) int {
	total := 0
	for _, num := range numbers {
		total += num
	}
	return total
}

// 7. Function làm tham số
func applyOperation(a, b int, operation func(int, int) int) int {
	return operation(a, b)
}

// 8. Function trả về function (closure)
func makeMultiplier(factor int) func(int) int {
	return func(n int) int {
		return n * factor
	}
}

// 9. Defer - thực thi sau khi function kết thúc
func demonstrateDefer() {
	defer fmt.Println("3. Defer: Thực thi cuối cùng")
	fmt.Println("1. Bắt đầu function")
	fmt.Println("2. Giữa function")
}

// 10. Recursive function
func factorial(n int) int {
	if n <= 1 {
		return 1
	}
	return n * factorial(n-1)
}

func main() {
	fmt.Println("=== 1. Function Cơ Bản ===")
	message := greet("Nam")
	fmt.Println(message)

	fmt.Println("\n=== 2. Function Với Nhiều Tham Số ===")
	result := add(5, 3)
	fmt.Printf("5 + 3 = %d\n", result)

	fmt.Println("\n=== 3. Shorthand Parameters ===")
	product := multiply(2, 3, 4)
	fmt.Printf("2 × 3 × 4 = %d\n", product)

	fmt.Println("\n=== 4. Multiple Return Values ===")
	quotient, err := divide(10, 2)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("10 ÷ 2 = %.2f\n", quotient)
	}

	_, err = divide(10, 0)
	if err != nil {
		fmt.Println("Error:", err)
	}

	fmt.Println("\n=== 5. Named Return Values ===")
	area, perimeter := rectangle(5, 3)
	fmt.Printf("Rectangle(5×3): Area=%.1f, Perimeter=%.1f\n", area, perimeter)

	fmt.Println("\n=== 6. Variadic Function ===")
	fmt.Printf("sum(1,2,3) = %d\n", sum(1, 2, 3))
	fmt.Printf("sum(1,2,3,4,5) = %d\n", sum(1, 2, 3, 4, 5))

	numbers := []int{10, 20, 30, 40}
	fmt.Printf("sum(10,20,30,40) = %d\n", sum(numbers...))

	fmt.Println("\n=== 7. Function Làm Tham Số ===")
	addFunc := func(x, y int) int { return x + y }
	subFunc := func(x, y int) int { return x - y }

	fmt.Printf("applyOperation(10, 5, add) = %d\n", applyOperation(10, 5, addFunc))
	fmt.Printf("applyOperation(10, 5, sub) = %d\n", applyOperation(10, 5, subFunc))

	fmt.Println("\n=== 8. Anonymous Function & Closure ===")
	// Anonymous function
	func(msg string) {
		fmt.Println("Anonymous:", msg)
	}("Hello from anonymous function")

	// Closure
	double := makeMultiplier(2)
	triple := makeMultiplier(3)
	fmt.Printf("double(5) = %d\n", double(5))
	fmt.Printf("triple(5) = %d\n", triple(5))

	fmt.Println("\n=== 9. Defer ===")
	demonstrateDefer()

	fmt.Println("\n=== 10. Recursive Function ===")
	fmt.Printf("factorial(5) = %d\n", factorial(5))
	fmt.Printf("factorial(10) = %d\n", factorial(10))

	fmt.Println("\n=== 11. Function với String Processing ===")
	processString := func(s string, ops ...func(string) string) string {
		result := s
		for _, op := range ops {
			result = op(result)
		}
		return result
	}

	text := "  Hello World  "
	processed := processString(text,
		strings.TrimSpace,
		strings.ToUpper,
		func(s string) string { return strings.ReplaceAll(s, " ", "_") },
	)
	fmt.Printf("Original: '%s'\n", text)
	fmt.Printf("Processed: '%s'\n", processed)
}
