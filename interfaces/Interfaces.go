package interfaces

import (
	"fmt"
	"math"
)

// Interfaces and Polymorphism in Go

// 1. Defining an Interface:
// An interface is a type that specifies a set of method signatures.
// A type is said to implement an interface if it provides the method(s) that the interface requires.

type Shape interface { // Interface 'Shape' with two methods: Area and Perimeter
	Area() float64
	Perimeter() float64
}

// 2. Implementing an Interface:
// The `Circle` struct implements the `Shape` interface by defining the `Area` and `Perimeter` methods.

type Circle struct { // Struct type 'Circle' with a radius field
	Radius float64
}

// Method to calculate the area of the circle
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// Method to calculate the perimeter (circumference) of the circle
func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

// 3. Another Implementation of the Interface:
// The `Rectangle` struct also implements the `Shape` interface with its own `Area` and `Perimeter` methods.

type Rectangle struct { // Struct type 'Rectangle' with length and width fields
	Length, Width float64
}

// Method to calculate the area of the rectangle
func (r Rectangle) Area() float64 {
	return r.Length * r.Width
}

// Method to calculate the perimeter of the rectangle
func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Length + r.Width)
}

// 4. Polymorphism with Interfaces:
// Polymorphism in Go is achieved using interfaces. You can write functions that accept interface types, allowing any struct that implements that interface to be used.

func PrintShapeInfo(s Shape) { // Function accepting the 'Shape' interface
	fmt.Printf("Area: %.2f\n", s.Area())
	fmt.Printf("Perimeter: %.2f\n", s.Perimeter())
}

func main() {
	// 5. Using Interfaces and Polymorphism:

	// Creating instances of Circle and Rectangle
	circle := Circle{Radius: 5}
	rectangle := Rectangle{Length: 4, Width: 3}

	// Using the same function to handle different types that implement the Shape interface
	fmt.Println("Circle Info:")
	PrintShapeInfo(circle) // Output will use Circle's Area and Perimeter methods

	fmt.Println("\nRectangle Info:")
	PrintShapeInfo(rectangle) // Output will use Rectangle's Area and Perimeter methods

	// Key Insights about Interfaces in Go:
	// - Implicit implementation: Types don't explicitly declare they implement an interface; if they have the required methods, they do.
	// - Interface composition: Interfaces can be composed using other interfaces.
	// - No generics: Go uses interfaces for code reusability and polymorphism instead of generics.
}
