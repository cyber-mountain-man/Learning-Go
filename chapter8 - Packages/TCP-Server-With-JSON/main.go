package main

import (
	"encoding/json" // ✅ Use JSON for encoding/decoding messages
	"fmt"
	"net" // ✅ Import net package for network communication
)

// ✅ Message struct to define the JSON format
type Message struct {
	Text string `json:"text"` // Defines a JSON key "text"
}

// ✅ server() - Starts a TCP server that listens for client messages
func server() {
	// 🛠 Start listening on TCP port 8888
	ln, err := net.Listen("tcp", ":8888")
	if err != nil {
		fmt.Println("❌ Error starting server:", err)
		return
	}
	defer ln.Close() // ✅ Ensure listener is closed when done
	fmt.Println("🚀 Server is running on port 8888...")

	for {
		// 🔄 Accept an incoming client connection
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("❌ Error accepting connection:", err)
			continue
		}
		fmt.Println("✅ Client connected!")

		// 🧵 Handle client connection concurrently using a goroutine
		go handleServerConnection(conn)
	}
}

// ✅ handleServerConnection() - Reads JSON data from the client
func handleServerConnection(conn net.Conn) {
	defer conn.Close() // ✅ Ensure connection is closed after handling

	// 📨 Read JSON data from the client
	var msg Message
	decoder := json.NewDecoder(conn)
	err := decoder.Decode(&msg) // Decodes incoming JSON into `msg`
	if err != nil {
		fmt.Println("❌ Error decoding JSON:", err)
		return
	}

	// ✅ Print the received message
	fmt.Println("📩 Received:", msg.Text)

	// ✅ Send a response back to the client
	response := Message{Text: "Message received successfully!"}
	encoder := json.NewEncoder(conn)
	err = encoder.Encode(response)
	if err != nil {
		fmt.Println("❌ Error encoding response:", err)
		return
	}

	fmt.Println("📤 Response sent to client!")
}

func main() {
	// 🏁 Start the server in a goroutine (non-blocking)
	go server()

	// 🕒 Keep the program running
	var input string
	fmt.Scanln(&input)
}

/*
✅ **Walkthrough of the Program:**

1️⃣ **Starting a TCP Server (`server()`)**
   - Uses `net.Listen("tcp", ":8888")` to **listen for incoming connections** on port `8888`.
   - Waits for clients and accepts connections using `ln.Accept()`.
   - Each client connection is handled **concurrently** using `go handleServerConnection(conn)`.

2️⃣ **Receiving Messages (`handleServerConnection()`)**
   - Reads data from the client using `json.NewDecoder(conn).Decode(&msg)`.
   - Uses a **struct (`Message`)** to parse JSON into a Go object.
   - Prints the **received message** to the console.
   - **Sends a confirmation response** (`"Message received successfully!"`).

3️⃣ **Running the Server (`main()`)**
   - Starts `server()` **in a separate goroutine** (`go server()`).
   - Uses `fmt.Scanln(&input)` to **keep the program running**.

4️⃣ **Expected Output (Example)**
	🚀 Server is running on port 8888... 
	✅ Client connected! 
	📩 Received: Hello from netcat using JSON! 
	📤 Response sent to client!


---

✅ **Why Use This Approach?**
- ✅ **Encapsulates a simple TCP communication system** – Allows server-client communication over a network.
- ✅ **Easier debugging** – JSON is **human-readable** and easier to test.
- ✅ **Cross-platform support** – Can be used with **Python, Bash, PowerShell, etc.**.
- ✅ **More secure** than `gob` for external clients.
- ✅ **Bi-directional communication** – Now supports **sending responses back** to clients.

📌 **Next Steps:**
- ✅ Modify the client to **send multiple messages** instead of just one.
- ✅ Implement **authentication** to prevent unauthorized access.
- ✅ Add **TLS encryption (`crypto/tls`)** for secure communication.
- ✅ Store messages in a **log file** for auditing and security monitoring.
*/
