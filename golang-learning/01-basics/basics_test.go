package basics

import "testing"

// Test cho functions
func TestAdd(t *testing.T) {
	result := 2 + 3
	expected := 5
	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}

func TestMultiply(t *testing.T) {
	tests := []struct {
		name     string
		a, b     int
		expected int
	}{
		{"positive numbers", 2, 3, 6},
		{"with zero", 5, 0, 0},
		{"negative numbers", -2, 3, -6},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.a * tt.b
			if result != tt.expected {
				t.Errorf("Expected %d, got %d", tt.expected, result)
			}
		})
	}
}

// Test cho arrays/slices
func TestSliceOperations(t *testing.T) {
	slice := []int{1, 2, 3, 4, 5}

	// Test length
	if len(slice) != 5 {
		t.Errorf("Expected length 5, got %d", len(slice))
	}

	// Test first element
	if slice[0] != 1 {
		t.Errorf("Expected first element 1, got %d", slice[0])
	}

	// Test append
	slice = append(slice, 6)
	if len(slice) != 6 {
		t.Errorf("Expected length 6 after append, got %d", len(slice))
	}
}

// Test cho maps
func TestMapOperations(t *testing.T) {
	m := make(map[string]int)
	m["apple"] = 5
	m["banana"] = 3

	if val, exists := m["apple"]; !exists || val != 5 {
		t.Errorf("Expected apple:5, got %d (exists: %t)", val, exists)
	}

	if _, exists := m["orange"]; exists {
		t.Error("orange should not exist in map")
	}
}

// Test cho structs
type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func TestRectangleArea(t *testing.T) {
	rect := Rectangle{Width: 5, Height: 3}
	expected := 15.0
	result := rect.Area()

	if result != expected {
		t.Errorf("Expected area %.2f, got %.2f", expected, result)
	}
}

// Benchmark example
func BenchmarkAppend(b *testing.B) {
	for i := 0; i < b.N; i++ {
		slice := []int{}
		for j := 0; j < 1000; j++ {
			slice = append(slice, j)
		}
	}
}

// Test error handling
func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, ErrDivisionByZero
	}
	return a / b, nil
}

var ErrDivisionByZero = error(nil)

func TestDivide(t *testing.T) {
	// Test successful division
	result, err := divide(10, 2)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if result != 5 {
		t.Errorf("Expected 5, got %.2f", result)
	}

	// Test division by zero
	_, err = divide(10, 0)
	if err == nil {
		t.Error("Expected error for division by zero")
	}
}
