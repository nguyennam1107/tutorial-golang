package main

import "fmt"

/*
BIẾN VÀ KIỂU DỮ LIỆU TRONG GO

1. Khai báo biến:
   - var name type = value
   - var name = value (type inference)
   - name := value (short declaration, chỉ trong function)

2. Các kiểu dữ liệu cơ bản:
   - bool
   - string
   - int, int8, int16, int32, int64
   - uint, uint8, uint16, uint32, uint64
   - float32, float64
   - complex64, complex128
   - byte (alias for uint8)
   - rune (alias for int32, đại diện cho Unicode code point)
*/

func main() {
	// 1. Khai báo biến với var
	var name string = "Nguyen Nam"
	var age int = 25
	var isStudent bool = true

	fmt.Println("=== Khai báo với var ===")
	fmt.Printf("Name: %s, Age: %d, Student: %t\n", name, age, isStudent)

	// 2. Type inference (Go tự suy ra kiểu)
	var city = "Ha Noi"
	var height = 1.75

	fmt.Println("\n=== Type Inference ===")
	fmt.Printf("City: %s, Height: %.2f\n", city, height)

	// 3. Short declaration (ngắn gọn hơn)
	email := "nam@example.com"
	score := 95.5

	fmt.Println("\n=== Short Declaration ===")
	fmt.Printf("Email: %s, Score: %.1f\n", email, score)

	// 4. Khai báo nhiều biến cùng lúc
	var (
		firstName = "Nam"
		lastName  = "Nguyen"
		yearBorn  = 1998
	)

	fmt.Println("\n=== Multiple Declaration ===")
	fmt.Printf("Full Name: %s %s, Year: %d\n", firstName, lastName, yearBorn)

	// 5. Constants (hằng số)
	const Pi = 3.14159
	const AppName = "Go Learning"

	fmt.Println("\n=== Constants ===")
	fmt.Printf("Pi: %.5f, App: %s\n", Pi, AppName)

	// 6. Zero values (giá trị mặc định)
	var defaultInt int
	var defaultFloat float64
	var defaultBool bool
	var defaultString string

	fmt.Println("\n=== Zero Values ===")
	fmt.Printf("int: %d, float: %.1f, bool: %t, string: '%s'\n",
		defaultInt, defaultFloat, defaultBool, defaultString)

	// 7. Type conversion (chuyển đổi kiểu)
	var x int = 42
	var y float64 = float64(x)
	var z uint = uint(x)

	fmt.Println("\n=== Type Conversion ===")
	fmt.Printf("int: %d, float64: %.1f, uint: %d\n", x, y, z)

	// 8. Các kiểu số khác nhau
	var (
		num8  int8   = 127  // -128 to 127
		num16 int16  = 32767
		num32 int32  = 2147483647
		num64 int64  = 9223372036854775807
		unum  uint   = 4294967295
		f32   float32 = 3.14
		f64   float64 = 3.141592653589793
	)

	fmt.Println("\n=== Các Kiểu Số ===")
	fmt.Printf("int8: %d, int16: %d, int32: %d\n", num8, num16, num32)
	fmt.Printf("int64: %d, uint: %d\n", num64, unum)
	fmt.Printf("float32: %.2f, float64: %.15f\n", f32, f64)

	// 9. Rune và Byte
	var letter rune = 'A'    // Unicode character
	var byteValue byte = 65  // ASCII value of 'A'

	fmt.Println("\n=== Rune & Byte ===")
	fmt.Printf("Rune: %c (%d), Byte: %c (%d)\n", letter, letter, byteValue, byteValue)

	// 10. Complex numbers
	var complex1 complex64 = 1 + 2i
	var complex2 complex128 = 3 + 4i

	fmt.Println("\n=== Complex Numbers ===")
	fmt.Printf("complex64: %v, complex128: %v\n", complex1, complex2)
}
