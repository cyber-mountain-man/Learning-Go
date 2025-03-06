package main

import (
	"fmt"
	"math"
)

// Rectangle struct to represent a rectangle with two diagonal points
type Rectangle struct {
	x1, y1, x2, y2 float64
}

// Circle struct to represent a circle with a given radius
type Circle struct {
	radius float64
}

// Function to calculate the Euclidean distance between two points
func distance(x1, y1, x2, y2 float64) float64 {
	a := x2 - x1
	b := y2 - y1
	return math.Sqrt(a*a + b*b) // Uses the Pythagorean theorem to compute distance
}

// Method to calculate the area of a rectangle using the distance function
func (r Rectangle) area() float64 {
	l := distance(r.x1, r.y1, r.x1, r.y2) // Calculates the height of the rectangle
	w := distance(r.x1, r.y1, r.x2, r.y1) // Calculates the width of the rectangle
	return l * w // Area = length × width
}

// Method to calculate the area of a circle
func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius // Formula for the area of a circle: πr²
}

func main() {
	// Create a rectangle with two opposite corner coordinates
	rect := Rectangle{x1: 0, y1: 0, x2: 10, y2: 10}

	// Create a circle with a given radius
	circ := Circle{radius: 5}

	// Calculate and print the areas of the rectangle and circle
	fmt.Println("Rectangle Area:", rect.area()) // Expected output: 100
	fmt.Println("Circle Area:", circ.area())    // Expected output: 78.54 (approximately)
}

/*
✅ Let’s walk through the program:

1️⃣ **Defining the Structs**
   - **Rectangle** stores four float64 values representing the coordinates of two opposite corners.
   - **Circle** stores one float64 value representing the radius.

2️⃣ **Calculating Distances**
   - The `distance()` function calculates the Euclidean distance between two points using the **Pythagorean theorem**.

3️⃣ **Computing Areas**
   - `Rectangle.area()` calculates:
     - **Length** as the vertical distance between `y1` and `y2`.
     - **Width** as the horizontal distance between `x1` and `x2`.
     - **Final area** as `length × width`.
   - `Circle.area()` uses the standard formula: **πr²**.

4️⃣ **Execution Flow**
   - A **Rectangle** with corners `(0,0)` and `(10,10)` is created.
   - A **Circle** with `radius = 5` is created.
   - The program **calls the `area()` methods** to calculate and print the areas.

5️⃣ **Expected Output**
	Rectangle Area: 100   Circle Area: 78.53981633974483


6️⃣ **Why Use Methods Instead of Standalone Functions?**
- ✅ **Encapsulation** – The area logic is directly associated with the `Rectangle` and `Circle` types.
- ✅ **Improves Readability** – We can call `rect.area()` instead of passing `rect` to a function.
- ✅ **Real-World Applications** – Similar to how object-oriented programming associates behaviors with objects.
*/