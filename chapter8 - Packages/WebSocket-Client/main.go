package main

import (
	"bufio"          // ✅ Reads user input from the console
	"fmt"            // ✅ Standard Go package for formatted I/O
	"log"            // ✅ Logs errors and events
	"net/url"        // ✅ Parses WebSocket URL
	"os"             // ✅ Handles system signals (e.g., Ctrl+C)
	"os/signal"      // ✅ Captures OS interrupt signals

	"github.com/gorilla/websocket" // ✅ WebSocket package for real-time communication
)

func main() {
	// ✅ Define the WebSocket server address
	serverAddr := "127.0.0.1:8080" // WebSocket server running locally on port 8080
	u := url.URL{Scheme: "ws", Host: serverAddr, Path: "/ws"} // Constructs the WebSocket URL

	// ✅ Connect to WebSocket server
	fmt.Println("🔗 Connecting to WebSocket server:", u.String())
	conn, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("❌ Connection failed:", err)
	}
	defer conn.Close() // ✅ Ensures connection closes on exit
	fmt.Println("✅ Connected to WebSocket Server!")

	// ✅ Handle incoming messages concurrently (Goroutine)
	done := make(chan struct{}) // Channel to signal when reading is complete
	go func() {
		defer close(done)
		for {
			// 📨 Read messages from the WebSocket server
			_, message, err := conn.ReadMessage()
			if err != nil {
				log.Println("❌ Read error:", err)
				return
			}
			fmt.Println("📩 Received:", string(message)) // ✅ Print received message
		}
	}()

	// ✅ Capture system interrupt signals (Ctrl+C)
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	// ✅ Read user input and send messages
	scanner := bufio.NewScanner(os.Stdin) // Scanner to read input from the console
	fmt.Println("✏️  Type messages to send (or 'exit' to quit):")
	for {
		select {
		case <-interrupt:
			fmt.Println("🚪 Closing connection gracefully...")
			// ✅ Send a close message to the WebSocket server before exiting
			err := conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, "Client exiting"))
			if err != nil {
				log.Println("❌ Close message error:", err)
			}
			return
		default:
			// ✅ Read user input
			if scanner.Scan() {
				msg := scanner.Text()
				if msg == "exit" { // ✅ If user types "exit", close connection properly
					fmt.Println("🚪 Exiting client...")
					err := conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, "Client exited manually"))
					if err != nil {
						log.Println("❌ Close message error:", err)
					}
					return
				}

				// ✅ Send message to WebSocket server
				err := conn.WriteMessage(websocket.TextMessage, []byte(msg))
				if err != nil {
					log.Println("❌ Write error:", err)
					return
				}
				fmt.Println("📤 Sent:", msg) // ✅ Confirm message was sent
			}
		}
	}
}

/*
✅ **Walkthrough of the Program:**

1️⃣ **Connecting to the WebSocket Server (`main()`)**
   - Constructs the **WebSocket URL** (`ws://127.0.0.1:8080/ws`).
   - Uses `websocket.DefaultDialer.Dial()` to connect.
   - **Gracefully closes** the connection using `defer conn.Close()`.

2️⃣ **Handling Incoming Messages (`goroutine`)**
   - Starts a **separate goroutine** to listen for messages.
   - Reads messages using `conn.ReadMessage()`.
   - **Prints received messages** to the console.

3️⃣ **Sending Messages (`user input`)**
   - Reads user input using `bufio.NewScanner(os.Stdin)`.
   - Sends messages using `conn.WriteMessage(websocket.TextMessage, msg)`.
   - **Allows user to type 'exit'** to disconnect.

4️⃣ **Handling Interruptions (`Ctrl+C`)**
   - Uses `os.Signal` to detect `Ctrl+C`.
   - **Sends a CloseMessage** before exiting.

---

✅ **Expected Behavior**
#### **Client (Terminal Output)**
		🔗 Connecting to WebSocket server: ws://127.0.0.1:8080/ws 
		✅ Connected to WebSocket Server! 
		✏️ Type messages to send (or 'exit' to quit): hello server 
		📤 Sent: hello server 
		📩 Received: hello server (echoed from server) ^C 
		🚪 Closing connection gracefully...

#### **Server (Terminal Output)**
		🚀 WebSocket Server is running on ws://127.0.0.1:8080/ws 
		✅ Client connected! 
		📩 Received: hello server 
		📤 Message echoed back! 
		❌ Client disconnected


---

✅ **Why Use This Approach?**
- ✅ **Gracefully handles disconnects** (Ctrl+C support).
- ✅ **Prevents resource leaks** (Ensures connections close properly).
- ✅ **Handles user input dynamically**.
- ✅ **Enables full-duplex communication** (bi-directional messaging).

📌 **Next Steps:**
- ✅ Modify the client to send JSON messages instead of plain text.
- ✅ Reconnect automatically if the connection is lost.
- ✅ Implement authentication (JWT or API key).
- ✅ Secure the WebSocket connection with TLS.
*/ 
