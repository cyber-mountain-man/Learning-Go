package main

import (
	"fmt"
	"net"
	"net/rpc"
)

// ✅ User struct represents a user in the system
type User struct {
	ID    int    // Unique identifier for the user
	Name  string // User's name
	Email string // User's email address
}

// ✅ UserService provides RPC methods for user management
type UserService struct {
	users map[int]User // Stores users in-memory with ID as the key
}

// ✅ CreateUser adds a new user
func (u *UserService) CreateUser(user User, reply *string) error {
	u.users[user.ID] = user // Store user in the map
	*reply = fmt.Sprintf("User %s created successfully!", user.Name)
	return nil
}

// ✅ GetUser retrieves a user by ID
func (u *UserService) GetUser(id int, reply *User) error {
	user, exists := u.users[id] // Check if user exists
	if !exists {
		return fmt.Errorf("User with ID %d not found", id)
	}
	*reply = user // Return the user details
	return nil
}

// ✅ DeleteUser removes a user by ID
func (u *UserService) DeleteUser(id int, reply *string) error {
	_, exists := u.users[id] // Check if user exists
	if !exists {
		return fmt.Errorf("User with ID %d not found", id)
	}
	delete(u.users, id) // Remove user from the map
	*reply = fmt.Sprintf("User with ID %d deleted successfully!", id)
	return nil
}

func main() {
	// ✅ Initialize UserService with an empty user map
	userService := &UserService{users: make(map[int]User)}

	// ✅ Register the RPC service to allow remote method calls
	err := rpc.Register(userService)
	if err != nil {
		fmt.Println("❌ Error registering RPC service:", err)
		return
	}

	// ✅ Start a TCP listener on port 8080
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("❌ Error starting server:", err)
		return
	}
	defer listener.Close()

	fmt.Println("🚀 RPC Server for User Management is running on tcp://127.0.0.1:8080")

	// ✅ Accept incoming client connections
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("❌ Connection error:", err)
			continue
		}
		// ✅ Serve the connection in a separate goroutine for concurrency
		go rpc.ServeConn(conn)
	}
}

/*
✅ **Walkthrough of the Program:**

1️⃣ **Defining a User Struct (`User`)**
   - Represents a user with `ID`, `Name`, and `Email`.

2️⃣ **Creating UserService with Three Methods**
   - `CreateUser(user User, reply *string)` → Adds a user.
   - `GetUser(id int, reply *User)` → Retrieves a user by ID.
   - `DeleteUser(id int, reply *string)` → Deletes a user by ID.

3️⃣ **Registering the RPC Service (`rpc.Register()`)**
   - Registers the `UserService` so that RPC clients can **call its methods remotely**.

4️⃣ **Starting the TCP Listener (`net.Listen("tcp", ":8080")`)**
   - Starts a **server** listening on port `8080`.

5️⃣ **Accepting and Serving Clients (`listener.Accept()`)**
   - Listens for client connections.
   - Uses `go rpc.ServeConn(conn)` to **handle each connection in a separate goroutine** (concurrent execution).

6️⃣ **Expected Usage**
   - A **client connects to the server** and can remotely call:
     - `CreateUser()` to add users.
     - `GetUser()` to retrieve user details.
     - `DeleteUser()` to remove users.

7️⃣ **Why Use This Approach?**
   - ✅ **RPC enables structured remote function calls** over a network.
   - ✅ **Works for distributed systems** where microservices communicate.
   - ✅ **Efficient for backend-to-backend communication**.

📌 **Next Steps:**
- ✅ Implement a **Go-based RPC client** to interact with this server.
- ✅ Store users in a **database instead of an in-memory map**.
- ✅ Use **RPC over HTTP** instead of TCP for better compatibility.
*/
