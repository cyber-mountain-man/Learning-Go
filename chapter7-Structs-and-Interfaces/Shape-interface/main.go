package main

import	(
	"fmt"  // Importing the fmt package for formatted output
	"math" // Importing the math package to use mathematical functions (e.g., Pi)
)

// Shape interface: Any type that implements area() satisfies Shape
type Shape interface {
	area() float64 // Defines a method signature that must return a float64 value
}

// Rectangle struct represents a rectangle with width and height
type Rectangle struct {
	width, height float64
}

// Circle struct represents a circle with a radius
type Circle struct {
	radius float64
}

// Method to calculate the area of a Rectangle
func (r Rectangle) area() float64 {
	return r.width * r.height // Area formula: width × height
}

// Method to calculate the area of a Circle
func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius // Area formula: π × radius²
}

// Function to print the area of any shape (uses the interface)
func printArea(s Shape) {
	fmt.Println("Area:", s.area()) // Calls the area() method on any Shape type
}

func main() {
	// Create a Rectangle with width 10 and height 5
	rect := Rectangle{width: 10, height: 5}

	// Create a Circle with radius 7
	circ := Circle{radius: 7}
	
	// Print areas using the interface
	printArea(rect) // Works because Rectangle implements area()
	printArea(circ) // Works because Circle implements area()
}

/*
✅ Let’s walk through the program:

1️⃣ **Defining an Interface**
   - `Shape` is an **interface** that requires any type that implements it to have an `area()` method returning a `float64`.
   - The interface allows **different shapes** to be treated the same way.

2️⃣ **Defining Structs**
   - `Rectangle` struct stores `width` and `height` as float64 values.
   - `Circle` struct stores `radius` as a float64 value.

3️⃣ **Implementing the `Shape` Interface**
   - `Rectangle` has a method `area()`, which calculates `width × height`.
   - `Circle` has a method `area()`, which calculates `π × radius²` using `math.Pi`.

4️⃣ **Using Interfaces in a Function**
   - `printArea(s Shape)` accepts **any type that implements `area()`**.
   - Since `Rectangle` and `Circle` both implement `area()`, they satisfy `Shape`.
   - Calling `printArea(rect)` and `printArea(circ)` automatically calls the correct method.

5️⃣ **Execution Flow**
   - `rect := Rectangle{width: 10, height: 5}` → Creates a rectangle.
   - `circ := Circle{radius: 7}` → Creates a circle.
   - `printArea(rect)` → Calls `Rectangle.area()`, prints **"Area: 50"**.
   - `printArea(circ)` → Calls `Circle.area()`, prints **"Area: 153.93804002589985"**.

6️⃣ **Expected Output**
		Area: 50 Area: 153.93804002589985

7️⃣ **Why Use Interfaces Here?**
- ✅ **Flexibility** – We can add more shapes (e.g., `Triangle`) without changing `printArea()`.
- ✅ **Encapsulation** – Each shape **handles its own logic**, making the code clean.
- ✅ **Reusability** – The `printArea()` function works for **any shape**, reducing duplicate code.
*/