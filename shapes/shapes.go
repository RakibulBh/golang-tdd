package shapes

import "math"

type Shape interface {
	Area() float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

type Circle struct {
	Radius float64
}

// Calculate the perimeter of a rectangle
func (s *Rectangle) Perimeter() float64 {
	return 2 * (s.Height + s.Width)
}

// Calculate the area of a rectangle
func (s *Rectangle) Area() float64 {
	return s.Height * s.Width
}

// Calculate the perimeter of a circle
func (s *Circle) Perimeter() float64 {
	return 2 * math.Pi * s.Radius
}

// Calculate the area of a circle
func (s *Circle) Area() float64 {
	return math.Pi * math.Pow(s.Radius, 2)
}
