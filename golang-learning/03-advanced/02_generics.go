package main

import (
	"fmt"
	"golang.org/x/exp/constraints"
)

// ============================================
// GENERIC STACK
// ============================================

type Stack[T any] struct {
	items []T
}

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{
		items: make([]T, 0),
	}
}

func (s *Stack[T]) Push(item T) {
	s.items = append(s.items, item)
}

func (s *Stack[T]) Pop() (T, bool) {
	if len(s.items) == 0 {
		var zero T
		return zero, false
	}
	item := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return item, true
}

func (s *Stack[T]) Peek() (T, bool) {
	if len(s.items) == 0 {
		var zero T
		return zero, false
	}
	return s.items[len(s.items)-1], true
}

func (s *Stack[T]) IsEmpty() bool {
	return len(s.items) == 0
}

func (s *Stack[T]) Size() int {
	return len(s.items)
}

// ============================================
// GENERIC MAP FUNCTIONS
// ============================================

func Map[T, U any](slice []T, fn func(T) U) []U {
	result := make([]U, len(slice))
	for i, v := range slice {
		result[i] = fn(v)
	}
	return result
}

func Filter[T any](slice []T, fn func(T) bool) []T {
	result := make([]T, 0)
	for _, v := range slice {
		if fn(v) {
			result = append(result, v)
		}
	}
	return result
}

func Reduce[T, U any](slice []T, initial U, fn func(U, T) U) U {
	result := initial
	for _, v := range slice {
		result = fn(result, v)
	}
	return result
}

func Find[T any](slice []T, fn func(T) bool) (T, bool) {
	for _, v := range slice {
		if fn(v) {
			return v, true
		}
	}
	var zero T
	return zero, false
}

// ============================================
// GENERIC CONSTRAINTS
// ============================================

// Ordered constraint
func Min[T constraints.Ordered](a, b T) T {
	if a < b {
		return a
	}
	return b
}

func Max[T constraints.Ordered](a, b T) T {
	if a > b {
		return a
	}
	return b
}

// Custom constraint
type Number interface {
	~int | ~int64 | ~float32 | ~float64
}

func Sum[T Number](numbers []T) T {
	var sum T
	for _, n := range numbers {
		sum += n
	}
	return sum
}

func Average[T Number](numbers []T) float64 {
	if len(numbers) == 0 {
		return 0
	}
	sum := Sum(numbers)
	return float64(sum) / float64(len(numbers))
}

// ============================================
// GENERIC PAIR
// ============================================

type Pair[K, V any] struct {
	Key   K
	Value V
}

func NewPair[K, V any](key K, value V) Pair[K, V] {
	return Pair[K, V]{Key: key, Value: value}
}

func (p Pair[K, V]) String() string {
	return fmt.Sprintf("(%v, %v)", p.Key, p.Value)
}

// ============================================
// GENERIC OPTIONAL
// ============================================

type Optional[T any] struct {
	value *T
}

func Some[T any](value T) Optional[T] {
	return Optional[T]{value: &value}
}

func None[T any]() Optional[T] {
	return Optional[T]{value: nil}
}

func (o Optional[T]) IsPresent() bool {
	return o.value != nil
}

func (o Optional[T]) Get() (T, bool) {
	if o.value == nil {
		var zero T
		return zero, false
	}
	return *o.value, true
}

func (o Optional[T]) OrElse(defaultValue T) T {
	if o.value == nil {
		return defaultValue
	}
	return *o.value
}

// ============================================
// GENERIC RESULT (for error handling)
// ============================================

type Result[T any] struct {
	value T
	err   error
}

func Ok[T any](value T) Result[T] {
	return Result[T]{value: value, err: nil}
}

func Err[T any](err error) Result[T] {
	var zero T
	return Result[T]{value: zero, err: err}
}

func (r Result[T]) IsOk() bool {
	return r.err == nil
}

func (r Result[T]) IsErr() bool {
	return r.err != nil
}

func (r Result[T]) Unwrap() (T, error) {
	return r.value, r.err
}

func (r Result[T]) Expect(msg string) T {
	if r.err != nil {
		panic(msg + ": " + r.err.Error())
	}
	return r.value
}

// ============================================
// EXAMPLES
// ============================================

func main() {
	fmt.Println("=== GENERIC STACK ===")
	stack := NewStack[int]()
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	fmt.Println("Size:", stack.Size())

	for !stack.IsEmpty() {
		item, _ := stack.Pop()
		fmt.Println("Popped:", item)
	}

	fmt.Println("\n=== MAP/FILTER/REDUCE ===")
	numbers := []int{1, 2, 3, 4, 5}

	// Map
	squared := Map(numbers, func(n int) int {
		return n * n
	})
	fmt.Println("Squared:", squared)

	// Filter
	evens := Filter(numbers, func(n int) bool {
		return n%2 == 0
	})
	fmt.Println("Evens:", evens)

	// Reduce
	sum := Reduce(numbers, 0, func(acc, n int) int {
		return acc + n
	})
	fmt.Println("Sum:", sum)

	// Find
	found, ok := Find(numbers, func(n int) bool {
		return n > 3
	})
	fmt.Printf("First > 3: %d, found: %v\n", found, ok)

	fmt.Println("\n=== MIN/MAX ===")
	fmt.Println("Min(3, 5):", Min(3, 5))
	fmt.Println("Max(3, 5):", Max(3, 5))
	fmt.Println("Min(3.14, 2.71):", Min(3.14, 2.71))

	fmt.Println("\n=== SUM/AVERAGE ===")
	intNums := []int{1, 2, 3, 4, 5}
	floatNums := []float64{1.5, 2.5, 3.5}

	fmt.Println("Int sum:", Sum(intNums))
	fmt.Println("Float sum:", Sum(floatNums))
	fmt.Println("Int average:", Average(intNums))
	fmt.Println("Float average:", Average(floatNums))

	fmt.Println("\n=== PAIR ===")
	p1 := NewPair("name", "John")
	p2 := NewPair("age", 30)
	fmt.Println(p1)
	fmt.Println(p2)

	fmt.Println("\n=== OPTIONAL ===")
	some := Some(42)
	none := None[int]()

	fmt.Println("Some present?", some.IsPresent())
	fmt.Println("None present?", none.IsPresent())

	if val, ok := some.Get(); ok {
		fmt.Println("Some value:", val)
	}

	fmt.Println("None or else:", none.OrElse(100))

	fmt.Println("\n=== RESULT ===")
	okResult := Ok(42)
	errResult := Err[int](fmt.Errorf("something went wrong"))

	fmt.Println("Ok result:", okResult.IsOk())
	fmt.Println("Err result:", errResult.IsErr())

	if val, err := okResult.Unwrap(); err == nil {
		fmt.Println("Ok value:", val)
	}

	if _, err := errResult.Unwrap(); err != nil {
		fmt.Println("Error:", err)
	}
}
