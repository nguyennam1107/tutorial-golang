package main

import (
	"fmt"
)
// ============= BÀI 7: METHODS & INTERFACES =============

// TODO: Define Shape interface
type Shape interface {
	Area() float64
}

// TODO: Implement Circle
type Circle struct {
	Radius float64
}


// TODO: Implement Rectangle
type Rectangle struct {
	Width  float64
	Height float64
}

// TODO: Implement Triangle
type Triangle struct {
	Base   float64
	Height float64
}

// TODO: Tính tổng diện tích
func totalArea(shapes []Shape) float64 {
	// Implement here
	var total float64
	for _, shape := range shapes {
		total += shape.Area()
	}
	return total
}

func (c Circle) Area() float64 {
	return 3.14159 * c.Radius * c.Radius
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (t Triangle) Area() float64 {
	return 0.5 * t.Base * t.Height
}

func exercise7() {
	fmt.Println("\n=== BÀI 7: METHODS & INTERFACES ===")

	shapes := []Shape{
		Circle{Radius: 5},
		Rectangle{Width: 10, Height: 5},
		Triangle{Base: 6, Height: 4},
	}

	fmt.Printf("Total area: %.2f\n", totalArea(shapes))
	fmt.Println("Total circle area:", shapes[0].Area())
	fmt.Println("Total rectangle area:", shapes[1].Area())
	fmt.Println("Total triangle area:", shapes[2].Area())
}

func main() {
	exercise7()
}