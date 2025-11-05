package main

import "fmt"

/*
POINTERS - CON TRỎ

Pointer lưu địa chỉ bộ nhớ của một biến
- & : lấy địa chỉ (address of)
- * : dereference (lấy giá trị tại địa chỉ)
*/

func main() {
	// 1. CƠ BẢN VỀ POINTER
	fmt.Println("=== 1. CƠ BẢN VỀ POINTER ===")

	x := 42
	fmt.Printf("x = %d\n", x)
	fmt.Printf("Địa chỉ của x = %p\n", &x)

	// Tạo pointer
	var p *int = &x
	fmt.Printf("p (pointer) = %p\n", p)
	fmt.Printf("*p (giá trị) = %d\n", *p)

	// 2. ZERO VALUE CỦA POINTER
	fmt.Println("\n=== 2. ZERO VALUE ===")

	var nilPointer *int
	fmt.Printf("nil pointer: %v (nil=%t)\n", nilPointer, nilPointer == nil)

	// 3. SỬA GIÁ TRỊ QUA POINTER
	fmt.Println("\n=== 3. SỬA GIÁ TRỊ QUA POINTER ===")

	y := 10
	fmt.Printf("Before: y = %d\n", y)

	py := &y
	*py = 20 // Sửa giá trị qua pointer
	fmt.Printf("After: y = %d\n", y)

	// 4. POINTER TRONG FUNCTION
	fmt.Println("\n=== 4. POINTER TRONG FUNCTION ===")

	// Pass by value
	incrementByValue := func(n int) {
		n++
		fmt.Printf("Inside incrementByValue: %d\n", n)
	}

	// Pass by pointer
	incrementByPointer := func(n *int) {
		*n++
		fmt.Printf("Inside incrementByPointer: %d\n", *n)
	}

	num := 10
	fmt.Printf("Original: %d\n", num)

	incrementByValue(num)
	fmt.Printf("After incrementByValue: %d (không đổi)\n", num)

	incrementByPointer(&num)
	fmt.Printf("After incrementByPointer: %d (đã đổi)\n", num)

	// 5. POINTER VỚI STRUCT
	fmt.Println("\n=== 5. POINTER VỚI STRUCT ===")

	type Person struct {
		Name string
		Age  int
	}

	// Tạo struct bình thường
	p1 := Person{Name: "Alice", Age: 25}
	fmt.Printf("p1: %+v\n", p1)

	// Tạo pointer to struct
	p2 := &Person{Name: "Bob", Age: 30}
	fmt.Printf("p2: %+v\n", p2)

	// Truy cập field (Go tự động dereference)
	p2.Age = 31 // Tương đương (*p2).Age = 31
	fmt.Printf("After update: %+v\n", p2)

	// 6. NEW FUNCTION
	fmt.Println("\n=== 6. NEW FUNCTION ===")

	// new() tạo zero value và trả về pointer
	ptr := new(int)
	fmt.Printf("new(int): %v, value: %d\n", ptr, *ptr)

	*ptr = 100
	fmt.Printf("After assign: %d\n", *ptr)

	// new() với struct
	personPtr := new(Person)
	personPtr.Name = "Charlie"
	personPtr.Age = 35
	fmt.Printf("new(Person): %+v\n", personPtr)

	// 7. POINTER VS VALUE RECEIVER
	fmt.Println("\n=== 7. POINTER VS VALUE RECEIVER ===")

	type Counter struct {
		Count int
	}

	// Value receiver - không thay đổi original
	func (c Counter) IncrementValue() {
		c.Count++
	}

	// Pointer receiver - thay đổi original
	func (c *Counter) IncrementPointer() {
		c.Count++
	}

	counter := Counter{Count: 0}
	fmt.Printf("Initial: %+v\n", counter)

	counter.IncrementValue()
	fmt.Printf("After IncrementValue: %+v (không đổi)\n", counter)

	counter.IncrementPointer()
	fmt.Printf("After IncrementPointer: %+v (đã đổi)\n", counter)

	// 8. SLICE VÀ MAP (ĐÃ LÀ REFERENCE)
	fmt.Println("\n=== 8. SLICE VÀ MAP ===")

	modifySlice := func(s []int) {
		s[0] = 999
		fmt.Printf("Inside function: %v\n", s)
	}

	slice := []int{1, 2, 3, 4, 5}
	fmt.Printf("Before: %v\n", slice)
	modifySlice(slice)
	fmt.Printf("After: %v (đã đổi - slice là reference)\n", slice)

	// 9. POINTER TO POINTER
	fmt.Println("\n=== 9. POINTER TO POINTER ===")

	a := 42
	pa := &a
	ppa := &pa

	fmt.Printf("a = %d\n", a)
	fmt.Printf("*pa = %d\n", *pa)
	fmt.Printf("**ppa = %d\n", **ppa)

	**ppa = 100
	fmt.Printf("After **ppa = 100, a = %d\n", a)

	// 10. ARRAY VS SLICE POINTER
	fmt.Println("\n=== 10. ARRAY VS SLICE POINTER ===")

	// Array passed by value
	modifyArray := func(arr [3]int) {
		arr[0] = 999
		fmt.Printf("Inside modifyArray: %v\n", arr)
	}

	arr := [3]int{1, 2, 3}
	fmt.Printf("Before: %v\n", arr)
	modifyArray(arr)
	fmt.Printf("After: %v (không đổi - array pass by value)\n", arr)

	// Array pointer
	modifyArrayPointer := func(arr *[3]int) {
		arr[0] = 999
		fmt.Printf("Inside modifyArrayPointer: %v\n", arr)
	}

	fmt.Printf("\nBefore pointer: %v\n", arr)
	modifyArrayPointer(&arr)
	fmt.Printf("After pointer: %v (đã đổi)\n", arr)

	// 11. PRACTICAL EXAMPLES
	fmt.Println("\n=== 11. PRACTICAL EXAMPLES ===")

	// Swap function
	swap := func(a, b *int) {
		*a, *b = *b, *a
	}

	x1, y1 := 10, 20
	fmt.Printf("Before swap: x=%d, y=%d\n", x1, y1)
	swap(&x1, &y1)
	fmt.Printf("After swap: x=%d, y=%d\n", x1, y1)

	// Update struct fields
	type User struct {
		Name  string
		Email string
		Age   int
	}

	updateUser := func(u *User, name, email string, age int) {
		u.Name = name
		u.Email = email
		u.Age = age
	}

	user := User{Name: "Old Name", Email: "old@email.com", Age: 20}
	fmt.Printf("\nBefore: %+v\n", user)
	updateUser(&user, "Alice", "alice@example.com", 25)
	fmt.Printf("After: %+v\n", user)

	// 12. LINKED LIST EXAMPLE
	fmt.Println("\n=== 12. LINKED LIST EXAMPLE ===")

	type Node struct {
		Value int
		Next  *Node
	}

	// Tạo linked list: 1 -> 2 -> 3
	node3 := &Node{Value: 3, Next: nil}
	node2 := &Node{Value: 2, Next: node3}
	node1 := &Node{Value: 1, Next: node2}

	// Traverse
	fmt.Print("Linked List: ")
	current := node1
	for current != nil {
		fmt.Printf("%d ", current.Value)
		if current.Next != nil {
			fmt.Print("-> ")
		}
		current = current.Next
	}
	fmt.Println()

	// 13. COMMON MISTAKES
	fmt.Println("\n=== 13. COMMON MISTAKES ===")

	// Mistake 1: Forgetting to initialize
	var uninit *Person
	// uninit.Name = "Test" // PANIC: nil pointer dereference

	// Correct way
	initialized := &Person{}
	initialized.Name = "Test"
	fmt.Printf("Initialized: %+v\n", initialized)

	// Mistake 2: Copying pointer instead of value
	original := Person{Name: "Alice", Age: 25}
	ptrCopy := &original // Both point to same address
	ptrCopy.Age = 30

	fmt.Printf("Original after modify via pointer: %+v\n", original)

	// To get independent copy
	valueCopy := original
	valueCopy.Age = 35
	fmt.Printf("Original after modify copy: %+v\n", original)
	fmt.Printf("Value copy: %+v\n", valueCopy)
}
