package shapes

import "math"

// Rectangle struct has a width and a heigth
type Rectangle struct {
	Width  float64
	Height float64
}

// Circle struct shape
type Circle struct {
	Radius float64
}

// Triangle struct shape
type Triangle struct {
	Width  float64
	Height float64
}

// Area of a Rectangle method
func (r Rectangle) Area() float64 {
	return (r.Width * r.Height)
}

// Perimeter of a Rectangle method
func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

// Area of a Circle method
func (c Circle) Area() float64 {
	return math.Pi * c.Radius
}

// Area of a Triangle method
func (t Triangle) Area() float64 {
	return (t.Width * t.Height) / 2
}
