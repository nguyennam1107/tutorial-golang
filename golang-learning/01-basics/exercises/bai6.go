package main

import (
	"fmt"
	"sort"
)
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
	return fmt.Sprintf("Name: %s, Age: %d, Email: %s", p.Name, p.Age, p.Email)
}

// TODO: Sort persons by age
func sortByAge(people []Person) {
	sort.Slice(people, func(i, j int) bool {
		return people[i].Age < people[j].Age
	})
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

func main() {
	exercise6()
}