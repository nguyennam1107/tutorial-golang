package main

import (
	"fmt"
	"time"
)

/*
CONTROL FLOW - LUỒNG ĐIỀU KHIỂN

1. if/else
2. switch
3. for loop
4. break, continue
5. goto (ít dùng)
*/

func main() {
	// 1. IF/ELSE
	fmt.Println("=== 1. IF/ELSE ===")

	age := 20
	if age >= 18 {
		fmt.Println("Bạn đã trưởng thành")
	} else {
		fmt.Println("Bạn chưa trưởng thành")
	}

	// If với short statement
	if score := 85; score >= 90 {
		fmt.Println("Điểm A")
	} else if score >= 80 {
		fmt.Println("Điểm B")
	} else if score >= 70 {
		fmt.Println("Điểm C")
	} else {
		fmt.Println("Cần cố gắng hơn")
	}

	// 2. SWITCH
	fmt.Println("\n=== 2. SWITCH ===")

	// Switch thông thường
	day := time.Now().Weekday()
	fmt.Println("Hôm nay là:", day)

	switch day {
	case time.Monday:
		fmt.Println("Đầu tuần, hãy tập trung!")
	case time.Friday:
		fmt.Println("Cuối tuần rồi!")
	case time.Saturday, time.Sunday:
		fmt.Println("Nghỉ ngơi thôi!")
	default:
		fmt.Println("Ngày thường trong tuần")
	}

	// Switch không điều kiện (thay thế if-else dài)
	hour := time.Now().Hour()
	switch {
	case hour < 12:
		fmt.Println("Chào buổi sáng!")
	case hour < 18:
		fmt.Println("Chào buổi chiều!")
	default:
		fmt.Println("Chào buổi tối!")
	}

	// Switch với type
	checkType := func(i interface{}) {
		switch v := i.(type) {
		case int:
			fmt.Printf("%v là integer\n", v)
		case string:
			fmt.Printf("%v là string\n", v)
		case bool:
			fmt.Printf("%v là boolean\n", v)
		default:
			fmt.Printf("%v là kiểu khác\n", v)
		}
	}

	fmt.Println("\n=== Switch Type ===")
	checkType(42)
	checkType("hello")
	checkType(true)
	checkType(3.14)

	// 3. FOR LOOP
	fmt.Println("\n=== 3. FOR LOOP ===")

	// For cơ bản
	fmt.Println("For cơ bản:")
	for i := 0; i < 5; i++ {
		fmt.Printf("%d ", i)
	}
	fmt.Println()

	// For như while
	fmt.Println("\nFor như while:")
	count := 0
	for count < 5 {
		fmt.Printf("%d ", count)
		count++
	}
	fmt.Println()

	// For vô hạn (với break)
	fmt.Println("\nFor vô hạn với break:")
	n := 0
	for {
		if n >= 5 {
			break
		}
		fmt.Printf("%d ", n)
		n++
	}
	fmt.Println()

	// For range với slice
	fmt.Println("\nFor range với slice:")
	fruits := []string{"Apple", "Banana", "Cherry"}
	for index, fruit := range fruits {
		fmt.Printf("%d: %s\n", index, fruit)
	}

	// For range chỉ lấy value
	fmt.Println("\nChỉ lấy value:")
	for _, fruit := range fruits {
		fmt.Printf("- %s\n", fruit)
	}

	// For range với map
	fmt.Println("\nFor range với map:")
	scores := map[string]int{
		"Math":    90,
		"English": 85,
		"Science": 88,
	}
	for subject, score := range scores {
		fmt.Printf("%s: %d\n", subject, score)
	}

	// For range với string (runes)
	fmt.Println("\nFor range với string:")
	text := "Xin chào"
	for index, char := range text {
		fmt.Printf("%d: %c\n", index, char)
	}

	// 4. BREAK và CONTINUE
	fmt.Println("\n=== 4. BREAK và CONTINUE ===")

	fmt.Println("Continue - bỏ qua số chẵn:")
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue // Bỏ qua số chẵn
		}
		fmt.Printf("%d ", i)
	}
	fmt.Println()

	fmt.Println("\nBreak - dừng khi gặp 5:")
	for i := 0; i < 10; i++ {
		if i == 5 {
			break // Dừng hẳn vòng lặp
		}
		fmt.Printf("%d ", i)
	}
	fmt.Println()

	// Label với break/continue
	fmt.Println("\nLabel với nested loop:")
outer:
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if i == 1 && j == 1 {
				fmt.Println("\nBreak outer loop at i=1, j=1")
				break outer
			}
			fmt.Printf("(%d,%d) ", i, j)
		}
	}
	fmt.Println()

	// 5. NESTED LOOPS
	fmt.Println("\n=== 5. NESTED LOOPS ===")
	fmt.Println("Bảng cửu chương 2-5:")
	for i := 2; i <= 5; i++ {
		for j := 1; j <= 5; j++ {
			fmt.Printf("%d×%d=%2d  ", i, j, i*j)
		}
		fmt.Println()
	}

	// 6. LOOP với CHANNEL (preview cho concurrent programming)
	fmt.Println("\n=== 6. PRACTICAL EXAMPLES ===")

	// Tìm số nguyên tố
	isPrime := func(n int) bool {
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

	fmt.Println("Số nguyên tố < 30:")
	for i := 2; i < 30; i++ {
		if isPrime(i) {
			fmt.Printf("%d ", i)
		}
	}
	fmt.Println()

	// Fibonacci
	fmt.Println("\n10 số Fibonacci đầu tiên:")
	a, b := 0, 1
	for i := 0; i < 10; i++ {
		fmt.Printf("%d ", a)
		a, b = b, a+b
	}
	fmt.Println()

	// FizzBuzz
	fmt.Println("\nFizzBuzz (1-20):")
	for i := 1; i <= 20; i++ {
		switch {
		case i%15 == 0:
			fmt.Print("FizzBuzz ")
		case i%3 == 0:
			fmt.Print("Fizz ")
		case i%5 == 0:
			fmt.Print("Buzz ")
		default:
			fmt.Printf("%d ", i)
		}
	}
	fmt.Println()
}
