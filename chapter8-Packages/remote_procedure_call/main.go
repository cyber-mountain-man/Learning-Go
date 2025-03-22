package main

import (
	"fmt"
	"net"
	"net/rpc"
	"time" // ✅ Added to introduce delay before client starts
)

// ✅ Server struct to hold RPC methods
type Server struct{}

// ✅ Negate method: Takes an integer and returns its negative value
func (srv *Server) Negate(i int64, reply *int64) error {
	*reply = -i
	return nil
}

// ✅ server() - Starts an RPC server
func server() {
	// ✅ Register `Server` type so its methods can be called remotely
	rpc.Register(new(Server))

	// ✅ Listen for incoming TCP connections on port 8888
	ln, err := net.Listen("tcp", ":8888")
	if err != nil {
		fmt.Println("❌ Error starting server:", err)
		return
	}
	defer ln.Close()

	fmt.Println("🚀 RPC Server is running on port 8888...")

	for {
		// ✅ Accept an incoming connection
		c, err := ln.Accept()
		if err != nil {
			continue 
		}

		// ✅ Serve the connection in a goroutine (handles multiple clients)
		go rpc.ServeConn(c)
	}
}

// ✅ client() - Calls the RPC server method
func client() {
	// ✅ Connect to the RPC server
	c, err := rpc.Dial("tcp", "127.0.0.1:8888")
	if err != nil {
		fmt.Println("❌ Error connecting to server:", err)
		return
	}
	defer c.Close() // ✅ Ensure connection is closed when done

	// ✅ Call `Server.Negate` method remotely
	var result int64
	err = c.Call("Server.Negate", int64(999), &result)
	if err != nil {
		fmt.Println("❌ RPC Call Error:", err)
	} else {
		fmt.Println("✅ Server.Negate(999) =", result)
	}
}

func main() {
	// ✅ Start the RPC server in a goroutine
	go server()

	// ✅ Wait briefly to allow the server to start before the client connects
	time.Sleep(1 * time.Second)

	// ✅ Start the RPC client
	go client()

	// ✅ Keep the program running
	var input string
	fmt.Scanln(&input)
}

/*
✅ **Walkthrough of the Program:**

1️⃣ **Setting Up an RPC Server (`server()`)**
   - Registers the `Server` struct as an **RPC service**.
   - Uses `net.Listen("tcp", ":8888")` to **start listening** for incoming connections.
   - Calls `rpc.ServeConn(c)` to **handle remote function calls** from clients.

2️⃣ **Defining an RPC Method (`Negate`)**
   - `Negate(i int64, reply *int64) error` takes an integer `i` and **returns its negative value** in `reply`.
   - The method is exposed as **"Server.Negate"** and can be called by clients.

3️⃣ **Client Calls RPC Server (`client()`)**
   - Connects to the **RPC server** using `rpc.Dial("tcp", "127.0.0.1:8888")`.
   - Calls the **remote `Negate` function** using `c.Call("Server.Negate", int64(999), &result)`.
   - Prints the **response from the server**.

4️⃣ **Handling Concurrent Connections**
   - The **server uses goroutines** (`go rpc.ServeConn(c)`) to **handle multiple clients at the same time**.
   - This allows the server to **process requests in parallel**.

5️⃣ **Expected Output**
		🚀 RPC Server is running on port 8888... 
		✅ Server.Negate(999) = -999

- The client sends `999` to the server, and the server **returns `-999`**.

6️⃣ **Why Use RPC?**
- ✅ **Encapsulates business logic** – The client doesn't need to know how `Negate()` works.
- ✅ **Enables distributed systems** – Clients and servers can run **on different machines**.
- ✅ **Allows seamless function calls** – Just like calling a **local function**, but remotely.

7️⃣ **Real-World Use Cases**
✅ **Microservices communication** – Services talk to each other over RPC.  
✅ **Distributed computing** – Remote servers process **complex calculations**.  
✅ **Database access layers** – APIs interact with databases remotely.  
✅ **Inter-process communication (IPC)** – Local programs communicate efficiently.

📌 **Next Steps:**
- ✅ Implement **multiple RPC functions** (e.g., `Multiply`, `Divide`).
- ✅ Use **JSON-RPC** (`net/rpc/jsonrpc`) for **cross-language compatibility**.
- ✅ Secure the connection using **TLS (`crypto/tls`)** for encrypted communication.
*/
