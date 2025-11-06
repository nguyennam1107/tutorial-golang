package main
import (
	"errors"
	"fmt"
)
// ============= BÀI 3: GENERIC DATA STRUCTURES =============

// TODO: Generic Stack
type Stack[T any] struct {
	// Implement here
	items []T
}

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{}
}

func (s *Stack[T]) Push(item T) {
	// Implement here
	s.items = append(s.items, item)
}

func (s *Stack[T]) Pop() (T, error) {
	// Implement here
	if len(s.items) == 0 {
		var zero T
		return zero, errors.New("stack is empty")
	}
	item := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return item, nil
}

func (s *Stack[T]) Peek() (T, error) {
	// Implement here
	if len(s.items) == 0 {
		var zero T
		return zero, errors.New("stack is empty")
	}
	return s.items[len(s.items)-1], nil
}

func (s *Stack[T]) IsEmpty() bool {
	// Implement here
	return len(s.items) == 0
}

func (s *Stack[T]) Size() int {
	// Implement here
	return len(s.items)
}

// TODO: Generic Queue
type Queue[T any] struct {
	// Implement here
	items []T
}


func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{}
}

func (q *Queue[T]) Enqueue(item T) {
	// Implement here
	q.items = append(q.items, item)
}

func (q *Queue[T]) Dequeue() (T, error) {
	// Implement here
	if len(q.items) == 0 {
		var zero T
		return zero, errors.New("queue is empty")
	}
	item := q.items[0]
	q.items = q.items[1:]
	return item, nil
}

func (q *Queue[T]) Front() (T, error) {
	// Implement here
	if len(q.items) == 0 {
		var zero T
		return zero, errors.New("queue is empty")
	}
	return q.items[0], nil
}


func (q *Queue[T]) IsEmpty() bool {
	// Implement here
	return len(q.items) == 0
}

func (q *Queue[T]) Size() int {
	// Implement here
	return len(q.items)
}

// TODO: Generic LinkedList
type Node[T any] struct {
	Value T
	Next  *Node[T]
}

type LinkedList[T comparable] struct {
	// Implement here
	head *Node[T]
	tail *Node[T]
	size int
}

func NewLinkedList[T comparable]() *LinkedList[T] {
	return &LinkedList[T]{}
}

func (ll *LinkedList[T]) Append(value T) {
	// Implement here
	newNode := &Node[T]{Value: value}
	if ll.head == nil {
		ll.head = newNode
		ll.tail = newNode
	} else {
		ll.tail.Next = newNode
		ll.tail = newNode
	}
	ll.size++
}

func (ll *LinkedList[T]) Prepend(value T) {
	// Implement here
	newNode := &Node[T]{Value: value}
	if ll.head == nil {
		ll.head = newNode
		ll.tail = newNode
	} else {
		newNode.Next = ll.head
		ll.head = newNode
	}
	ll.size++
}

func (ll *LinkedList[T]) Remove(value T) bool {
	// Implement here
	if ll.head == nil {
		return false
	}
	if ll.head.Value == value {
		ll.head = ll.head.Next
		ll.size--
		return true
	}
	current := ll.head
	for current.Next != nil {
		if current.Next.Value == value {
			current.Next = current.Next.Next
			ll.size--
			return true
		}
		current = current.Next
	}
	return false
}

func (ll *LinkedList[T]) Find(value T) bool {
	// Implement here
	current := ll.head
	for current != nil {
		if current.Value == value {
			return true
		}
		current = current.Next
	}
	return false
}

func (ll *LinkedList[T]) ToSlice() []T {
	// Implement here
	var result []T
	current := ll.head
	for current != nil {
		result = append(result, current.Value)
		current = current.Next
	}
	return result
}


func exercise3() {
	fmt.Println("\n=== BÀI 3: GENERIC DATA STRUCTURES ===")

	// TODO: Test Stack
	fmt.Println("Stack test:")
	stack := NewStack[int]()
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	fmt.Println("Size:", stack.Size())
	val, _ := stack.Pop()
	fmt.Println("Popped:", val)

	// TODO: Test Queue
	fmt.Println("\nQueue test:")
	queue := NewQueue[string]()
	queue.Enqueue("first")
	queue.Enqueue("second")
	valStr, _ := queue.Dequeue()
	fmt.Println("Dequeued:", valStr)

	// TODO: Test LinkedList
	fmt.Println("\nLinkedList test:")
	list := NewLinkedList[int]()
	list.Append(1)
	list.Append(2)
	list.Prepend(0)
	fmt.Println("List:", list.ToSlice())
}
func main() {
	exercise3()
}