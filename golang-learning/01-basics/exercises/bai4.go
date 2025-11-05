package main

import (
	"fmt"
)	
// ============= BÀI 4: ARRAYS & SLICES =============

// TODO: Reverse slice
func reverseSlice(slice []int) []int {
	// Implement here
	for i, j := 0, len(slice)-1; i < j; i, j = i+1, j-1 {
		slice[i], slice[j] = slice[j], slice[i]
	}
	return slice
}
// TODO: Remove element tại index
func removeAt(slice []int, index int) []int {
	// Implement here
	if index < 0 || index >= len(slice) {
		return slice
	}
	return append(slice[:index], slice[index+1:]...)
}

// TODO: Insert element tại index
func insertAt(slice []int, index, value int) []int {
	// Implement here
	if index < 0 || index > len(slice) {
		return slice
	}
	return append(append(slice[:index], value), slice[index:]...)
	return slice
}

// TODO: Merge 2 sorted slices
func mergeSorted(a, b []int) []int {
	// Implement here
	merged := make([]int, 0, len(a)+len(b))
	i, j := 0, 0
	for i < len(a) && j < len(b) {
		if a[i] < b[j] {
			merged = append(merged, a[i])
			i++
		} else {
			merged = append(merged, b[j])
			j++
		}
	}
	// Append any remaining elements from either slice
	merged = append(merged, a[i:]...)
	merged = append(merged, b[j:]...)
	return merged
}

// TODO: Tìm phần tử xuất hiện nhiều nhất
func mostFrequent(slice []int) int {
	// Implement here
	frequency := make(map[int]int)
	for _, v := range slice {
		frequency[v]++
	}
	maxFreq := 0
	mostFreqElem := 0
	for k, v := range frequency {
		if v > maxFreq {
			maxFreq = v
			mostFreqElem = k
		}
	}
	return mostFreqElem
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
func main() {
	exercise4()
}