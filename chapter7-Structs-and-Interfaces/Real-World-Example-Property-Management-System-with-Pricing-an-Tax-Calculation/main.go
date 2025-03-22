package main

import "fmt"

// Property interface: Defines the behavior of all property types
type Property interface {
	area() float64         // Each property must have an area calculation
	getType() string       // Returns the type of property
	getPrice() float64     // Returns price per square foot
	getTaxRate() float64   // Returns tax rate based on property type	
}

// House struct with width and length
type House struct {
	width, length float64
}

// Apartment struct with number of rooms and average room size
type Apartment struct {
	numRooms    int
	avgRoomSize float64
}

// Office struct with width, length, and number of floors
type Office struct {
	width, length float64
	floors        int
}

// Portfolio struct: A collection of different properties
type Portfolio struct {
	properties []Property // Stores multiple Property types
}

// Method to calculate the area of a House
func (h House) area() float64 {
	return h.width * h.length // Area = width × length
}

// Method to get the type of property
func (h House) getType() string {
	return "House"
}

// Method to get price per sqft for a House
func (h House) getPrice() float64 {
	return 200 // Price per sqft in dollars
}

// Method to get the tax rate for a House
func (h House) getTaxRate() float64 {
	return 0.01 // 1% property tax
}

// Method to calculate the area of an Apartment
func (a Apartment) area() float64 {
	return float64(a.numRooms) * a.avgRoomSize // Total area = room count × avg room size
}

// Method to get the type of property
func (a Apartment) getType() string {
	return "Apartment"
}

// Method to get price per sqft for an Apartment
func (a Apartment) getPrice() float64 {
	return 100 // Price per sqft in dollars
}

// Method to get tax rate for an Apartment
func (a Apartment) getTaxRate() float64 {
	return 0.015 // 1.5% property tax
}

// Method to calculate the area of an Office
func (o Office) area() float64 {
	return (o.width * o.length) * float64(o.floors) // Total area = base area × floors
}

// Method to get the type of property
func (o Office) getType() string {
	return "Office"
}

// Method to get price per sqft for an Office
func (o Office) getPrice() float64 {
	return 250 // Price per sqft in dollars
}

// Method to get tax rate for an Office
func (o Office) getTaxRate() float64 {
	return 0.02 // 2% property tax
}

// Method to calculate the total area of all properties in Portfolio
func (p *Portfolio) totalArea() float64 {
	var total float64
	for _, property := range p.properties {
		total += property.area()
	}
	return total
}

// Method to calculate the total price of all properties
func (p *Portfolio) totalPrice() float64 {
	var total float64
	for _, property := range p.properties {
		total += property.area() * property.getPrice()
	}
	return total
}

// Method to calculate the total tax of all properties
func (p *Portfolio) totalTax() float64 {
	var total float64
	for _, property := range p.properties {
		total += property.area() * property.getPrice() * property.getTaxRate()
	}
	return total
}

// Function to display property details
func (p *Portfolio) displayProperties() {
	fmt.Println("\nProperty Details:")
	for _, property := range p.properties {
		fmt.Printf("%s: %.2f sqft, Price per sqft: $%.2f, Total Price: $%.2f, Tax: $%.2f\n",
			property.getType(),
			property.area(),
			property.getPrice(),
			property.area()*property.getPrice(),
			property.area()*property.getPrice()*property.getTaxRate(),
		)
	}
}

func main() {
	// Create different properties
	house := House{width: 40, length: 30}                 // House with 40x30 dimensions
	apartment := Apartment{numRooms: 5, avgRoomSize: 250} // Apartment with 5 rooms, 250 sqft each
	office := Office{width: 50, length: 60, floors: 3}    // Office with 3 floors of 50x60 sqft each 

	// Create a Portfolio and add properties
	portfolio := Portfolio{
		properties: []Property{house, apartment, office},
	}

	// Display details of each property
	portfolio.displayProperties()

	// Print the total area, price, and tax of all properties with proper formatting
	fmt.Println("\nTotal Managed Property Area:", portfolio.totalArea(), "sqft")
	fmt.Printf("Total Portfolio Price: $%.2f\n", portfolio.totalPrice())  // Formats as a normal decimal
	fmt.Printf("Total Portfolio Tax: $%.2f\n", portfolio.totalTax())      // Formats tax properly
}

/*
✅ Let’s walk through the program:

1️⃣ **Defining an Interface**
   - `Property` is an **interface** that requires any type that implements it to have:
     - `area()` → Calculates the total area of the property.
     - `getType()` → Returns the type of property.
     - `getPrice()` → Returns the price per square foot.
     - `getTaxRate()` → Returns the applicable tax rate.

2️⃣ **Defining Structs**
   - `House` struct stores `width` and `length` values.
   - `Apartment` struct stores `numRooms` (number of rooms) and `avgRoomSize` (size of each room).
   - `Office` struct stores `width`, `length`, and `floors` (number of floors).
   - `Portfolio` struct stores **multiple Property objects**.

3️⃣ **Implementing the `Property` Interface**
   - Each property type (`House`, `Apartment`, `Office`) **implements `Property`**.
   - `area()` calculates the total space for each property.
   - `getPrice()` defines the **cost per square foot** for each property type.
   - `getTaxRate()` sets different **tax rates** for each property type.

4️⃣ **Aggregating Data in `Portfolio`**
   - `totalArea()` → Sums up the area of all properties.
   - `totalPrice()` → Sums up the total property values.
   - `totalTax()` → Sums up the property tax for all properties.
   - `displayProperties()` → Prints detailed information about each property.

5️⃣ **Execution Flow**
   - `house := House{width: 40, length: 30}` → Creates a house.
   - `apartment := Apartment{numRooms: 5, avgRoomSize: 250}` → Creates an apartment.
   - `office := Office{width: 50, length: 60, floors: 3}` → Creates an office.
   - `portfolio := Portfolio{properties: []Property{house, apartment, office}}` → Adds all properties to the portfolio.
   - Calls `displayProperties()`, `totalArea()`, `totalPrice()`, and `totalTax()` to **output property details**.

6️⃣ **Expected Output**
	Property Details:
	House: 1200.00 sqft, Price per sqft: $200.00, Total Price: $240000.00, Tax: $2400.00 
	Apartment: 1250.00 sqft, Price per sqft: $100.00, Total Price: $125000.00,Tax: $1875.00 
	Office: 9000.00 sqft, Price per sqft: $250.00, Total Price: $2250000.00, Tax: $45000.00

	Total Managed Property Area: 11450.00 sqft
	Total Portfolio Price: $ 2615000.00 
	Total Portfolio Tax: $ 49275.00

7️⃣ **Why Use This Approach?**
	- ✅ **Encapsulates multiple property types into a single system**.
	- ✅ **Easily expandable** – We can add more property types without changing `Portfolio` logic.
	- ✅ **Promotes clean and modular design** using interfaces.	
*/