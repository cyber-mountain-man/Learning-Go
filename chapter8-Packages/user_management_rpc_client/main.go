package main

import (
	"fmt"
	"net/rpc"
)

// ✅ User struct to match the server's User struct
type User struct {
	ID    int
	Name  string
	Email string
}

func main() {
	// 🔌 Connect to the RPC server running on 127.0.0.1:8080
	client, err := rpc.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		fmt.Println("❌ Connection failed:", err)
		return
	}
	defer client.Close()

	// ✅ Create a new user
	newUser := User{ID: 1, Name: "John Doe", Email: "john@example.com"}
	var createReply string
	err = client.Call("UserService.CreateUser", newUser, &createReply)
	if err != nil {
		fmt.Println("❌ Error creating user:", err)
	} else {
		fmt.Println("✅", createReply)
	}

	// ✅ Retrieve the user
	var retrievedUser User
	err = client.Call("UserService.GetUser", newUser.ID, &retrievedUser)
	if err != nil {
		fmt.Println("❌ Error retrieving user:", err)
	} else {
		fmt.Printf("📩 Retrieved User: %+v\n", retrievedUser)
	}

	// ✅ Delete the user
	var deleteReply string
	err = client.Call("UserService.DeleteUser", newUser.ID, &deleteReply)
	if err != nil {
		fmt.Println("❌ Error deleting user:", err)
	} else {
		fmt.Println("✅", deleteReply)
	}

	// ✅ Try retrieving the user again after deletion
	err = client.Call("UserService.GetUser", newUser.ID, &retrievedUser)
	if err != nil {
		fmt.Println("❌ Expected error: User not found")
	} else {
		fmt.Printf("❌ Unexpected: Retrieved User: %+v\n", retrievedUser)
	}
}

/*
✅ **Walkthrough of the Program:**

1️⃣ **Connecting to the RPC Server (`rpc.Dial()`)**
   - The client establishes a TCP connection to `127.0.0.1:8080` where the RPC server is running.

2️⃣ **Creating a User (`CreateUser`)**
   - Calls `UserService.CreateUser` remotely to **add a new user**.
   - Expects a **success message** confirming user creation.

3️⃣ **Retrieving a User (`GetUser`)**
   - Calls `UserService.GetUser` remotely to **fetch user details**.
   - Prints the retrieved user data.

4️⃣ **Deleting a User (`DeleteUser`)**
   - Calls `UserService.DeleteUser` remotely to **remove a user by ID**.
   - Expects a **success message** confirming deletion.

5️⃣ **Attempting to Retrieve a Deleted User**
   - Calls `UserService.GetUser` again to check if the user was truly removed.
   - **Expected output:** `"User not found"` error.

6️⃣ **Expected Output**
	✅ User John Doe created successfully!  
	📩 Retrieved User: {ID:1 Name:John Doe Email:john@example.com}  
	✅ User with ID 1 deleted successfully!  
	❌ Expected error: User not found  

7️⃣ **Why Use This Approach?**
   - ✅ **RPC enables structured function calls across networks.**
   - ✅ **Ideal for microservices** where services need to interact.
   - ✅ **Efficient for backend-to-backend communication** without exposing REST APIs.

📌 **Next Steps:**
- ✅ Implement **authentication** for secure RPC calls.
- ✅ Store users in a **database** instead of in-memory storage.
- ✅ Convert to **RPC over HTTP** for better cross-platform integration.
*/
