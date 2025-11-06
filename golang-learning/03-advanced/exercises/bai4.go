package main

import (
	"fmt"
)
// ============= BÀI 4: GENERIC FUNCTIONS =============

// TODO: Map function
func Map[T, U any](slice []T, fn func(T) U) []U {
	// Implement here
	var result []U
	for _, item := range slice {
		result = append(result, fn(item))
	}
	return result
}

// TODO: Filter function
func Filter[T any](slice []T, fn func(T) bool) []T {
	var result []T
	for _, item := range slice {
		if fn(item) {
			result = append(result, item)
		}
	}
	return result
}

// TODO: Reduce function
func Reduce[T, U any](slice []T, initial U, fn func(U, T) U) U {
	// Implement here
	var zero U
	for _, item := range slice {
		zero = fn(zero, item)
	}
	return zero
}
// TODO: Find function
func Find[T any](slice []T, fn func(T) bool) (T, bool) {
	// Implement here
	var zero T
	for _, item := range slice {
		if fn(item) {
			return item, true
		}
	}
	return zero, false
}

// TODO: Contains function
func Contains[T comparable](slice []T, item T) bool {
	// Implement here
	for _, v := range slice {
		if v == item {
			return true
		}
	}
	return false
}

// TODO: Reverse function
func Reverse[T any](slice []T) []T {
	var result []T
	for i := len(slice) - 1; i >= 0; i-- {
		result = append(result, slice[i])
	}
	return result
}

// TODO: Min function
func Min[T interface{ ~int | ~float64 | ~string }](items ...T) T {
	// Implement here
	var zero T
	for _, item := range items {
		if item < zero {
			zero = item
		}
	}
	return zero
}

// TODO: Max function
func Max[T interface{ ~int | ~float64 | ~string }](items ...T) T {
	var zero T
	for _, item := range items {
		if item > zero {
			zero = item
		}
	}
	return zero
}
func Avg[T interface{ ~int | ~float64 }](items ...T) T {
	// Implement here
	var sum T
	for _, item := range items {
		sum += item
	}
	return sum / T(len(items))
}

func exercise4() {
	fmt.Println("\n=== BÀI 4: GENERIC FUNCTIONS ===")

	nums := []int{1, 2, 3, 4, 5}

	// TODO: Test Map
	squared := Map(nums, func(n int) int { return n * n })
	fmt.Println("Squared:", squared)

	// TODO: Test Filter
	evens := Filter(nums, func(n int) bool { return n%2 == 0 })
	fmt.Println("Evens:", evens)

	// TODO: Test Reduce
	sum := Reduce(nums, 0, func(acc, n int) int { return acc + n })
	fmt.Println("Sum:", sum)

	// TODO: Test Contains
	fmt.Println("Contains 3:", Contains(nums, 3))
	fmt.Println("Contains 10:", Contains(nums, 10))

	// TODO: Test Min/Max
	fmt.Println("Min:", Min(1, 2, 3, 4, 5))
	fmt.Println("Max:", Max(1, 2, 3, 4, 5))

	fmt.Println("Avg:", Avg(1, 2, 3, 4, 5))
}

func main() {
	exercise4()
}