package main

import (
	"fmt"
	"reflect"
	"sync"
	"time"
)

// ============================================
// 01. REFLECTION EXAMPLES
// ============================================

func reflectionExamples() {
	fmt.Println("=== REFLECTION EXAMPLES ===")

	// Basic reflection
	var x float64 = 3.14
	v := reflect.ValueOf(x)
	fmt.Println("Type:", v.Type())
	fmt.Println("Kind:", v.Kind())
	fmt.Println("Value:", v.Float())

	// Struct reflection
	type Person struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}

	p := Person{Name: "John", Age: 30}
	t := reflect.TypeOf(p)
	v = reflect.ValueOf(p)

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		value := v.Field(i)
		fmt.Printf("Field: %s, Type: %s, Value: %v, Tag: %s\n",
			field.Name, field.Type, value.Interface(), field.Tag.Get("json"))
	}

	// Modify via reflection
	p2 := Person{Name: "Jane", Age: 25}
	v = reflect.ValueOf(&p2).Elem()
	v.FieldByName("Age").SetInt(26)
	fmt.Println("Modified:", p2)
}

// ============================================
// 02. GENERICS EXAMPLES
// ============================================

// Generic function
func Print[T any](items []T) {
	for _, item := range items {
		fmt.Println(item)
	}
}

// Generic type
type Pair[T, U any] struct {
	First  T
	Second U
}

// Generic constraint
type Number interface {
	int | int64 | float64
}

func Sum[T Number](numbers []T) T {
	var sum T
	for _, n := range numbers {
		sum += n
	}
	return sum
}

func genericsExamples() {
	fmt.Println("\n=== GENERICS EXAMPLES ===")

	// Generic function
	Print([]int{1, 2, 3})
	Print([]string{"a", "b", "c"})

	// Generic type
	p1 := Pair[string, int]{First: "age", Second: 30}
	fmt.Println(p1)

	// Generic constraint
	intSum := Sum([]int{1, 2, 3, 4, 5})
	floatSum := Sum([]float64{1.1, 2.2, 3.3})
	fmt.Println("Int sum:", intSum)
	fmt.Println("Float sum:", floatSum)
}

// ============================================
// 03. DESIGN PATTERNS
// ============================================

// Singleton Pattern
type singleton struct {
	data string
}

var (
	instance *singleton
	once     sync.Once
)

func GetInstance() *singleton {
	once.Do(func() {
		instance = &singleton{data: "singleton data"}
	})
	return instance
}

// Factory Pattern
type Shape interface {
	Draw() string
}

type Circle struct {
	Radius float64
}

func (c Circle) Draw() string {
	return fmt.Sprintf("Circle with radius %.2f", c.Radius)
}

type Rectangle struct {
	Width, Height float64
}

func (r Rectangle) Draw() string {
	return fmt.Sprintf("Rectangle %.2fx%.2f", r.Width, r.Height)
}

func CreateShape(shapeType string) Shape {
	switch shapeType {
	case "circle":
		return Circle{Radius: 5}
	case "rectangle":
		return Rectangle{Width: 10, Height: 5}
	default:
		return nil
	}
}

// Builder Pattern
type Car struct {
	Brand  string
	Model  string
	Color  string
	Engine string
}

type CarBuilder struct {
	car Car
}

func NewCarBuilder() *CarBuilder {
	return &CarBuilder{}
}

func (b *CarBuilder) Brand(brand string) *CarBuilder {
	b.car.Brand = brand
	return b
}

func (b *CarBuilder) Model(model string) *CarBuilder {
	b.car.Model = model
	return b
}

func (b *CarBuilder) Color(color string) *CarBuilder {
	b.car.Color = color
	return b
}

func (b *CarBuilder) Engine(engine string) *CarBuilder {
	b.car.Engine = engine
	return b
}

func (b *CarBuilder) Build() Car {
	return b.car
}

// Observer Pattern
type EventObserver interface {
	OnEvent(event string)
}

type EventPublisher struct {
	observers []EventObserver
}

func (ep *EventPublisher) Subscribe(observer EventObserver) {
	ep.observers = append(ep.observers, observer)
}

func (ep *EventPublisher) Publish(event string) {
	for _, observer := range ep.observers {
		observer.OnEvent(event)
	}
}

type EmailNotifier struct {
	email string
}

func (en EmailNotifier) OnEvent(event string) {
	fmt.Printf("Email to %s: %s\n", en.email, event)
}

type SMSNotifier struct {
	phone string
}

func (sn SMSNotifier) OnEvent(event string) {
	fmt.Printf("SMS to %s: %s\n", sn.phone, event)
}

func designPatternsExamples() {
	fmt.Println("\n=== DESIGN PATTERNS EXAMPLES ===")

	// Singleton
	s1 := GetInstance()
	s2 := GetInstance()
	fmt.Println("Same instance?", s1 == s2)

	// Factory
	circle := CreateShape("circle")
	rectangle := CreateShape("rectangle")
	fmt.Println(circle.Draw())
	fmt.Println(rectangle.Draw())

	// Builder
	car := NewCarBuilder().
		Brand("Toyota").
		Model("Camry").
		Color("Red").
		Engine("V6").
		Build()
	fmt.Printf("Car: %+v\n", car)

	// Observer
	publisher := &EventPublisher{}
	publisher.Subscribe(EmailNotifier{email: "user@example.com"})
	publisher.Subscribe(SMSNotifier{phone: "+84123456789"})
	publisher.Publish("New message received!")
}

// ============================================
// 04. PERFORMANCE EXAMPLES
// ============================================

// Bad: String concatenation
func badStringConcat(strs []string) string {
	result := ""
	for _, s := range strs {
		result += s // Creates new string each time
	}
	return result
}

// Good: Using strings.Builder
func goodStringConcat(strs []string) string {
	var builder = &struct {
		buf []byte
	}{}
	for _, s := range strs {
		builder.buf = append(builder.buf, s...)
	}
	return string(builder.buf)
}

// Object pooling with sync.Pool
type Buffer struct {
	data []byte
}

var bufferPool = sync.Pool{
	New: func() interface{} {
		return &Buffer{data: make([]byte, 0, 1024)}
	},
}

func processWithPool(input []byte) []byte {
	buf := bufferPool.Get().(*Buffer)
	defer bufferPool.Put(buf)

	buf.data = buf.data[:0] // Reset
	buf.data = append(buf.data, input...)
	// Process...
	return buf.data
}

func performanceExamples() {
	fmt.Println("\n=== PERFORMANCE EXAMPLES ===")

	strs := []string{"Hello", " ", "World", "!", " ", "Golang"}

	// Time bad version
	start := time.Now()
	_ = badStringConcat(strs)
	fmt.Println("Bad concat:", time.Since(start))

	// Time good version
	start = time.Now()
	_ = goodStringConcat(strs)
	fmt.Println("Good concat:", time.Since(start))

	// Object pool
	input := []byte("test data")
	result := processWithPool(input)
	fmt.Println("Pool result:", string(result))
}

// ============================================
// 05. ADVANCED CONCURRENCY
// ============================================

// Worker Pool Pattern
func workerPoolExample() {
	fmt.Println("\n=== WORKER POOL EXAMPLE ===")

	jobs := make(chan int, 100)
	results := make(chan int, 100)

	// Start workers
	for w := 1; w <= 3; w++ {
		go func(id int) {
			for j := range jobs {
				fmt.Printf("Worker %d processing job %d\n", id, j)
				time.Sleep(100 * time.Millisecond)
				results <- j * 2
			}
		}(w)
	}

	// Send jobs
	for j := 1; j <= 5; j++ {
		jobs <- j
	}
	close(jobs)

	// Collect results
	for a := 1; a <= 5; a++ {
		result := <-results
		fmt.Println("Result:", result)
	}
}

// Pipeline Pattern
func pipelineExample() {
	fmt.Println("\n=== PIPELINE EXAMPLE ===")

	// Stage 1: Generate numbers
	generate := func(nums ...int) <-chan int {
		out := make(chan int)
		go func() {
			for _, n := range nums {
				out <- n
			}
			close(out)
		}()
		return out
	}

	// Stage 2: Square numbers
	square := func(in <-chan int) <-chan int {
		out := make(chan int)
		go func() {
			for n := range in {
				out <- n * n
			}
			close(out)
		}()
		return out
	}

	// Stage 3: Filter even
	filterEven := func(in <-chan int) <-chan int {
		out := make(chan int)
		go func() {
			for n := range in {
				if n%2 == 0 {
					out <- n
				}
			}
			close(out)
		}()
		return out
	}

	// Build pipeline
	nums := generate(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	squared := square(nums)
	evens := filterEven(squared)

	// Consume
	for n := range evens {
		fmt.Println(n)
	}
}

// Rate Limiter Example (Simple)
type SimpleRateLimiter struct {
	ticker *time.Ticker
}

func NewSimpleRateLimiter(requestsPerSecond int) *SimpleRateLimiter {
	interval := time.Second / time.Duration(requestsPerSecond)
	return &SimpleRateLimiter{
		ticker: time.NewTicker(interval),
	}
}

func (rl *SimpleRateLimiter) Wait() {
	<-rl.ticker.C
}

func rateLimiterExample() {
	fmt.Println("\n=== RATE LIMITER EXAMPLE ===")

	limiter := NewSimpleRateLimiter(2) // 2 requests per second

	for i := 1; i <= 6; i++ {
		limiter.Wait()
		fmt.Printf("Request %d at %v\n", i, time.Now().Format("15:04:05"))
	}
}

// ============================================
// MAIN
// ============================================

func main() {
	reflectionExamples()
	genericsExamples()
	designPatternsExamples()
	performanceExamples()
	workerPoolExample()
	pipelineExample()
	rateLimiterExample()

	fmt.Println("\n=== ALL EXAMPLES COMPLETED ===")
}
