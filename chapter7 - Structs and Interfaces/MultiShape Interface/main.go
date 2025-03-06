package main

import (
	"fmt"  // Importing fmt package for formatted output
	"math" // Importing math package for mathematical functions like Pi
)

// Shape interface: Any type that implements area() satisfies Shape
type Shape interface {
	area() float64 // Defines a method that returns a float64 value
}

// Rectangle struct represents a rectangle with width and height
type Rectangle struct {
	width, height float64
}

// Circle struct represents a circle with a radius
type Circle struct {
	radius float64
}

// Triangle struct represents a triangle with base and height
type Triangle struct {
	base, height float64
}

// MultiShape struct: A collection that holds multiple shapes
type MultiShape struct {
	shapes []Shape // A slice that can store multiple Shape objects
}

// Method to calculate the area of a Rectangle
func (r Rectangle) area() float64 {
	return r.width * r.height // Formula: width × height
}

// Method to calculate the area of a Circle
func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius // Formula: π × radius²
}

// Method to calculate the area of a Triangle
func (t Triangle) area() float64 {
	return 0.5 * t.base * t.height // Formula: (1/2) × base × height
}

// Method to calculate the total area of all shapes inside MultiShape
func (m *MultiShape) area() float64 {
	var area float64
	for _, s := range m.shapes { // Loop through each shape in the slice
		area += s.area() // Calls the area() method of each shape and accumulates the total area
	}
	return area
}

func main() {
	// Create a Rectangle with width 10 and height 10
	rect := Rectangle{width: 10, height: 10}

	// Create a Circle with radius 5
	circ := Circle{radius: 5}

	// Create a Triangle with base 6 and height 8
	tri := Triangle{base: 6, height: 8}

	// Create a MultiShape struct that stores multiple shapes
	multiShape := MultiShape{
		shapes: []Shape{rect, circ, tri}, // Adds Rectangle, Circle, and Triangle to the MultiShape
	}

	// Calculate and print individual areas
	fmt.Println("Rectangle Area:", rect.area())       // Output: 100
	fmt.Println("Circle Area:", circ.area())          // Output: 78.54 (approx)
	fmt.Println("Triangle Area:", tri.area())         // Output: 24
	fmt.Println("Total MultiShape Area:", multiShape.area()) // Output: 202.54 (approx)
}

/*
✅ Let’s walk through the program:

1️⃣ **Defining an Interface**
   - `Shape` is an **interface** that requires any type that implements it to have an `area()` method returning a `float64`.

2️⃣ **Defining Structs**
   - `Rectangle` struct has `width` and `height` fields.
   - `Circle` struct has a `radius` field.
   - `Triangle` struct has `base` and `height` fields.
   - `MultiShape` struct stores **multiple Shape objects** in a slice (`[]Shape`).

3️⃣ **Implementing the `Shape` Interface**
   - `Rectangle` implements `area()` by calculating **width × height**.
   - `Circle` implements `area()` using **π × radius²**.
   - `Triangle` implements `area()` using **(1/2) × base × height**.
   - `MultiShape` implements `area()` by **iterating over all stored shapes and summing their areas**.

4️⃣ **Using the Interface in a Function**
   - **Each shape automatically satisfies `Shape`** since they implement `area()`.
   - The `MultiShape` struct itself also satisfies `Shape` because it implements `area()`.
   - `multiShape.area()` calls `area()` on all contained shapes and adds up the values.

5️⃣ **Execution Flow**
   - `rect := Rectangle{width: 10, height: 10}` → Creates a rectangle.
   - `circ := Circle{radius: 5}` → Creates a circle.
   - `tri := Triangle{base: 6, height: 8}` → Creates a triangle.
   - `multiShape := MultiShape{shapes: []Shape{rect, circ, tri}}` → Stores all shapes inside `MultiShape`.
   - Calling `rect.area()` → Prints **100**.
   - Calling `circ.area()` → Prints **78.54**.
   - Calling `tri.area()` → Prints **24**.
   - Calling `multiShape.area()` → Adds up **100 + 78.54 + 24 = 202.54**.

6️⃣ **Expected Output**
Rectangle Area: 100
Circle Area: 78.53981633974483
Triangle Area: 24
Total MultiShape Area: 202.53981633974483

7️⃣ **Why Use `MultiShape`?**
- ✅ **Encapsulates multiple shapes into one object**.
- ✅ **Provides an easy way to calculate the total area**.
- ✅ **Flexible—can add more shapes later without modifying existing logic**.
*/