package main

import (
	"fmt"
	"net/http"
	"github.com/gorilla/websocket" // ✅ Import WebSocket library
)

// ✅ Upgrade HTTP connection to WebSocket
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true }, // Allows connections from any origin
}

// ✅ handleWebSocket handles WebSocket connections
func handleWebSocket(w http.ResponseWriter, r *http.Request) {
	// 🛠 Upgrade HTTP request to WebSocket
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("❌ Error upgrading to WebSocket:", err)
		return
	}
	defer conn.Close() // ✅ Ensure connection is closed when done

	fmt.Println("✅ Client connected!")

	for {
		// 📨 Read message from client
		_, msg, err := conn.ReadMessage()
		if err != nil {
			fmt.Println("❌ Client disconnected:", err)
			break
		}

		fmt.Println("📩 Received:", string(msg))

		// 🔁 Echo message back to client
		err = conn.WriteMessage(websocket.TextMessage, msg)
		if err != nil {
			fmt.Println("❌ Error sending message:", err)
			break
		}

		fmt.Println("📤 Message echoed back!")
	}
}

func main() {
	// ✅ Define WebSocket route
	http.HandleFunc("/ws", handleWebSocket)

	// ✅ Start the server
	port := "8080"
	fmt.Println("🚀 WebSocket Server is running on ws://127.0.0.1:" + port + "/ws")
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		fmt.Println("❌ Server error:", err)
	}
}

/*
✅ **Walkthrough of the Program:**

1️⃣ **Upgrading HTTP to WebSocket (`upgrader.Upgrade()`)**
   - Uses `github.com/gorilla/websocket` to **upgrade** an HTTP connection to a WebSocket connection.
   - The `CheckOrigin` function **allows all origins**, making it easier for testing.

2️⃣ **Handling WebSocket Connections (`handleWebSocket()`)**
   - **Accepts a WebSocket connection** from a client.
   - **Listens for messages** from the client using `conn.ReadMessage()`.
   - **Logs received messages** and **echoes them back** using `conn.WriteMessage()`.

3️⃣ **Starting the WebSocket Server (`main()`)**
   - Uses `http.HandleFunc("/ws", handleWebSocket)` to define the **WebSocket endpoint**.
   - Starts the **HTTP server** on port `8080`.
   - Clients can connect via **ws://127.0.0.1:8080/ws**.

4️⃣ **Testing the WebSocket Server**
   - **With `wscat` (Node.js package):**
     ```sh
     npm install -g wscat
     wscat -c ws://127.0.0.1:8080/ws
     ```
   - **With JavaScript in Browser Console:**
     ```js
     let ws = new WebSocket("ws://127.0.0.1:8080/ws");
     ws.onopen = () => ws.send("Hello WebSocket!");
     ws.onmessage = (msg) => console.log("Received:", msg.data);
     ```
   - **With cURL (WebSocket support required):**
     ```sh
     curl --include --no-buffer --header "Connection: Upgrade" --header "Upgrade: websocket" ws://127.0.0.1:8080/ws
     ```

5️⃣ **Expected Output (Example)**
	🚀 WebSocket Server is running on ws://127.0.0.1:8080/ws 
	✅ Client connected! 
	📩 Received: Hello WebSocket! 
	📤 Message echoed back!

---


---

### ✅ **Why Use WebSockets?**
- **Bidirectional Communication** → Clients & servers can send messages **asynchronously**.
- **Efficient for Real-time Apps** → No need to poll for updates like HTTP.
- **Use Cases** → **Chat apps, stock tickers, collaborative tools, real-time notifications.**

📌 **Next Steps:**
- ✅ Implement **multiple clients** and broadcast messages.
- ✅ Secure the WebSocket connection using **TLS (`wss://`)**.
- ✅ Store chat history or messages in a **database**.
*/
