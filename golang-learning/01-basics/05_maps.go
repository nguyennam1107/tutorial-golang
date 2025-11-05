package main

import "fmt"

/*
MAPS - BẢN ĐỒ (Key-Value)

Map là cấu trúc dữ liệu lưu trữ cặp key-value
- Key phải là kiểu so sánh được (comparable)
- Value có thể là bất kỳ kiểu nào
- Không có thứ tự
*/

func main() {
	// 1. KHAI BÁO MAP
	fmt.Println("=== 1. KHAI BÁO MAP ===")

	// Map literal
	ages := map[string]int{
		"Alice":   25,
		"Bob":     30,
		"Charlie": 35,
	}
	fmt.Println("Ages:", ages)

	// Tạo map rỗng với make
	scores := make(map[string]int)
	fmt.Println("Empty map:", scores)

	// Nil map (không thể thêm phần tử)
	var nilMap map[string]int
	fmt.Printf("Nil map: %v (nil=%t)\n", nilMap, nilMap == nil)

	// 2. THÊM VÀ SỬA
	fmt.Println("\n=== 2. THÊM VÀ SỬA ===")

	scores["Math"] = 90
	scores["English"] = 85
	scores["Science"] = 88
	fmt.Println("Scores:", scores)

	// Sửa giá trị
	scores["Math"] = 95
	fmt.Println("After update:", scores)

	// 3. TRUY CẬP PHẦN TỬ
	fmt.Println("\n=== 3. TRUY CẬP PHẦN TỬ ===")

	mathScore := scores["Math"]
	fmt.Printf("Math score: %d\n", mathScore)

	// Kiểm tra key có tồn tại không
	historyScore, exists := scores["History"]
	if exists {
		fmt.Printf("History score: %d\n", historyScore)
	} else {
		fmt.Println("History score not found")
	}

	// Truy cập key không tồn tại -> zero value
	physicsScore := scores["Physics"]
	fmt.Printf("Physics score (not exists): %d\n", physicsScore)

	// 4. XÓA PHẦN TỬ
	fmt.Println("\n=== 4. XÓA PHẦN TỬ ===")

	fmt.Println("Before delete:", scores)
	delete(scores, "English")
	fmt.Println("After delete English:", scores)

	// Xóa key không tồn tại - không lỗi
	delete(scores, "NotExist")

	// 5. DUYỆT MAP
	fmt.Println("\n=== 5. DUYỆT MAP ===")

	countries := map[string]string{
		"VN": "Vietnam",
		"US": "United States",
		"JP": "Japan",
		"KR": "Korea",
	}

	fmt.Println("Countries:")
	for code, name := range countries {
		fmt.Printf("%s: %s\n", code, name)
	}

	// Chỉ lấy keys
	fmt.Println("\nCountry codes:")
	for code := range countries {
		fmt.Printf("%s ", code)
	}
	fmt.Println()

	// Chỉ lấy values
	fmt.Println("\nCountry names:")
	for _, name := range countries {
		fmt.Printf("%s ", name)
	}
	fmt.Println()

	// 6. MAP VỚI STRUCT
	fmt.Println("\n=== 6. MAP VỚI STRUCT ===")

	type Person struct {
		Name string
		Age  int
		City string
	}

	people := map[string]Person{
		"p1": {Name: "Alice", Age: 25, City: "Hanoi"},
		"p2": {Name: "Bob", Age: 30, City: "HCMC"},
	}

	fmt.Println("People:")
	for id, person := range people {
		fmt.Printf("%s: %s (%d) from %s\n",
			id, person.Name, person.Age, person.City)
	}

	// 7. NESTED MAPS
	fmt.Println("\n=== 7. NESTED MAPS ===")

	// Map of maps
	studentGrades := map[string]map[string]int{
		"Alice": {
			"Math":    90,
			"English": 85,
		},
		"Bob": {
			"Math":    88,
			"English": 92,
		},
	}

	fmt.Println("Student grades:")
	for student, grades := range studentGrades {
		fmt.Printf("%s:\n", student)
		for subject, grade := range grades {
			fmt.Printf("  %s: %d\n", subject, grade)
		}
	}

	// 8. MAP VỚI SLICE
	fmt.Println("\n=== 8. MAP VỚI SLICE ===")

	// Map với value là slice
	hobbies := map[string][]string{
		"Alice": {"Reading", "Swimming"},
		"Bob":   {"Coding", "Gaming", "Music"},
	}

	fmt.Println("Hobbies:")
	for person, hobbyList := range hobbies {
		fmt.Printf("%s: %v\n", person, hobbyList)
	}

	// Thêm hobby
	hobbies["Alice"] = append(hobbies["Alice"], "Traveling")
	fmt.Println("After adding hobby:", hobbies)

	// 9. MAP OPERATIONS
	fmt.Println("\n=== 9. MAP OPERATIONS ===")

	// Length
	fmt.Printf("Number of scores: %d\n", len(scores))

	// Copy map (phải copy thủ công)
	original := map[string]int{"a": 1, "b": 2}
	copied := make(map[string]int)
	for k, v := range original {
		copied[k] = v
	}
	fmt.Println("Original:", original)
	fmt.Println("Copied:", copied)

	// Merge maps
	map1 := map[string]int{"a": 1, "b": 2}
	map2 := map[string]int{"c": 3, "d": 4}
	for k, v := range map2 {
		map1[k] = v
	}
	fmt.Println("Merged map:", map1)

	// 10. PRACTICAL EXAMPLES
	fmt.Println("\n=== 10. PRACTICAL EXAMPLES ===")

	// Word counter
	text := "go is great go is simple go go go"
	words := []string{}
	word := ""
	for _, char := range text + " " {
		if char == ' ' {
			if word != "" {
				words = append(words, word)
				word = ""
			}
		} else {
			word += string(char)
		}
	}

	wordCount := make(map[string]int)
	for _, w := range words {
		wordCount[w]++
	}

	fmt.Println("Word count:")
	for word, count := range wordCount {
		fmt.Printf("%s: %d\n", word, count)
	}

	// Character frequency
	charFreq := make(map[rune]int)
	testString := "hello world"
	for _, char := range testString {
		if char != ' ' {
			charFreq[char]++
		}
	}

	fmt.Println("\nCharacter frequency:")
	for char, freq := range charFreq {
		fmt.Printf("%c: %d\n", char, freq)
	}

	// Group by
	type Student struct {
		Name  string
		Class string
		Score int
	}

	students := []Student{
		{"Alice", "A", 90},
		{"Bob", "B", 85},
		{"Charlie", "A", 88},
		{"David", "B", 92},
	}

	groupedByClass := make(map[string][]Student)
	for _, student := range students {
		groupedByClass[student.Class] = append(
			groupedByClass[student.Class],
			student,
		)
	}

	fmt.Println("\nStudents grouped by class:")
	for class, studentList := range groupedByClass {
		fmt.Printf("Class %s:\n", class)
		for _, s := range studentList {
			fmt.Printf("  %s (score: %d)\n", s.Name, s.Score)
		}
	}

	// Cache/Memoization example
	fibonacci := func() func(int) int {
		cache := make(map[int]int)
		var fib func(int) int
		fib = func(n int) int {
			if n <= 1 {
				return n
			}
			if val, exists := cache[n]; exists {
				return val
			}
			cache[n] = fib(n-1) + fib(n-2)
			return cache[n]
		}
		return fib
	}

	fib := fibonacci()
	fmt.Println("\nFibonacci with memoization:")
	for i := 0; i <= 10; i++ {
		fmt.Printf("fib(%d) = %d\n", i, fib(i))
	}
}
