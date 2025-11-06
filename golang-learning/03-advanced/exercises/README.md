# Bài Tập Thực Hành - Advanced

## 🎯 Mục tiêu

Áp dụng kiến thức nâng cao: Reflection, Generics, Design Patterns, Performance Optimization

---

## Bài 1: Reflection Basics

**Mục tiêu:** Hiểu và sử dụng reflection để inspect types

**Nhiệm vụ:**
1. Viết hàm `inspectStruct` nhận interface{} và in ra:
   - Type name
   - Number of fields
   - Field names, types, và values
   - Struct tags

2. Viết hàm `callMethod` nhận struct và method name, gọi method đó

3. Viết hàn `modifyField` thay đổi giá trị field qua reflection

**Test cases:**
```go
type User struct {
    Name  string `json:"name" validate:"required"`
    Email string `json:"email" validate:"email"`
    Age   int    `json:"age" validate:"min=0,max=120"`
}

user := User{Name: "John", Email: "john@example.com", Age: 30}
inspectStruct(user)
// Output:
// Type: User
// Fields: 3
// Field 0: Name (string) = "John", tag: json:"name" validate:"required"
// ...
```

---

## Bài 2: JSON Serializer với Reflection

**Mục tiêu:** Build custom JSON serializer sử dụng reflection

**Nhiệm vụ:**
1. Implement `Marshal(v interface{}) ([]byte, error)`:
   - Support struct, map, slice, basic types
   - Respect struct tags (json, omitempty)
   - Handle nested structs

2. Implement `Unmarshal(data []byte, v interface{}) error`:
   - Parse JSON string
   - Set values qua reflection
   - Type conversion

**Test cases:**
```go
type Person struct {
    Name    string   `json:"name"`
    Age     int      `json:"age"`
    Emails  []string `json:"emails,omitempty"`
    Address *Address `json:"address,omitempty"`
}

type Address struct {
    City    string `json:"city"`
    Country string `json:"country"`
}

p := Person{Name: "John", Age: 30}
data, _ := Marshal(p)
// {"name":"John","age":30}
```

---

## Bài 3: Generics - Data Structures

**Mục tiêu:** Implement generic data structures

**Nhiệm vụ:**

### 3.1 Generic Stack
```go
type Stack[T any] struct {
    // Implement
}

func (s *Stack[T]) Push(item T)
func (s *Stack[T]) Pop() (T, error)
func (s *Stack[T]) Peek() (T, error)
func (s *Stack[T]) IsEmpty() bool
func (s *Stack[T]) Size() int
```

### 3.2 Generic Queue
```go
type Queue[T any] struct {
    // Implement
}

func (q *Queue[T]) Enqueue(item T)
func (q *Queue[T]) Dequeue() (T, error)
func (q *Queue[T]) Front() (T, error)
func (q *Queue[T]) IsEmpty() bool
func (q *Queue[T]) Size() int
```

### 3.3 Generic LinkedList
```go
type LinkedList[T any] struct {
    // Implement
}

func (ll *LinkedList[T]) Append(value T)
func (ll *LinkedList[T]) Prepend(value T)
func (ll *LinkedList[T]) Remove(value T) bool
func (ll *LinkedList[T]) Find(value T) bool
func (ll *LinkedList[T]) ToSlice() []T
```

**Test cases:**
```go
// Stack
stack := Stack[int]{}
stack.Push(1)
stack.Push(2)
val, _ := stack.Pop() // 2

// Queue
queue := Queue[string]{}
queue.Enqueue("first")
queue.Enqueue("second")
val, _ := queue.Dequeue() // "first"
```

---

## Bài 4: Generic Functions

**Mục tiêu:** Viết generic utility functions

**Nhiệm vụ:**
1. `Map[T, U any](slice []T, fn func(T) U) []U`
2. `Filter[T any](slice []T, fn func(T) bool) []T`
3. `Reduce[T, U any](slice []T, initial U, fn func(U, T) U) U`
4. `Find[T any](slice []T, fn func(T) bool) (T, bool)`
5. `Contains[T comparable](slice []T, item T) bool`
6. `Reverse[T any](slice []T) []T`
7. `Min[T constraints.Ordered](items ...T) T`
8. `Max[T constraints.Ordered](items ...T) T`

**Test cases:**
```go
nums := []int{1, 2, 3, 4, 5}

// Map
squared := Map(nums, func(n int) int { return n * n })
// [1, 4, 9, 16, 25]

// Filter
evens := Filter(nums, func(n int) bool { return n%2 == 0 })
// [2, 4]

// Reduce
sum := Reduce(nums, 0, func(acc, n int) int { return acc + n })
// 15
```

---

## Bài 5: Generic Cache

**Mục tiêu:** Build type-safe cache với generics

**Nhiệm vụ:**
```go
type Cache[K comparable, V any] struct {
    // Implement
}

func NewCache[K comparable, V any]() *Cache[K, V]
func (c *Cache[K, V]) Set(key K, value V)
func (c *Cache[K, V]) Get(key K) (V, bool)
func (c *Cache[K, V]) Delete(key K)
func (c *Cache[K, V]) Clear()
func (c *Cache[K, V]) Size() int

// Advanced features
func (c *Cache[K, V]) SetWithTTL(key K, value V, ttl time.Duration)
func (c *Cache[K, V]) SetMaxSize(maxSize int) // LRU eviction
```

**Requirements:**
- Thread-safe
- TTL support
- LRU eviction khi full
- Metrics (hits, misses)

**Test cases:**
```go
cache := NewCache[string, int]()
cache.Set("one", 1)
cache.Set("two", 2)

val, ok := cache.Get("one") // val=1, ok=true
val, ok = cache.Get("three") // val=0, ok=false

cache.SetWithTTL("temp", 42, 100*time.Millisecond)
time.Sleep(200*time.Millisecond)
val, ok = cache.Get("temp") // val=0, ok=false (expired)
```

---

## Bài 6: Design Patterns - Creational

### 6.1 Singleton Pattern
```go
type Database struct {
    // Connection pool, etc.
}

// Implement thread-safe singleton
func GetDatabase() *Database
```

### 6.2 Factory Pattern
```go
type Logger interface {
    Log(message string)
}

type ConsoleLogger struct{}
type FileLogger struct{}

func NewLogger(loggerType string) Logger
```

### 6.3 Builder Pattern
```go
type HTTPRequest struct {
    URL     string
    Method  string
    Headers map[string]string
    Body    []byte
}

type RequestBuilder struct {
    // Implement
}

func NewRequestBuilder() *RequestBuilder
func (b *RequestBuilder) URL(url string) *RequestBuilder
func (b *RequestBuilder) Method(method string) *RequestBuilder
func (b *RequestBuilder) Header(key, value string) *RequestBuilder
func (b *RequestBuilder) Body(body []byte) *RequestBuilder
func (b *RequestBuilder) Build() *HTTPRequest
```

**Usage:**
```go
req := NewRequestBuilder().
    URL("https://api.example.com").
    Method("POST").
    Header("Content-Type", "application/json").
    Body([]byte(`{"name":"John"}`)).
    Build()
```

---

## Bài 7: Design Patterns - Structural

### 7.1 Adapter Pattern
```go
// Old interface
type LegacyPrinter interface {
    PrintOld(text string)
}

// New interface
type ModernPrinter interface {
    Print(text string)
}

// Implement adapter
type PrinterAdapter struct {
    legacyPrinter LegacyPrinter
}

func (a *PrinterAdapter) Print(text string)
```

### 7.2 Decorator Pattern
```go
type Coffee interface {
    Cost() float64
    Description() string
}

type SimpleCoffee struct{}

func (c SimpleCoffee) Cost() float64       { return 5.0 }
func (c SimpleCoffee) Description() string { return "Simple Coffee" }

// Decorators
type MilkDecorator struct {
    coffee Coffee
}

type SugarDecorator struct {
    coffee Coffee
}

// Implement decorators
```

**Usage:**
```go
coffee := SimpleCoffee{}
withMilk := MilkDecorator{coffee}
withMilkAndSugar := SugarDecorator{withMilk}

fmt.Println(withMilkAndSugar.Description()) // "Simple Coffee, Milk, Sugar"
fmt.Println(withMilkAndSugar.Cost())        // 7.5
```

---

## Bài 8: Design Patterns - Behavioral

### 8.1 Observer Pattern
```go
type Observer interface {
    Update(data interface{})
}

type Subject interface {
    Attach(observer Observer)
    Detach(observer Observer)
    Notify(data interface{})
}

// Implement concrete subject và observers
```

### 8.2 Strategy Pattern
```go
type PaymentStrategy interface {
    Pay(amount float64) error
}

type CreditCardPayment struct{}
type PayPalPayment struct{}
type CryptoPayment struct{}

type ShoppingCart struct {
    paymentStrategy PaymentStrategy
}

func (sc *ShoppingCart) SetPaymentStrategy(strategy PaymentStrategy)
func (sc *ShoppingCart) Checkout(amount float64) error
```

### 8.3 Command Pattern
```go
type Command interface {
    Execute() error
    Undo() error
}

type TextEditor struct {
    text string
}

type InsertCommand struct {
    editor *TextEditor
    text   string
}

type DeleteCommand struct {
    editor *TextEditor
    start  int
    end    int
}

// Implement Execute() và Undo() cho mỗi command
```

---

## Bài 9: Performance Optimization

**Mục tiêu:** Profile và optimize code

### 9.1 Benchmarking
Viết benchmarks cho:
```go
// String concatenation
func BenchmarkStringConcat(b *testing.B)
func BenchmarkStringBuilder(b *testing.B)
func BenchmarkStringJoin(b *testing.B)

// Slice operations
func BenchmarkSliceAppend(b *testing.B)
func BenchmarkSlicePrealloc(b *testing.B)

// Map operations
func BenchmarkMapAccess(b *testing.B)
func BenchmarkMapIteration(b *testing.B)
```

### 9.2 Memory Optimization
```go
// Optimize these functions
func inefficientFilter(numbers []int) []int {
    var result []int
    for _, n := range numbers {
        if n%2 == 0 {
            result = append(result, n)
        }
    }
    return result
}

func inefficientConcat(strs []string) string {
    result := ""
    for _, s := range strs {
        result += s
    }
    return result
}
```

**Task:** 
- Profile memory usage
- Optimize allocations
- Compare với benchmarks

### 9.3 Sync.Pool
```go
// Implement object pool cho frequent allocations
type Buffer struct {
    data []byte
}

var bufferPool = sync.Pool{
    New: func() interface{} {
        return &Buffer{data: make([]byte, 0, 1024)}
    },
}

func processData(input []byte) []byte {
    // Use buffer from pool
    // Process data
    // Return buffer to pool
}
```

---

## Bài 10: Advanced Concurrency Patterns

### 10.1 Rate Limiter
```go
type RateLimiter struct {
    // Implement token bucket algorithm
}

func NewRateLimiter(rate int, burst int) *RateLimiter
func (rl *RateLimiter) Allow() bool
func (rl *RateLimiter) Wait(ctx context.Context) error
```

**Test:**
```go
limiter := NewRateLimiter(5, 10) // 5 req/sec, burst 10

for i := 0; i < 20; i++ {
    if limiter.Allow() {
        fmt.Printf("Request %d: ALLOWED\n", i)
    } else {
        fmt.Printf("Request %d: DENIED\n", i)
    }
}
```

### 10.2 Circuit Breaker
```go
type CircuitBreaker struct {
    // States: Closed, Open, Half-Open
}

func NewCircuitBreaker(maxFailures int, timeout time.Duration) *CircuitBreaker
func (cb *CircuitBreaker) Call(fn func() error) error
func (cb *CircuitBreaker) State() string
```

**Test:**
```go
cb := NewCircuitBreaker(3, 5*time.Second)

for i := 0; i < 10; i++ {
    err := cb.Call(func() error {
        // Simulated API call
        return errors.New("service unavailable")
    })
    
    if err != nil {
        fmt.Printf("Call %d failed: %v\n", i, err)
    }
}
```

### 10.3 Worker Pool với Retry
```go
type Job struct {
    ID      int
    Payload interface{}
    Retry   int
}

type WorkerPool struct {
    workers   int
    jobs      chan Job
    results   chan Result
    maxRetry  int
}

func NewWorkerPool(workers int, maxRetry int) *WorkerPool
func (wp *WorkerPool) Submit(job Job)
func (wp *WorkerPool) Start()
func (wp *WorkerPool) Stop()
```

---

## 🎯 Challenges

### Challenge 1: ORM Framework

Build mini ORM framework với:
- Struct to SQL mapping
- CRUD operations
- Query builder
- Transactions
- Migrations

**Example usage:**
```go
type User struct {
    ID        int    `db:"id,primary"`
    Username  string `db:"username,unique"`
    Email     string `db:"email"`
    CreatedAt time.Time `db:"created_at"`
}

db := NewDB("sqlite3", "test.db")

// Create
user := User{Username: "john", Email: "john@example.com"}
db.Create(&user)

// Find
var users []User
db.Where("username = ?", "john").Find(&users)

// Update
db.Model(&user).Update("email", "newemail@example.com")

// Delete
db.Delete(&user)
```

### Challenge 2: Dependency Injection Container

```go
type Container struct {
    // Implement DI container
}

func NewContainer() *Container
func (c *Container) Register(key string, factory func() interface{})
func (c *Container) Singleton(key string, factory func() interface{})
func (c *Container) Resolve(key string) (interface{}, error)
func (c *Container) ResolveAs(key string, target interface{}) error
```

**Usage:**
```go
container := NewContainer()

// Register dependencies
container.Singleton("database", func() interface{} {
    return NewDatabase()
})

container.Register("userService", func() interface{} {
    db, _ := container.Resolve("database")
    return NewUserService(db.(*Database))
})

// Resolve
userService, _ := container.Resolve("userService")
```

### Challenge 3: Event Bus

Build event-driven system:
```go
type EventBus struct {
    // Implement pub/sub pattern
}

func NewEventBus() *EventBus
func (eb *EventBus) Subscribe(topic string, handler func(data interface{}))
func (eb *EventBus) Publish(topic string, data interface{})
func (eb *EventBus) Unsubscribe(topic string, handler func(data interface{}))
```

**Features:**
- Multiple subscribers per topic
- Async event handling
- Priority subscribers
- Event history
- Wildcard topics

### Challenge 4: High-Performance JSON Parser

Build JSON parser nhanh hơn `encoding/json`:
```go
func Parse(data []byte) (interface{}, error)
func ParseObject(data []byte) (map[string]interface{}, error)
func ParseArray(data []byte) ([]interface{}, error)
```

**Requirements:**
- Zero allocation khi có thể
- Streaming support
- Faster than stdlib
- Benchmark comparison

---

## 📊 Evaluation Criteria

### Code Quality
- [ ] Clean code, readable
- [ ] Proper error handling
- [ ] Good naming conventions
- [ ] Comments for complex logic

### Performance
- [ ] Benchmarks provided
- [ ] Memory efficient
- [ ] No unnecessary allocations
- [ ] Proper use of sync primitives

### Testing
- [ ] Unit tests
- [ ] Edge cases covered
- [ ] Table-driven tests
- [ ] Race detection (-race flag)

### Best Practices
- [ ] Follow Go idioms
- [ ] Proper package structure
- [ ] godoc comments
- [ ] No global state (when possible)

---

## 💡 Tips

1. **Reflection:** Dùng khi thực sự cần, cache reflection results
2. **Generics:** Prefer khi logic giống nhau cho nhiều types
3. **Patterns:** Không over-engineer, giữ code đơn giản
4. **Performance:** Profile trước khi optimize
5. **Testing:** Viết tests trước khi implement (TDD)

---

## 📚 Resources

- [Go Reflection](https://go.dev/blog/laws-of-reflection)
- [Go Generics Tutorial](https://go.dev/doc/tutorial/generics)
- [Go Design Patterns](https://github.com/tmrts/go-patterns)
- [High Performance Go](https://dave.cheney.net/high-performance-go-workshop)

---

**Good luck! 🚀**
