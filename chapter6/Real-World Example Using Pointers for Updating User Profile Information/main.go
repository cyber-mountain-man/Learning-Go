/*
✅ Real-World Example: Using Pointers for Updating User Profile Information
📌 Scenario:
You are building a user profile system where users can update their email address.
Since we need to modify the original user data, we use pointers.

📌 Concepts Used from Chapter 6:
✅ Pass-by-Value vs. Pass-by-Pointer – Understanding when data changes persist.
✅ Structs – Representing a user profile.
✅ Pointers (*T) – Modifying original values instead of copies.
*/

// Updating a User's Email Using Pointers

package main

import (
	"fmt"
)

// Define a User struct with a name and email field
type User struct{
	Name string
	Email string
}

// Function to update the email of a user (pass-by-pointer)
func updateEmail(userPtr *User, newEmail string) {
	userPtr.Email = newEmail // Modify the actual User struct in memory
}

func main() {
	// Create a user profile
	user := User{Name: "Alice", Email: "alice@example.com"}

	// Display original profile
	fmt.Println("Before update:", user)

	// Update the email using a pointer function
	updateEmail(&user, "alice.new@example.com")

	// Display update profile
	fmt.Println("After update:", user)
}