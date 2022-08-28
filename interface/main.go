package main

import "fmt"

// Interface to be implement for struct
type shape interface {
	getArea() float64
}

// Triangle struct to calculate area : base * heigt / 2
type triangle struct {
	base   float64
	height float64
}

// Square struct to calculate area : side * side
type square struct {
	sideLength float64
}

// Implemented interface function
func printArea(s shape) {
	fmt.Println(s.getArea())
}

// Receiver function with name function same as interface shape getArea()
func (t triangle) getArea() float64 {
	return (t.base * t.height) * 0.5
}

// Receiver function with name function same as interface shape getArea()
func (s square) getArea() float64 {
	return (s.sideLength * s.sideLength)
}

func main() {
	// New struct triangle
	newTriangle := triangle{
		base:   5,
		height: 5,
	}
	// New struct square
	newSquare := square{
		sideLength: 5,
	}

	// Print all the data with the current interface
	printArea(newTriangle)
	printArea(newSquare)

}
