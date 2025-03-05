/*
✅ Real-World Example: Updating a Price in a Shopping Cart
📌 Scenario:
A user is shopping online, and we need to apply a discount to an item’s price.
Since we want the change to persist, we use a pointer.

📌 Concepts Used from Chapter 6:
✅ Pass-by-Value vs. Pass-by-Pointer – Understanding when changes persist.
✅ Pointers (*int and *float64) – Modifying values directly in memory.
*/

// Applying a Discount to an Item Price


package main

import "fmt"

// Function to apply a discount to an item's price (pass-by-pointer)
func applyDiscount(price *float64, discount float64) {
	*price = *price * (1 - discount/100) // Modify the original price
}

func main() {
	itemPrice := 100.0 // original price in dollars

	fmt.Println("Original Price: $", itemPrice)

	applyDiscount(&itemPrice, 10) // Apply a 10% discount

	fmt.Println("Discounted Price: $", itemPrice) // Price is now updated
}

/*
🔹 Why Use Pointers in This Case?
✅ Without pointers, Go would pass a copy of itemPrice, so changes wouldn’t persist.
✅ With pointers (*float64), applyDiscount() modifies the actual price in memory.
✅ Useful in real-world scenarios like shopping carts, pricing systems, and financial apps.
*/