package main

import (
	"fmt"
	"math"
	"strings"
)

/*
METHODS VÀ INTERFACES

1. Methods: Functions gắn với type (receiver)
2. Interfaces: Định nghĩa behavior, không phải implementation
*/

// ========== METHODS ==========

// Struct để demo methods
type Rectangle struct {
	Width  float64
	Height float64
}

type Circle struct {
	Radius float64
}

// Method với value receiver
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

// Method với pointer receiver (có thể sửa struct)
func (r *Rectangle) Scale(factor float64) {
	r.Width *= factor
	r.Height *= factor
}

// Methods cho Circle
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

// ========== INTERFACES ==========

// Interface định nghĩa behavior
type Shape interface {
	Area() float64
	Perimeter() float64
}

// Hàm nhận interface
func PrintShapeInfo(s Shape) {
	fmt.Printf("Area: %.2f, Perimeter: %.2f\n", s.Area(), s.Perimeter())
}

// Empty interface (interface{} hoặc any)
func PrintAnything(value interface{}) {
	fmt.Printf("Type: %T, Value: %v\n", value, value)
}

// Interface với nhiều methods
type Animal interface {
	Speak() string
	Move() string
}

type Dog struct {
	Name string
}

func (d Dog) Speak() string {
	return "Woof!"
}

func (d Dog) Move() string {
	return "Running"
}

type Cat struct {
	Name string
}

func (c Cat) Speak() string {
	return "Meow!"
}

func (c Cat) Move() string {
	return "Jumping"
}

// Interface embedding
type Writer interface {
	Write([]byte) (int, error)
}

type Closer interface {
	Close() error
}

type WriteCloser interface {
	Writer
	Closer
}

// Type assertion và type switch
func CheckType(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Integer: %d\n", v)
	case string:
		fmt.Printf("String: %s\n", v)
	case bool:
		fmt.Printf("Boolean: %t\n", v)
	default:
		fmt.Printf("Unknown type: %T\n", v)
	}
}

// ========== CUSTOM TYPES ==========

type MyString string

// Method cho custom type
func (ms MyString) Upper() string {
	return strings.ToUpper(string(ms))
}

func (ms MyString) Lower() string {
	return strings.ToLower(string(ms))
}

// ========== STRINGER INTERFACE ==========

type Person struct {
	Name string
	Age  int
}

// Implement fmt.Stringer interface
func (p Person) String() string {
	return fmt.Sprintf("%s (%d years old)", p.Name, p.Age)
}

// ========== PRACTICAL EXAMPLES ==========

// Payment processor example
type PaymentMethod interface {
	Pay(amount float64) string
	GetName() string
}

type CreditCard struct {
	CardNumber string
	Holder     string
}

func (cc CreditCard) Pay(amount float64) string {
	return fmt.Sprintf("Paid $%.2f with Credit Card ending in %s",
		amount, cc.CardNumber[len(cc.CardNumber)-4:])
}

func (cc CreditCard) GetName() string {
	return "Credit Card"
}

type PayPal struct {
	Email string
}

func (pp PayPal) Pay(amount float64) string {
	return fmt.Sprintf("Paid $%.2f via PayPal account %s", amount, pp.Email)
}

func (pp PayPal) GetName() string {
	return "PayPal"
}

func ProcessPayment(pm PaymentMethod, amount float64) {
	fmt.Printf("Payment Method: %s\n", pm.GetName())
	fmt.Println(pm.Pay(amount))
}

func main() {
	// 1. BASIC METHODS
	fmt.Println("=== 1. BASIC METHODS ===")

	rect := Rectangle{Width: 10, Height: 5}
	fmt.Printf("Rectangle: %+v\n", rect)
	fmt.Printf("Area: %.2f\n", rect.Area())
	fmt.Printf("Perimeter: %.2f\n", rect.Perimeter())

	// 2. POINTER RECEIVER METHODS
	fmt.Println("\n=== 2. POINTER RECEIVER METHODS ===")

	fmt.Printf("Before scale: %+v\n", rect)
	rect.Scale(2)
	fmt.Printf("After scale(2): %+v\n", rect)
	fmt.Printf("New area: %.2f\n", rect.Area())

	// 3. INTERFACES
	fmt.Println("\n=== 3. INTERFACES ===")

	circle := Circle{Radius: 5}
	rectangle := Rectangle{Width: 10, Height: 5}

	// Cả hai đều implement Shape interface
	fmt.Println("Circle info:")
	PrintShapeInfo(circle)

	fmt.Println("\nRectangle info:")
	PrintShapeInfo(rectangle)

	// Slice of interfaces
	shapes := []Shape{
		Circle{Radius: 3},
		Rectangle{Width: 4, Height: 6},
		Circle{Radius: 7},
	}

	fmt.Println("\nAll shapes:")
	totalArea := 0.0
	for i, shape := range shapes {
		fmt.Printf("Shape %d: ", i+1)
		PrintShapeInfo(shape)
		totalArea += shape.Area()
	}
	fmt.Printf("Total area: %.2f\n", totalArea)

	// 4. EMPTY INTERFACE
	fmt.Println("\n=== 4. EMPTY INTERFACE ===")

	PrintAnything(42)
	PrintAnything("Hello")
	PrintAnything(true)
	PrintAnything(3.14)
	PrintAnything([]int{1, 2, 3})

	// 5. TYPE ASSERTION
	fmt.Println("\n=== 5. TYPE ASSERTION ===")

	var i interface{} = "hello"

	// Type assertion
	s, ok := i.(string)
	if ok {
		fmt.Printf("String value: %s\n", s)
	}

	// This would panic without ok check
	// n := i.(int) // panic

	n, ok := i.(int)
	if !ok {
		fmt.Println("Not an integer")
	} else {
		fmt.Printf("Integer value: %d\n", n)
	}

	// 6. TYPE SWITCH
	fmt.Println("\n=== 6. TYPE SWITCH ===")

	CheckType(42)
	CheckType("hello")
	CheckType(true)
	CheckType(3.14)

	// 7. ANIMAL INTERFACE
	fmt.Println("\n=== 7. ANIMAL INTERFACE ===")

	animals := []Animal{
		Dog{Name: "Buddy"},
		Cat{Name: "Whiskers"},
		Dog{Name: "Max"},
	}

	for _, animal := range animals {
		fmt.Printf("%s says: %s and is %s\n",
			animal, animal.Speak(), animal.Move())
	}

	// 8. CUSTOM TYPE METHODS
	fmt.Println("\n=== 8. CUSTOM TYPE METHODS ===")

	text := MyString("Hello World")
	fmt.Printf("Original: %s\n", text)
	fmt.Printf("Upper: %s\n", text.Upper())
	fmt.Printf("Lower: %s\n", text.Lower())

	// 9. STRINGER INTERFACE
	fmt.Println("\n=== 9. STRINGER INTERFACE ===")

	person := Person{Name: "Alice", Age: 25}
	// Tự động gọi String() method
	fmt.Println(person)
	fmt.Printf("Person: %v\n", person)

	// 10. PAYMENT PROCESSING EXAMPLE
	fmt.Println("\n=== 10. PAYMENT PROCESSING ===")

	card := CreditCard{
		CardNumber: "1234-5678-9012-3456",
		Holder:     "Alice",
	}

	paypal := PayPal{
		Email: "alice@example.com",
	}

	ProcessPayment(card, 99.99)
	fmt.Println()
	ProcessPayment(paypal, 49.99)

	// 11. INTERFACE VALUES
	fmt.Println("\n=== 11. INTERFACE VALUES ===")

	var shape Shape
	fmt.Printf("Nil interface: %v (nil=%t)\n", shape, shape == nil)

	shape = Circle{Radius: 5}
	fmt.Printf("After assignment: %v (nil=%t)\n", shape, shape == nil)
	fmt.Printf("Type: %T, Value: %+v\n", shape, shape)

	// 12. INTERFACE COMPOSITION
	fmt.Println("\n=== 12. INTERFACE COMPOSITION ===")

	// Smaller interfaces are better
	type Reader interface {
		Read([]byte) (int, error)
	}

	type Writer interface {
		Write([]byte) (int, error)
	}

	// Composed interface
	type ReadWriter interface {
		Reader
		Writer
	}

	fmt.Println("Interface composition allows building complex interfaces from simple ones")

	// 13. POLYMORPHISM
	fmt.Println("\n=== 13. POLYMORPHISM ===")

	// Database example
	type Database interface {
		Connect() string
		Query(sql string) string
	}

	type MySQL struct {
		Host string
	}

	func (m MySQL) Connect() string {
		return "Connected to MySQL at " + m.Host
	}

	func (m MySQL) Query(sql string) string {
		return "MySQL executing: " + sql
	}

	type PostgreSQL struct {
		Host string
	}

	func (p PostgreSQL) Connect() string {
		return "Connected to PostgreSQL at " + p.Host
	}

	func (p PostgreSQL) Query(sql string) string {
		return "PostgreSQL executing: " + sql
	}

	executeQuery := func(db Database, query string) {
		fmt.Println(db.Connect())
		fmt.Println(db.Query(query))
		fmt.Println()
	}

	mysql := MySQL{Host: "localhost:3306"}
	postgres := PostgreSQL{Host: "localhost:5432"}

	executeQuery(mysql, "SELECT * FROM users")
	executeQuery(postgres, "SELECT * FROM users")
}
