package main

import (
	"fmt"
	"sort"
)

/*
ARRAYS VÀ SLICES

1. Array: Kích thước cố định
   var arr [5]int

2. Slice: Kích thước động, linh hoạt hơn
   var slice []int
   slice := make([]int, length, capacity)
*/

func main() {
	// 1. ARRAYS - Mảng cố định
	fmt.Println("=== 1. ARRAYS ===")

	// Khai báo array
	var numbers [5]int
	fmt.Println("Zero value array:", numbers)

	// Khởi tạo với giá trị
	primes := [5]int{2, 3, 5, 7, 11}
	fmt.Println("Prime numbers:", primes)

	// Compiler tự đếm độ dài
	weekdays := [...]string{"Mon", "Tue", "Wed", "Thu", "Fri"}
	fmt.Println("Weekdays:", weekdays)
	fmt.Println("Length:", len(weekdays))

	// Truy cập phần tử
	fmt.Printf("First day: %s, Last day: %s\n", weekdays[0], weekdays[4])

	// Sửa giá trị
	numbers[0] = 10
	numbers[4] = 50
	fmt.Println("Modified array:", numbers)

	// Lặp qua array
	fmt.Println("\nLặp qua array:")
	for i, v := range primes {
		fmt.Printf("primes[%d] = %d\n", i, v)
	}

	// 2. SLICES - Mảng động
	fmt.Println("\n=== 2. SLICES ===")

	// Tạo slice từ array
	arr := [5]int{1, 2, 3, 4, 5}
	slice1 := arr[1:4] // Lấy từ index 1 đến 3
	fmt.Println("Array:", arr)
	fmt.Println("Slice [1:4]:", slice1)

	// Tạo slice trực tiếp
	colors := []string{"Red", "Green", "Blue"}
	fmt.Println("Colors slice:", colors)

	// Tạo slice với make
	numbers2 := make([]int, 5)       // length = 5, capacity = 5
	numbers3 := make([]int, 3, 10)   // length = 3, capacity = 10
	fmt.Printf("make([]int, 5): %v (len=%d, cap=%d)\n",
		numbers2, len(numbers2), cap(numbers2))
	fmt.Printf("make([]int, 3, 10): %v (len=%d, cap=%d)\n",
		numbers3, len(numbers3), cap(numbers3))

	// 3. APPEND - Thêm phần tử
	fmt.Println("\n=== 3. APPEND ===")

	fruits := []string{"Apple", "Banana"}
	fmt.Println("Original:", fruits)

	fruits = append(fruits, "Cherry")
	fmt.Println("After append:", fruits)

	fruits = append(fruits, "Durian", "Elderberry")
	fmt.Println("Append multiple:", fruits)

	// Append slice vào slice
	moreFruits := []string{"Fig", "Grape"}
	fruits = append(fruits, moreFruits...)
	fmt.Println("Append slice:", fruits)

	// 4. COPY - Sao chép slice
	fmt.Println("\n=== 4. COPY ===")

	src := []int{1, 2, 3, 4, 5}
	dst := make([]int, len(src))
	copied := copy(dst, src)

	fmt.Printf("Source: %v\n", src)
	fmt.Printf("Destination: %v\n", dst)
	fmt.Printf("Elements copied: %d\n", copied)

	// Copy một phần
	partial := make([]int, 3)
	copy(partial, src)
	fmt.Printf("Partial copy: %v\n", partial)

	// 5. SLICING - Cắt slice
	fmt.Println("\n=== 5. SLICING ===")

	nums := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println("Original:", nums)
	fmt.Println("nums[2:5]:", nums[2:5])   // [2 3 4]
	fmt.Println("nums[:4]:", nums[:4])     // [0 1 2 3]
	fmt.Println("nums[6:]:", nums[6:])     // [6 7 8 9]
	fmt.Println("nums[:]:", nums[:])       // toàn bộ

	// 6. 2D SLICES - Slice đa chiều
	fmt.Println("\n=== 6. 2D SLICES ===")

	// Ma trận 3x3
	matrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	fmt.Println("Matrix:")
	for i, row := range matrix {
		fmt.Printf("Row %d: %v\n", i, row)
	}

	// 7. SLICE OPERATIONS
	fmt.Println("\n=== 7. SLICE OPERATIONS ===")

	// Remove element at index
	removeAt := func(slice []int, index int) []int {
		return append(slice[:index], slice[index+1:]...)
	}

	data := []int{10, 20, 30, 40, 50}
	fmt.Println("Original:", data)
	data = removeAt(data, 2)
	fmt.Println("After removing index 2:", data)

	// Insert element at index
	insertAt := func(slice []int, index, value int) []int {
		slice = append(slice[:index], append([]int{value}, slice[index:]...)...)
		return slice
	}

	data = insertAt(data, 2, 35)
	fmt.Println("After inserting 35 at index 2:", data)

	// Filter
	filter := func(slice []int, condition func(int) bool) []int {
		result := []int{}
		for _, v := range slice {
			if condition(v) {
				result = append(result, v)
			}
		}
		return result
	}

	evens := filter(data, func(n int) bool { return n%2 == 0 })
	fmt.Println("Even numbers:", evens)

	// 8. SORTING
	fmt.Println("\n=== 8. SORTING ===")

	unsorted := []int{5, 2, 8, 1, 9, 3}
	fmt.Println("Unsorted:", unsorted)

	sort.Ints(unsorted)
	fmt.Println("Sorted:", unsorted)

	names := []string{"Charlie", "Alice", "Bob"}
	fmt.Println("Unsorted names:", names)
	sort.Strings(names)
	fmt.Println("Sorted names:", names)

	// Custom sort
	people := []struct {
		Name string
		Age  int
	}{
		{"Alice", 25},
		{"Bob", 30},
		{"Charlie", 20},
	}

	sort.Slice(people, func(i, j int) bool {
		return people[i].Age < people[j].Age
	})

	fmt.Println("\nSorted by age:")
	for _, p := range people {
		fmt.Printf("%s: %d\n", p.Name, p.Age)
	}

	// 9. NIL SLICES
	fmt.Println("\n=== 9. NIL SLICES ===")

	var nilSlice []int
	fmt.Printf("Nil slice: %v (len=%d, cap=%d, nil=%t)\n",
		nilSlice, len(nilSlice), cap(nilSlice), nilSlice == nil)

	// Có thể append vào nil slice
	nilSlice = append(nilSlice, 1, 2, 3)
	fmt.Println("After append:", nilSlice)

	// 10. PRACTICAL EXAMPLES
	fmt.Println("\n=== 10. PRACTICAL EXAMPLES ===")

	// Reverse slice
	reverse := func(s []int) []int {
		for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
			s[i], s[j] = s[j], s[i]
		}
		return s
	}

	test := []int{1, 2, 3, 4, 5}
	fmt.Println("Original:", test)
	reverse(test)
	fmt.Println("Reversed:", test)

	// Unique elements
	unique := func(slice []int) []int {
		seen := make(map[int]bool)
		result := []int{}
		for _, v := range slice {
			if !seen[v] {
				seen[v] = true
				result = append(result, v)
			}
		}
		return result
	}

	duplicates := []int{1, 2, 2, 3, 3, 3, 4, 5, 5}
	fmt.Println("With duplicates:", duplicates)
	fmt.Println("Unique:", unique(duplicates))
}
