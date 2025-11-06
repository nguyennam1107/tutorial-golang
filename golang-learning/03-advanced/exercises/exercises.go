package main

import (
	"context"
	"errors"
	"fmt"
	"reflect"
	"sync"
	"time"
)

// ============= BÀI 1: REFLECTION BASICS =============

// TODO: Inspect struct và in ra thông tin
func inspectStruct(v interface{}) {
	// Implement here
	// Sử dụng reflect.TypeOf và reflect.ValueOf
	// In ra: type name, number of fields, field details
}

// TODO: Call method qua reflection
func callMethod(v interface{}, methodName string) error {
	// Implement here
	// Tìm method theo tên
	// Gọi method
	// Return error nếu method không tồn tại
	return nil
}

// TODO: Modify field value qua reflection
func modifyField(v interface{}, fieldName string, newValue interface{}) error {
	// Implement here
	// Check if value is pointer
	// Get field by name
	// Set new value
	return nil
}

func exercise1() {
	fmt.Println("=== BÀI 1: REFLECTION BASICS ===")

	type User struct {
		Name  string `json:"name" validate:"required"`
		Email string `json:"email" validate:"email"`
		Age   int    `json:"age" validate:"min=0,max=120"`
	}

	user := User{Name: "John", Email: "john@example.com", Age: 30}

	// TODO: Test inspectStruct
	inspectStruct(user)

	// TODO: Test modifyField
	// modifyField(&user, "Age", 31)
	// fmt.Println("Modified user:", user)
}

// ============= BÀI 2: JSON SERIALIZER =============

// TODO: Marshal struct to JSON
func Marshal(v interface{}) ([]byte, error) {
	// Implement here
	// Handle struct, map, slice, basic types
	// Respect struct tags
	return nil, errors.New("not implemented")
}

// TODO: Unmarshal JSON to struct
func Unmarshal(data []byte, v interface{}) error {
	// Implement here
	// Parse JSON string
	// Set values qua reflection
	return errors.New("not implemented")
}

func exercise2() {
	fmt.Println("\n=== BÀI 2: JSON SERIALIZER ===")

	type Address struct {
		City    string `json:"city"`
		Country string `json:"country"`
	}

	type Person struct {
		Name    string   `json:"name"`
		Age     int      `json:"age"`
		Emails  []string `json:"emails,omitempty"`
		Address *Address `json:"address,omitempty"`
	}

	// TODO: Test Marshal
	// p := Person{Name: "John", Age: 30}
	// data, err := Marshal(p)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(string(data))

	// TODO: Test Unmarshal
	// var p2 Person
	// err = Unmarshal(data, &p2)
	// fmt.Println("Unmarshaled:", p2)
}

// ============= BÀI 3: GENERIC DATA STRUCTURES =============

// TODO: Generic Stack
type Stack[T any] struct {
	// Implement here
}

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{}
}

func (s *Stack[T]) Push(item T) {
	// Implement here
}

func (s *Stack[T]) Pop() (T, error) {
	// Implement here
	var zero T
	return zero, errors.New("stack is empty")
}

func (s *Stack[T]) Peek() (T, error) {
	// Implement here
	var zero T
	return zero, errors.New("stack is empty")
}

func (s *Stack[T]) IsEmpty() bool {
	// Implement here
	return true
}

func (s *Stack[T]) Size() int {
	// Implement here
	return 0
}

// TODO: Generic Queue
type Queue[T any] struct {
	// Implement here
}

func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{}
}

func (q *Queue[T]) Enqueue(item T) {
	// Implement here
}

func (q *Queue[T]) Dequeue() (T, error) {
	// Implement here
	var zero T
	return zero, errors.New("queue is empty")
}

func (q *Queue[T]) Front() (T, error) {
	// Implement here
	var zero T
	return zero, errors.New("queue is empty")
}

func (q *Queue[T]) IsEmpty() bool {
	// Implement here
	return true
}

func (q *Queue[T]) Size() int {
	// Implement here
	return 0
}

// TODO: Generic LinkedList
type Node[T any] struct {
	Value T
	Next  *Node[T]
}

type LinkedList[T any] struct {
	// Implement here
}

func NewLinkedList[T any]() *LinkedList[T] {
	return &LinkedList[T]{}
}

func (ll *LinkedList[T]) Append(value T) {
	// Implement here
}

func (ll *LinkedList[T]) Prepend(value T) {
	// Implement here
}

func (ll *LinkedList[T]) Remove(value T) bool {
	// Implement here
	return false
}

func (ll *LinkedList[T]) Find(value T) bool {
	// Implement here
	return false
}

func (ll *LinkedList[T]) ToSlice() []T {
	// Implement here
	return nil
}

func exercise3() {
	fmt.Println("\n=== BÀI 3: GENERIC DATA STRUCTURES ===")

	// TODO: Test Stack
	fmt.Println("Stack test:")
	// stack := NewStack[int]()
	// stack.Push(1)
	// stack.Push(2)
	// stack.Push(3)
	// fmt.Println("Size:", stack.Size())
	// val, _ := stack.Pop()
	// fmt.Println("Popped:", val)

	// TODO: Test Queue
	fmt.Println("\nQueue test:")
	// queue := NewQueue[string]()
	// queue.Enqueue("first")
	// queue.Enqueue("second")
	// val, _ := queue.Dequeue()
	// fmt.Println("Dequeued:", val)

	// TODO: Test LinkedList
	fmt.Println("\nLinkedList test:")
	// list := NewLinkedList[int]()
	// list.Append(1)
	// list.Append(2)
	// list.Prepend(0)
	// fmt.Println("List:", list.ToSlice())
}

// ============= BÀI 4: GENERIC FUNCTIONS =============

// TODO: Map function
func Map[T, U any](slice []T, fn func(T) U) []U {
	// Implement here
	return nil
}

// TODO: Filter function
func Filter[T any](slice []T, fn func(T) bool) []T {
	// Implement here
	return nil
}

// TODO: Reduce function
func Reduce[T, U any](slice []T, initial U, fn func(U, T) U) U {
	// Implement here
	var zero U
	return zero
}

// TODO: Find function
func Find[T any](slice []T, fn func(T) bool) (T, bool) {
	// Implement here
	var zero T
	return zero, false
}

// TODO: Contains function
func Contains[T comparable](slice []T, item T) bool {
	// Implement here
	return false
}

// TODO: Reverse function
func Reverse[T any](slice []T) []T {
	// Implement here
	return nil
}

// TODO: Min function
func Min[T interface{ ~int | ~float64 | ~string }](items ...T) T {
	// Implement here
	var zero T
	return zero
}

// TODO: Max function
func Max[T interface{ ~int | ~float64 | ~string }](items ...T) T {
	// Implement here
	var zero T
	return zero
}

func exercise4() {
	fmt.Println("\n=== BÀI 4: GENERIC FUNCTIONS ===")

	nums := []int{1, 2, 3, 4, 5}

	// TODO: Test Map
	// squared := Map(nums, func(n int) int { return n * n })
	// fmt.Println("Squared:", squared)

	// TODO: Test Filter
	// evens := Filter(nums, func(n int) bool { return n%2 == 0 })
	// fmt.Println("Evens:", evens)

	// TODO: Test Reduce
	// sum := Reduce(nums, 0, func(acc, n int) int { return acc + n })
	// fmt.Println("Sum:", sum)

	// TODO: Test Contains
	// fmt.Println("Contains 3:", Contains(nums, 3))
	// fmt.Println("Contains 10:", Contains(nums, 10))

	// TODO: Test Min/Max
	// fmt.Println("Min:", Min(1, 2, 3, 4, 5))
	// fmt.Println("Max:", Max(1, 2, 3, 4, 5))
}

// ============= BÀI 5: GENERIC CACHE =============

// TODO: Generic Cache
type CacheItem[V any] struct {
	value      V
	expireTime time.Time
}

type Cache[K comparable, V any] struct {
	// Implement here
	// items map[K]CacheItem[V]
	// mu sync.RWMutex
	// maxSize int
}

func NewCache[K comparable, V any]() *Cache[K, V] {
	// Implement here
	return &Cache[K, V]{}
}

func (c *Cache[K, V]) Set(key K, value V) {
	// Implement here
}

func (c *Cache[K, V]) Get(key K) (V, bool) {
	// Implement here
	var zero V
	return zero, false
}

func (c *Cache[K, V]) SetWithTTL(key K, value V, ttl time.Duration) {
	// Implement here
	// Set expireTime = now + ttl
}

func (c *Cache[K, V]) Delete(key K) {
	// Implement here
}

func (c *Cache[K, V]) Clear() {
	// Implement here
}

func (c *Cache[K, V]) Size() int {
	// Implement here
	return 0
}

func (c *Cache[K, V]) SetMaxSize(maxSize int) {
	// Implement here
	// Implement LRU eviction
}

func exercise5() {
	fmt.Println("\n=== BÀI 5: GENERIC CACHE ===")

	// TODO: Test basic operations
	// cache := NewCache[string, int]()
	// cache.Set("one", 1)
	// cache.Set("two", 2)
	// val, ok := cache.Get("one")
	// fmt.Printf("Get 'one': %d, %v\n", val, ok)

	// TODO: Test TTL
	// cache.SetWithTTL("temp", 42, 100*time.Millisecond)
	// val, ok := cache.Get("temp")
	// fmt.Printf("Before expiry: %d, %v\n", val, ok)
	// time.Sleep(150 * time.Millisecond)
	// val, ok = cache.Get("temp")
	// fmt.Printf("After expiry: %d, %v\n", val, ok)
}

// ============= BÀI 6: SINGLETON PATTERN =============

// TODO: Implement Singleton
type Database struct {
	connectionString string
}

var (
	dbInstance *Database
	dbOnce     sync.Once
)

func GetDatabase() *Database {
	// Implement here
	// Sử dụng sync.Once
	return nil
}

func (db *Database) Connect(connStr string) {
	db.connectionString = connStr
	fmt.Println("Connected to:", connStr)
}

func exercise6() {
	fmt.Println("\n=== BÀI 6: SINGLETON PATTERN ===")

	// TODO: Test singleton
	// db1 := GetDatabase()
	// db1.Connect("localhost:5432")

	// db2 := GetDatabase()
	// fmt.Println("Same instance?", db1 == db2)
}

// ============= BÀI 7: FACTORY PATTERN =============

// TODO: Factory Pattern
type Logger interface {
	Log(message string)
}

type ConsoleLogger struct{}

func (cl ConsoleLogger) Log(message string) {
	fmt.Println("[CONSOLE]", message)
}

type FileLogger struct {
	filename string
}

func (fl FileLogger) Log(message string) {
	fmt.Printf("[FILE:%s] %s\n", fl.filename, message)
}

// TODO: Logger factory
func NewLogger(loggerType string, params ...string) Logger {
	// Implement here
	// "console" -> ConsoleLogger
	// "file" -> FileLogger
	return nil
}

func exercise7() {
	fmt.Println("\n=== BÀI 7: FACTORY PATTERN ===")

	// TODO: Test factory
	// logger1 := NewLogger("console")
	// logger1.Log("Hello from console")

	// logger2 := NewLogger("file", "app.log")
	// logger2.Log("Hello from file")
}

// ============= BÀI 8: BUILDER PATTERN =============

// TODO: Builder Pattern
type HTTPRequest struct {
	URL     string
	Method  string
	Headers map[string]string
	Body    []byte
}

type RequestBuilder struct {
	request HTTPRequest
}

func NewRequestBuilder() *RequestBuilder {
	return &RequestBuilder{
		request: HTTPRequest{
			Headers: make(map[string]string),
		},
	}
}

func (b *RequestBuilder) URL(url string) *RequestBuilder {
	// Implement here
	return b
}

func (b *RequestBuilder) Method(method string) *RequestBuilder {
	// Implement here
	return b
}

func (b *RequestBuilder) Header(key, value string) *RequestBuilder {
	// Implement here
	return b
}

func (b *RequestBuilder) Body(body []byte) *RequestBuilder {
	// Implement here
	return b
}

func (b *RequestBuilder) Build() *HTTPRequest {
	return &b.request
}

func exercise8() {
	fmt.Println("\n=== BÀI 8: BUILDER PATTERN ===")

	// TODO: Test builder
	// req := NewRequestBuilder().
	// 	URL("https://api.example.com").
	// 	Method("POST").
	// 	Header("Content-Type", "application/json").
	// 	Body([]byte(`{"name":"John"}`)).
	// 	Build()

	// fmt.Printf("Request: %+v\n", req)
}

// ============= BÀI 9: OBSERVER PATTERN =============

// TODO: Observer Pattern
type Observer interface {
	Update(data interface{})
}

type Subject interface {
	Attach(observer Observer)
	Detach(observer Observer)
	Notify(data interface{})
}

// TODO: Concrete Subject
type NewsAgency struct {
	observers []Observer
}

func (na *NewsAgency) Attach(observer Observer) {
	// Implement here
}

func (na *NewsAgency) Detach(observer Observer) {
	// Implement here
}

func (na *NewsAgency) Notify(data interface{}) {
	// Implement here
}

func (na *NewsAgency) PublishNews(news string) {
	fmt.Println("Publishing:", news)
	na.Notify(news)
}

// TODO: Concrete Observer
type NewsChannel struct {
	name string
}

func (nc NewsChannel) Update(data interface{}) {
	fmt.Printf("[%s] Received: %v\n", nc.name, data)
}

func exercise9() {
	fmt.Println("\n=== BÀI 9: OBSERVER PATTERN ===")

	// TODO: Test observer
	// agency := &NewsAgency{}

	// channel1 := NewsChannel{name: "CNN"}
	// channel2 := NewsChannel{name: "BBC"}

	// agency.Attach(channel1)
	// agency.Attach(channel2)

	// agency.PublishNews("Breaking news!")
}

// ============= BÀI 10: RATE LIMITER =============

// TODO: Token Bucket Rate Limiter
type RateLimiter struct {
	rate   int           // tokens per second
	burst  int           // max tokens
	tokens int           // current tokens
	mu     sync.Mutex    // protect tokens
	ticker *time.Ticker  // refill ticker
}

func NewRateLimiter(rate int, burst int) *RateLimiter {
	// Implement here
	// Start goroutine để refill tokens
	return nil
}

func (rl *RateLimiter) Allow() bool {
	// Implement here
	// Check if token available
	// Consume token if yes
	return false
}

func (rl *RateLimiter) Wait(ctx context.Context) error {
	// Implement here
	// Wait until token available
	// Respect context cancellation
	return nil
}

func exercise10() {
	fmt.Println("\n=== BÀI 10: RATE LIMITER ===")

	// TODO: Test rate limiter
	// limiter := NewRateLimiter(5, 10) // 5 req/sec, burst 10

	// for i := 0; i < 20; i++ {
	// 	if limiter.Allow() {
	// 		fmt.Printf("Request %d: ALLOWED at %v\n", i, time.Now())
	// 	} else {
	// 		fmt.Printf("Request %d: DENIED\n", i)
	// 	}
	// 	time.Sleep(100 * time.Millisecond)
	// }
}

// ============= MAIN =============

func main() {
	exercise1()
	exercise2()
	exercise3()
	exercise4()
	exercise5()
	exercise6()
	exercise7()
	exercise8()
	exercise9()
	exercise10()
}
