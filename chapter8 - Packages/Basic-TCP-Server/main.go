package main

import (
	"fmt"
	"net" // ✅ Import the `net` package for networking
	"bufio" // ✅ Allows reading incoming text line-by-line
	"strings" // ✅ For string manipulation
)

// ✅ server() - Starts a basic TCP server that listens for text-based messages
func server() {
	// 🛠 Start listening on TCP port 8888
	ln, err := net.Listen("tcp", ":8888")
	if err != nil {
		fmt.Println("❌ Error starting server:", err)
		return
	}
	defer ln.Close() // ✅ Ensure listener is closed when done
	fmt.Println("🚀 Basic TCP Server is running on port 8888...")

	for {
		// 🔄 Accept an incoming client connection
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("❌ Error accepting connection:", err)
			continue
		}
		fmt.Println("✅ Client connected!")

		// 🧵 Handle the client connection concurrently
		go handleServerConnection(conn)
	}
}

// ✅ handleServerConnection() - Reads text data from the client
func handleServerConnection(conn net.Conn) {
	defer conn.Close() // ✅ Ensure the connection is closed after handling

	// 🔍 Read the incoming text message
	reader := bufio.NewReader(conn)
	for {
		message, err := reader.ReadString('\n') // Read until newline
		if err != nil {
			fmt.Println("❌ Connection closed by client.")
			break
		}

		// 📨 Print the received message
		fmt.Print("📩 Received: ", message)

		// 🔄 Echo the message back to the client
		response := strings.ToUpper(message) // Convert to uppercase for testing
		conn.Write([]byte(response)) // Send back the response
		fmt.Println("📤 Response sent!")
	}
}

// ✅ main() - Starts the TCP server
func main() {
	// 🏁 Start the server in a goroutine (non-blocking)
	go server()

	// 🕒 Keep the program running
	var input string
	fmt.Scanln(&input)
}

/* 📌 Walkthrough of the Program:

1️⃣ Starting a TCP Server (server())

		Uses net.Listen("tcp", ":8888") to listen for incoming connections on port 8888.
		Waits for clients and accepts connections using ln.Accept().
		Each client connection is handled concurrently using go handleServerConnection(conn).

2️⃣ Receiving Messages (handleServerConnection())

		Reads raw text-based messages from the client using bufio.Reader.
		Prints the received message.
		Echoes the message back to the client (converted to uppercase).

3️⃣ Running the Server (main())

		Starts server() in a separate goroutine (go server()).
		Uses fmt.Scanln(&input) to keep the program running.

4️⃣ Expected Output (Example)
		🚀 Basic TCP Server is running on port 8888...
		✅ Client connected!
		📩 Received: hello server
		📤 Response sent!


✅ Why Use This Approach?
	✅ Basic implementation – Simple text-based communication.
	✅ Lightweight – No need for complex parsing (like JSON).
	✅ Good for learning – Helps understand TCP server fundamentals.
	✅ Easily extendable – You can modify it for logging, authentication, or more.
	
📌 Next Steps:

	✅ Modify the client to send multiple messages instead of just one.
	✅ Implement a simple command system (e.g., "EXIT" to close the connection).
	✅ Add logging to store messages received from clients.
	✅ Convert the program into a chat server where multiple clients can talk to each other.
*/