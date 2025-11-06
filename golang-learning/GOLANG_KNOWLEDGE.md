# Kiến Thức Golang - Từ Cơ Bản Đến Nâng Cao

## 📚 Tổng Quan

Tài liệu này tổng hợp toàn bộ kiến thức Golang từ cơ bản đến nâng cao, phù hợp cho Backend Developer.

---

# 🎯 PHẦN 1: BASICS - CƠ BẢN

## 1. Variables & Data Types

### Khai báo biến
```go
// Cách 1: var với type
var name string = "Go"
var age int = 10

// Cách 2: var với type inference
var language = "Golang"

// Cách 3: Short declaration (chỉ trong function)
city := "Hanoi"

// Multiple variables
var (
    firstName string = "Nguyen"
    lastName  string = "Van A"
    year      int    = 2025
)
```

### Basic Types
```go
// Boolean
var isActive bool = true

// Numeric types
var i8 int8       // -128 to 127
var i16 int16     // -32768 to 32767
var i32 int32     // -2^31 to 2^31-1
var i64 int64     // -2^63 to 2^63-1
var ui8 uint8     // 0 to 255
var ui16 uint16   // 0 to 65535
var ui32 uint32   // 0 to 2^32-1
var ui64 uint64   // 0 to 2^64-1

// Float
var f32 float32 = 3.14
var f64 float64 = 3.14159265359

// String
var text string = "Hello, 世界"

// Rune (Unicode code point)
var r rune = '世' // rune is alias for int32

// Byte
var b byte = 'A' // byte is alias for uint8
```

### Constants
```go
const Pi = 3.14
const AppName = "MyApp"

// Iota - auto increment
const (
    Sunday = iota    // 0
    Monday           // 1
    Tuesday          // 2
    Wednesday        // 3
)
```

### Type Conversion
```go
var i int = 42
var f float64 = float64(i)
var u uint = uint(f)

// String conversion
s := strconv.Itoa(42)        // int to string
n, _ := strconv.Atoi("42")   // string to int
```

## 2. Functions

### Basic Function
```go
func add(a int, b int) int {
    return a + b
}

// Same type parameters
func multiply(a, b, c int) int {
    return a * b * c
}
```

### Multiple Return Values
```go
func divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, errors.New("division by zero")
    }
    return a / b, nil
}

// Named return values
func swap(a, b string) (first, second string) {
    first = b
    second = a
    return // naked return
}
```

### Variadic Functions
```go
func sum(numbers ...int) int {
    total := 0
    for _, num := range numbers {
        total += num
    }
    return total
}

// Usage
result := sum(1, 2, 3, 4, 5)
```

### Anonymous Functions & Closures
```go
// Anonymous function
add := func(a, b int) int {
    return a + b
}

// Closure
func counter() func() int {
    count := 0
    return func() int {
        count++
        return count
    }
}

c := counter()
fmt.Println(c()) // 1
fmt.Println(c()) // 2
```

### Defer
```go
func readFile(filename string) {
    file, _ := os.Open(filename)
    defer file.Close() // Thực thi cuối cùng
    
    // Read file...
}
```

## 3. Control Flow

### If-Else
```go
if age >= 18 {
    fmt.Println("Adult")
} else if age >= 13 {
    fmt.Println("Teenager")
} else {
    fmt.Println("Child")
}

// If with short statement
if err := doSomething(); err != nil {
    return err
}
```

### Switch
```go
// Basic switch
switch day {
case "Monday":
    fmt.Println("Start of week")
case "Friday":
    fmt.Println("TGIF!")
default:
    fmt.Println("Regular day")
}

// Switch without condition (like if-else)
switch {
case score >= 90:
    grade = "A"
case score >= 80:
    grade = "B"
default:
    grade = "C"
}

// Type switch
switch v := i.(type) {
case int:
    fmt.Printf("Integer: %d", v)
case string:
    fmt.Printf("String: %s", v)
}
```

### For Loop
```go
// Traditional for
for i := 0; i < 10; i++ {
    fmt.Println(i)
}

// While-like
i := 0
for i < 10 {
    i++
}

// Infinite loop
for {
    // break to exit
}

// Range over slice/array
nums := []int{1, 2, 3, 4, 5}
for index, value := range nums {
    fmt.Printf("%d: %d\n", index, value)
}

// Range over map
m := map[string]int{"a": 1, "b": 2}
for key, value := range m {
    fmt.Printf("%s: %d\n", key, value)
}
```

## 4. Arrays & Slices

### Arrays (Fixed size)
```go
var arr [5]int
arr[0] = 1

// Initialize
arr := [5]int{1, 2, 3, 4, 5}

// Auto size
arr := [...]int{1, 2, 3}

// Length
len(arr) // 3
```

### Slices (Dynamic size)
```go
// Create slice
var s []int
s = []int{1, 2, 3}

// Make slice
s := make([]int, 5)      // length 5, capacity 5
s := make([]int, 5, 10)  // length 5, capacity 10

// Append
s = append(s, 6, 7, 8)

// Slicing
arr := []int{0, 1, 2, 3, 4, 5}
s := arr[1:4]  // [1, 2, 3]
s := arr[:3]   // [0, 1, 2]
s := arr[3:]   // [3, 4, 5]

// Copy
dst := make([]int, len(src))
copy(dst, src)

// Remove element
s = append(s[:i], s[i+1:]...)
```

### Slice operations
```go
// Length and capacity
len(s)
cap(s)

// 2D slice
matrix := [][]int{
    {1, 2, 3},
    {4, 5, 6},
}
```

## 5. Maps

### Create Map
```go
// Make map
m := make(map[string]int)

// Initialize
m := map[string]int{
    "apple":  5,
    "banana": 3,
}
```

### Map Operations
```go
// Set
m["orange"] = 7

// Get
value := m["apple"]

// Get with check
value, exists := m["grape"]
if exists {
    fmt.Println(value)
}

// Delete
delete(m, "banana")

// Length
len(m)

// Iterate
for key, value := range m {
    fmt.Printf("%s: %d\n", key, value)
}
```

## 6. Structs

### Define Struct
```go
type Person struct {
    Name    string
    Age     int
    Email   string
}
```

### Create & Initialize
```go
// Method 1
var p Person
p.Name = "John"
p.Age = 30

// Method 2
p := Person{
    Name: "John",
    Age:  30,
}

// Method 3
p := Person{"John", 30, "john@example.com"}
```

### Embedded Structs
```go
type Address struct {
    City    string
    Country string
}

type Person struct {
    Name    string
    Address // Embedded
}

p := Person{
    Name: "John",
    Address: Address{
        City:    "Hanoi",
        Country: "Vietnam",
    },
}

// Access
fmt.Println(p.City) // Direct access to embedded field
```

### Struct Tags
```go
type User struct {
    ID       int    `json:"id"`
    Username string `json:"username"`
    Email    string `json:"email,omitempty"`
}
```

## 7. Methods & Interfaces

### Methods
```go
type Rectangle struct {
    Width  float64
    Height float64
}

// Value receiver
func (r Rectangle) Area() float64 {
    return r.Width * r.Height
}

// Pointer receiver
func (r *Rectangle) Scale(factor float64) {
    r.Width *= factor
    r.Height *= factor
}
```

### Interfaces
```go
type Shape interface {
    Area() float64
    Perimeter() float64
}

type Circle struct {
    Radius float64
}

func (c Circle) Area() float64 {
    return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
    return 2 * math.Pi * c.Radius
}

// Circle implements Shape interface
var s Shape = Circle{Radius: 5}
```

### Empty Interface
```go
var i interface{} // Can hold any type

i = 42
i = "hello"
i = []int{1, 2, 3}

// Type assertion
value, ok := i.(string)
if ok {
    fmt.Println("String:", value)
}
```

## 8. Pointers

### Basics
```go
x := 42
p := &x     // p là pointer to x
fmt.Println(*p) // Dereference: 42

*p = 21     // Thay đổi giá trị qua pointer
fmt.Println(x)  // 21
```

### Pointers with Structs
```go
type Person struct {
    Name string
    Age  int
}

func (p *Person) Birthday() {
    p.Age++
}

p := &Person{Name: "John", Age: 30}
p.Birthday()
```

## 9. Error Handling

### Basic Error
```go
func divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, errors.New("division by zero")
    }
    return a / b, nil
}

result, err := divide(10, 0)
if err != nil {
    log.Fatal(err)
}
```

### Custom Error
```go
type ValidationError struct {
    Field   string
    Message string
}

func (e *ValidationError) Error() string {
    return fmt.Sprintf("%s: %s", e.Field, e.Message)
}

func validate(email string) error {
    if !strings.Contains(email, "@") {
        return &ValidationError{
            Field:   "email",
            Message: "invalid email format",
        }
    }
    return nil
}
```

### Error Wrapping (Go 1.13+)
```go
if err != nil {
    return fmt.Errorf("failed to read file: %w", err)
}

// Unwrap
errors.Unwrap(err)

// Check error type
if errors.Is(err, os.ErrNotExist) {
    // File doesn't exist
}

// Error as type
var pathError *os.PathError
if errors.As(err, &pathError) {
    fmt.Println(pathError.Path)
}
```

---

# ⚡ PHẦN 2: INTERMEDIATE - TRUNG CẤP

## 1. Goroutines

### Basic Goroutine
```go
func sayHello() {
    fmt.Println("Hello")
}

go sayHello() // Chạy concurrent
```

### Anonymous Goroutine
```go
go func() {
    fmt.Println("Anonymous goroutine")
}()
```

### Wait for Goroutines
```go
var wg sync.WaitGroup

for i := 0; i < 5; i++ {
    wg.Add(1)
    go func(n int) {
        defer wg.Done()
        fmt.Println(n)
    }(i)
}

wg.Wait() // Chờ tất cả goroutines
```

## 2. Channels

### Basic Channel
```go
ch := make(chan int)

// Send
go func() {
    ch <- 42
}()

// Receive
value := <-ch
```

### Buffered Channel
```go
ch := make(chan int, 3) // Buffer size 3

ch <- 1
ch <- 2
ch <- 3 // Không block

value := <-ch // 1
```

### Close Channel
```go
ch := make(chan int)

go func() {
    for i := 0; i < 5; i++ {
        ch <- i
    }
    close(ch)
}()

// Range over channel
for value := range ch {
    fmt.Println(value)
}
```

### Channel Directions
```go
// Send-only
func send(ch chan<- int) {
    ch <- 42
}

// Receive-only
func receive(ch <-chan int) {
    value := <-ch
}
```

## 3. Select

### Basic Select
```go
ch1 := make(chan string)
ch2 := make(chan string)

select {
case msg1 := <-ch1:
    fmt.Println("Received:", msg1)
case msg2 := <-ch2:
    fmt.Println("Received:", msg2)
}
```

### Select with Timeout
```go
select {
case msg := <-ch:
    fmt.Println(msg)
case <-time.After(1 * time.Second):
    fmt.Println("Timeout!")
}
```

### Non-blocking Select
```go
select {
case msg := <-ch:
    fmt.Println(msg)
default:
    fmt.Println("No message")
}
```

## 4. Sync Package

### WaitGroup
```go
var wg sync.WaitGroup

for i := 0; i < 10; i++ {
    wg.Add(1)
    go func(n int) {
        defer wg.Done()
        fmt.Println(n)
    }(i)
}

wg.Wait()
```

### Mutex
```go
var (
    counter int
    mu      sync.Mutex
)

func increment() {
    mu.Lock()
    counter++
    mu.Unlock()
}
```

### RWMutex
```go
var (
    data map[string]string
    mu   sync.RWMutex
)

// Read
func read(key string) string {
    mu.RLock()
    defer mu.RUnlock()
    return data[key]
}

// Write
func write(key, value string) {
    mu.Lock()
    defer mu.Unlock()
    data[key] = value
}
```

### Once
```go
var once sync.Once

func initialize() {
    once.Do(func() {
        fmt.Println("Chỉ chạy 1 lần")
    })
}
```

## 5. Patterns

### Worker Pool
```go
func worker(id int, jobs <-chan int, results chan<- int) {
    for j := range jobs {
        fmt.Printf("Worker %d processing job %d\n", id, j)
        results <- j * 2
    }
}

func main() {
    jobs := make(chan int, 100)
    results := make(chan int, 100)

    // Start workers
    for w := 1; w <= 3; w++ {
        go worker(w, jobs, results)
    }

    // Send jobs
    for j := 1; j <= 9; j++ {
        jobs <- j
    }
    close(jobs)

    // Collect results
    for a := 1; a <= 9; a++ {
        <-results
    }
}
```

### Pipeline
```go
func generate(nums ...int) <-chan int {
    out := make(chan int)
    go func() {
        for _, n := range nums {
            out <- n
        }
        close(out)
    }()
    return out
}

func square(in <-chan int) <-chan int {
    out := make(chan int)
    go func() {
        for n := range in {
            out <- n * n
        }
        close(out)
    }()
    return out
}

// Usage
nums := generate(1, 2, 3, 4)
squared := square(nums)
for n := range squared {
    fmt.Println(n)
}
```

### Fan-Out Fan-In
```go
func fanOut(in <-chan int, n int) []<-chan int {
    channels := make([]<-chan int, n)
    for i := 0; i < n; i++ {
        channels[i] = worker(in)
    }
    return channels
}

func fanIn(channels ...<-chan int) <-chan int {
    out := make(chan int)
    var wg sync.WaitGroup
    
    for _, ch := range channels {
        wg.Add(1)
        go func(c <-chan int) {
            defer wg.Done()
            for n := range c {
                out <- n
            }
        }(ch)
    }
    
    go func() {
        wg.Wait()
        close(out)
    }()
    
    return out
}
```

## 6. Context

### WithCancel
```go
ctx, cancel := context.WithCancel(context.Background())
defer cancel()

go func() {
    for {
        select {
        case <-ctx.Done():
            return
        default:
            // Do work
        }
    }
}()

time.Sleep(1 * time.Second)
cancel() // Cancel goroutine
```

### WithTimeout
```go
ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
defer cancel()

select {
case <-time.After(3 * time.Second):
    fmt.Println("Work done")
case <-ctx.Done():
    fmt.Println("Timeout!")
}
```

### WithValue
```go
ctx := context.WithValue(context.Background(), "userID", 123)

func processRequest(ctx context.Context) {
    userID := ctx.Value("userID").(int)
    fmt.Println("User ID:", userID)
}
```

---

# 🚀 PHẦN 3: ADVANCED - NÂNG CAO

## 1. Reflection

### Basic Reflection
```go
import "reflect"

var x float64 = 3.14
t := reflect.TypeOf(x)
v := reflect.ValueOf(x)

fmt.Println("Type:", t)
fmt.Println("Value:", v)
fmt.Println("Kind:", v.Kind())
```

### Struct Reflection
```go
type Person struct {
    Name string `json:"name"`
    Age  int    `json:"age"`
}

p := Person{Name: "John", Age: 30}
v := reflect.ValueOf(p)

for i := 0; i < v.NumField(); i++ {
    field := v.Type().Field(i)
    value := v.Field(i)
    
    fmt.Printf("%s: %v (tag: %s)\n",
        field.Name,
        value.Interface(),
        field.Tag.Get("json"))
}
```

### Modify Values
```go
x := 42
v := reflect.ValueOf(&x)
v.Elem().SetInt(100)
fmt.Println(x) // 100
```

## 2. Generics (Go 1.18+)

### Generic Function
```go
func Min[T constraints.Ordered](a, b T) T {
    if a < b {
        return a
    }
    return b
}

// Usage
fmt.Println(Min(3, 5))         // 3
fmt.Println(Min(3.14, 2.71))   // 2.71
fmt.Println(Min("abc", "xyz")) // "abc"
```

### Generic Type
```go
type Stack[T any] struct {
    items []T
}

func (s *Stack[T]) Push(item T) {
    s.items = append(s.items, item)
}

func (s *Stack[T]) Pop() T {
    item := s.items[len(s.items)-1]
    s.items = s.items[:len(s.items)-1]
    return item
}

// Usage
intStack := Stack[int]{}
intStack.Push(1)
intStack.Push(2)
```

### Type Constraints
```go
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
```

## 3. Design Patterns

### Singleton
```go
type singleton struct{}

var instance *singleton
var once sync.Once

func GetInstance() *singleton {
    once.Do(func() {
        instance = &singleton{}
    })
    return instance
}
```

### Factory
```go
type Animal interface {
    Speak() string
}

type Dog struct{}
func (d Dog) Speak() string { return "Woof!" }

type Cat struct{}
func (c Cat) Speak() string { return "Meow!" }

func NewAnimal(animalType string) Animal {
    switch animalType {
    case "dog":
        return Dog{}
    case "cat":
        return Cat{}
    default:
        return nil
    }
}
```

### Builder
```go
type House struct {
    windows int
    doors   int
    floors  int
}

type HouseBuilder struct {
    house House
}

func NewHouseBuilder() *HouseBuilder {
    return &HouseBuilder{}
}

func (b *HouseBuilder) Windows(n int) *HouseBuilder {
    b.house.windows = n
    return b
}

func (b *HouseBuilder) Doors(n int) *HouseBuilder {
    b.house.doors = n
    return b
}

func (b *HouseBuilder) Build() House {
    return b.house
}

// Usage
house := NewHouseBuilder().
    Windows(4).
    Doors(2).
    Build()
```

### Observer
```go
type Observer interface {
    Update(data string)
}

type Subject struct {
    observers []Observer
}

func (s *Subject) Attach(o Observer) {
    s.observers = append(s.observers, o)
}

func (s *Subject) Notify(data string) {
    for _, o := range s.observers {
        o.Update(data)
    }
}
```

## 4. Performance Optimization

### Benchmarking
```go
func BenchmarkSum(b *testing.B) {
    for i := 0; i < b.N; i++ {
        sum(1, 2, 3, 4, 5)
    }
}

// Run: go test -bench=.
```

### Profiling
```go
import _ "net/http/pprof"

func main() {
    go func() {
        log.Println(http.ListenAndServe("localhost:6060", nil))
    }()
    
    // Application code...
}

// Visit: http://localhost:6060/debug/pprof/
```

### Memory Optimization
```go
// Bad: Creates many allocations
func badConcat(strs []string) string {
    result := ""
    for _, s := range strs {
        result += s
    }
    return result
}

// Good: Uses strings.Builder
func goodConcat(strs []string) string {
    var builder strings.Builder
    for _, s := range strs {
        builder.WriteString(s)
    }
    return builder.String()
}
```

### Sync.Pool
```go
var bufferPool = sync.Pool{
    New: func() interface{} {
        return new(bytes.Buffer)
    },
}

func processData(data []byte) {
    buf := bufferPool.Get().(*bytes.Buffer)
    defer bufferPool.Put(buf)
    
    buf.Reset()
    buf.Write(data)
    // Process...
}
```

## 5. Advanced Concurrency

### Rate Limiting
```go
type RateLimiter struct {
    ticker *time.Ticker
}

func NewRateLimiter(rps int) *RateLimiter {
    return &RateLimiter{
        ticker: time.NewTicker(time.Second / time.Duration(rps)),
    }
}

func (rl *RateLimiter) Wait() {
    <-rl.ticker.C
}
```

### Circuit Breaker
```go
type CircuitBreaker struct {
    maxFailures int
    failures    int
    state       string // "closed", "open", "half-open"
    mu          sync.Mutex
}

func (cb *CircuitBreaker) Call(fn func() error) error {
    cb.mu.Lock()
    defer cb.mu.Unlock()
    
    if cb.state == "open" {
        return errors.New("circuit breaker is open")
    }
    
    err := fn()
    if err != nil {
        cb.failures++
        if cb.failures >= cb.maxFailures {
            cb.state = "open"
        }
        return err
    }
    
    cb.failures = 0
    cb.state = "closed"
    return nil
}
```

### Semaphore
```go
type Semaphore struct {
    ch chan struct{}
}

func NewSemaphore(n int) *Semaphore {
    return &Semaphore{
        ch: make(chan struct{}, n),
    }
}

func (s *Semaphore) Acquire() {
    s.ch <- struct{}{}
}

func (s *Semaphore) Release() {
    <-s.ch
}
```

---

## 📚 Best Practices

### 1. Error Handling
- Luôn check errors
- Wrap errors với context
- Tạo custom errors khi cần
- Sử dụng errors.Is và errors.As

### 2. Concurrency
- Tránh goroutine leaks
- Luôn close channels
- Sử dụng context cho cancellation
- Test với -race flag

### 3. Code Organization
- Một package một mục đích
- Export chỉ những gì cần thiết
- Tránh circular dependencies
- Viết godoc comments

### 4. Performance
- Profile trước khi optimize
- Tránh premature optimization
- Sử dụng benchmark
- Reuse objects với sync.Pool

### 5. Testing
- Viết unit tests
- Sử dụng table-driven tests
- Test edge cases
- Mock external dependencies

---

## 🎯 Tips & Tricks

1. **Sử dụng gofmt**: Format code tự động
2. **Chạy go vet**: Tìm bugs tiềm ẩn
3. **Sử dụng golangci-lint**: Linter tổng hợp
4. **Đọc Effective Go**: Best practices chính thức
5. **Join Go community**: Học từ người khác

---

## 📖 Resources

- [Official Go Documentation](https://go.dev/doc/)
- [Go by Example](https://gobyexample.com/)
- [Effective Go](https://go.dev/doc/effective_go)
- [Go Blog](https://go.dev/blog/)
- [Awesome Go](https://github.com/avelino/awesome-go)

---

**Happy Coding! 🚀**
