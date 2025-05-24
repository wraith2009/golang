package main

import "fmt"

// defining an interface
type shape interface {
	area() float32
	perimeter() float32
}

// defining a structure
type rect struct {
	width  float32
	height float32
}

// implementing area method
func (r rect) area() float32 {
	return r.width * r.height
}

// implementing perimeter method
func (r rect) perimeter() float32 {
	return 2 * (r.width + r.height)
}

func main() {
	var s shape                // Declare a variable of type 'shape' interface
	s = rect{width: 20, height: 10} // Assign a 'rect' instance to the interface

	fmt.Println("Area:", s.area())
	fmt.Println("Perimeter:", s.perimeter())
}
