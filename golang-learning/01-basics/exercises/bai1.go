package main

import (
	"fmt"
)

// ============= BÀI 1: VARIABLES =============

func exercise1() {
	fmt.Println("=== BÀI 1: VARIABLES ===")

	// TODO: Khai báo biến name (string), age (int), height (float64)
	var (
		name   string  = "Nguyen Nam"
		age    int     = 25
		height float64 = 1.75
	)
	// TODO: In ra màn hình
	fmt.Printf("Name: %s, Age: %d, Height: %.2f\n", name, age, height)
	// TODO: Convert age sang float64 và nhân với 1.5
	ageFloat := float64(age) * 1.5
	fmt.Printf("Age sau khi nhân với 1.5: %.2f\n", ageFloat)
}
func main() {
	exercise1()
}