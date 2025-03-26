package main

import (
	"fmt"
	"math"
)

// Shape interface defines a method for calculating area.
type Shape interface {
	Area() float64
}

// Rectangle represents a rectangle with a width and height.
type Rectangle struct {
	Width, Height float64
}

// NewRectangle creates a new rectangle instance.
func NewRectangle(width, height float64) *Rectangle {
	return &Rectangle{Width: width, Height: height}
}

// Area calculates the area of a rectangle.
func (r *Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Square represents a square with a side length.
// It is modeled as a distinct type to avoid unwanted behavior inherited from Rectangle.
type Square struct {
	Side float64
}

// NewSquare creates a new square instance.
func NewSquare(side float64) *Square {
	return &Square{Side: side}
}

// Area calculates the area of a square.
func (s *Square) Area() float64 {
	return math.Pow(s.Side, 2)
}

// DisplayArea prints the area of any shape.
func DisplayArea(s Shape) {
	fmt.Printf("Area: %.2f\n", s.Area())
}

func main() {
	// Create a rectangle and a square.
	rect := NewRectangle(5, 10)
	square := NewSquare(7)

	// Display areas for both shapes using the Shape interface.
	fmt.Println("Rectangle:")
	DisplayArea(rect)

	fmt.Println("Square:")
	DisplayArea(square)
}
