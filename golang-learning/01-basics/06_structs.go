package main

import (
	"fmt"
	"time"
)

/*
STRUCTS - CẤU TRÚC DỮ LIỆU

Struct là kiểu dữ liệu tự định nghĩa, nhóm các field có liên quan
Tương tự class trong OOP nhưng đơn giản hơn
*/

// 1. Định nghĩa struct cơ bản
type Person struct {
	Name string
	Age  int
	City string
}

// 2. Struct với tags (dùng cho JSON, validation, etc.)
type User struct {
	ID        int    `json:"id"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

// 3. Struct lồng nhau
type Address struct {
	Street  string
	City    string
	Country string
	ZipCode string
}

type Employee struct {
	Name    string
	Age     int
	Address Address // Embedded struct
	Salary  float64
}

// 4. Embedded struct (inheritance-like)
type Animal struct {
	Name   string
	Weight float64
}

type Dog struct {
	Animal // Embedded - Dog có tất cả fields của Animal
	Breed  string
}

// 5. Anonymous struct
func demonstrateAnonymousStruct() {
	// Tạo struct không cần định nghĩa trước
	person := struct {
		Name string
		Age  int
	}{
		Name: "Alice",
		Age:  25,
	}
	fmt.Printf("Anonymous struct: %+v\n", person)
}

// 6. Empty struct
type EmptyStruct struct{}

func main() {
	// CÁCH 1: Khởi tạo với field names
	fmt.Println("=== 1. KHỞI TẠO STRUCT ===")

	p1 := Person{
		Name: "Alice",
		Age:  25,
		City: "Hanoi",
	}
	fmt.Printf("Person 1: %+v\n", p1)

	// CÁCH 2: Khởi tạo theo thứ tự (không khuyến khích)
	p2 := Person{"Bob", 30, "HCMC"}
	fmt.Printf("Person 2: %+v\n", p2)

	// CÁCH 3: Khởi tạo từng phần
	var p3 Person
	p3.Name = "Charlie"
	p3.Age = 35
	p3.City = "Da Nang"
	fmt.Printf("Person 3: %+v\n", p3)

	// CÁCH 4: Zero value
	var p4 Person
	fmt.Printf("Person 4 (zero value): %+v\n", p4)

	// 2. TRUY CẬP VÀ SỬA FIELD
	fmt.Println("\n=== 2. TRUY CẬP VÀ SỬA FIELD ===")

	fmt.Printf("Name: %s, Age: %d, City: %s\n", p1.Name, p1.Age, p1.City)

	p1.Age = 26
	p1.City = "Hai Phong"
	fmt.Printf("After update: %+v\n", p1)

	// 3. POINTERS TO STRUCTS
	fmt.Println("\n=== 3. POINTERS TO STRUCTS ===")

	p5 := &Person{Name: "David", Age: 40, City: "Hue"}
	fmt.Printf("Pointer to struct: %+v\n", p5)

	// Go tự động dereference
	p5.Age = 41 // Tương đương (*p5).Age = 41
	fmt.Printf("After update: %+v\n", p5)

	// 4. COMPARING STRUCTS
	fmt.Println("\n=== 4. COMPARING STRUCTS ===")

	s1 := Person{Name: "Alice", Age: 25, City: "Hanoi"}
	s2 := Person{Name: "Alice", Age: 25, City: "Hanoi"}
	s3 := Person{Name: "Bob", Age: 30, City: "HCMC"}

	fmt.Printf("s1 == s2: %t\n", s1 == s2)
	fmt.Printf("s1 == s3: %t\n", s1 == s3)

	// 5. STRUCT VỚI TAGS
	fmt.Println("\n=== 5. STRUCT VỚI TAGS ===")

	user := User{
		ID:        1,
		Username:  "alice",
		Email:     "alice@example.com",
		CreatedAt: time.Now(),
	}
	fmt.Printf("User: %+v\n", user)

	// 6. NESTED STRUCTS
	fmt.Println("\n=== 6. NESTED STRUCTS ===")

	emp := Employee{
		Name: "John Doe",
		Age:  35,
		Address: Address{
			Street:  "123 Main St",
			City:    "Hanoi",
			Country: "Vietnam",
			ZipCode: "100000",
		},
		Salary: 50000,
	}

	fmt.Printf("Employee: %+v\n", emp)
	fmt.Printf("Address: %s, %s, %s\n",
		emp.Address.Street, emp.Address.City, emp.Address.Country)

	// 7. EMBEDDED STRUCTS
	fmt.Println("\n=== 7. EMBEDDED STRUCTS ===")

	dog := Dog{
		Animal: Animal{
			Name:   "Buddy",
			Weight: 25.5,
		},
		Breed: "Golden Retriever",
	}

	fmt.Printf("Dog: %+v\n", dog)
	// Có thể truy cập trực tiếp field của Animal
	fmt.Printf("Name: %s, Weight: %.1f, Breed: %s\n",
		dog.Name, dog.Weight, dog.Breed)

	// 8. ANONYMOUS STRUCTS
	fmt.Println("\n=== 8. ANONYMOUS STRUCTS ===")

	demonstrateAnonymousStruct()

	// Sử dụng cho config, temporary data
	config := struct {
		Host string
		Port int
	}{
		Host: "localhost",
		Port: 8080,
	}
	fmt.Printf("Config: %+v\n", config)

	// 9. ARRAY/SLICE OF STRUCTS
	fmt.Println("\n=== 9. ARRAY/SLICE OF STRUCTS ===")

	people := []Person{
		{Name: "Alice", Age: 25, City: "Hanoi"},
		{Name: "Bob", Age: 30, City: "HCMC"},
		{Name: "Charlie", Age: 35, City: "Da Nang"},
	}

	fmt.Println("People:")
	for i, person := range people {
		fmt.Printf("%d: %s (%d) from %s\n",
			i, person.Name, person.Age, person.City)
	}

	// 10. MAP OF STRUCTS
	fmt.Println("\n=== 10. MAP OF STRUCTS ===")

	users := map[string]Person{
		"user1": {Name: "Alice", Age: 25, City: "Hanoi"},
		"user2": {Name: "Bob", Age: 30, City: "HCMC"},
	}

	fmt.Println("Users:")
	for id, user := range users {
		fmt.Printf("%s: %+v\n", id, user)
	}

	// 11. STRUCT COPY
	fmt.Println("\n=== 11. STRUCT COPY ===")

	original := Person{Name: "Alice", Age: 25, City: "Hanoi"}
	copied := original // Copy by value

	copied.Age = 26
	copied.City = "HCMC"

	fmt.Printf("Original: %+v\n", original)
	fmt.Printf("Copied: %+v\n", copied)

	// 12. PRACTICAL EXAMPLES
	fmt.Println("\n=== 12. PRACTICAL EXAMPLES ===")

	// Blog post
	type BlogPost struct {
		ID        int
		Title     string
		Content   string
		Author    string
		Tags      []string
		Published bool
		CreatedAt time.Time
	}

	post := BlogPost{
		ID:      1,
		Title:   "Learning Go",
		Content: "Go is a great language...",
		Author:  "Alice",
		Tags:    []string{"golang", "programming", "tutorial"},
		Published: true,
		CreatedAt: time.Now(),
	}

	fmt.Printf("Blog Post:\n")
	fmt.Printf("  Title: %s\n", post.Title)
	fmt.Printf("  Author: %s\n", post.Author)
	fmt.Printf("  Tags: %v\n", post.Tags)
	fmt.Printf("  Published: %t\n", post.Published)

	// Product with nested struct
	type Product struct {
		ID    int
		Name  string
		Price float64
		Stock int
		Supplier struct {
			Name    string
			Contact string
		}
	}

	product := Product{
		ID:    101,
		Name:  "Laptop",
		Price: 1200.00,
		Stock: 50,
	}
	product.Supplier.Name = "Tech Corp"
	product.Supplier.Contact = "tech@corp.com"

	fmt.Printf("\nProduct:\n")
	fmt.Printf("  Name: %s\n", product.Name)
	fmt.Printf("  Price: $%.2f\n", product.Price)
	fmt.Printf("  Supplier: %s (%s)\n",
		product.Supplier.Name, product.Supplier.Contact)

	// 13. EMPTY STRUCT
	fmt.Println("\n=== 13. EMPTY STRUCT ===")
	
	// Empty struct không chiếm bộ nhớ, dùng cho set, signal
	set := make(map[string]struct{})
	set["apple"] = struct{}{}
	set["banana"] = struct{}{}
	
	fmt.Println("Set:", set)
	_, exists := set["apple"]
	fmt.Printf("'apple' exists: %t\n", exists)
}
