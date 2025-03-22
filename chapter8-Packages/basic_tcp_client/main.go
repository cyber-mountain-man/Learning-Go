package main

import (
	"bufio" // ✅ For reading user input from the console
	"fmt"   // ✅ For printing output
	"net"   // ✅ For TCP networking
	"os"    // ✅ For handling OS-related tasks like reading input
)

func main() {
	// ✅ Connect to the TCP server
	conn, err := net.Dial("tcp", "127.0.0.1:8888") // Connects to the server at 127.0.0.1:8888
	if err != nil {
		fmt.Println("❌ Connection failed:", err)
		return
	}
	defer conn.Close() // ✅ Ensures the connection is closed when the program exits

	fmt.Println("✅ Connected to server!")
	scanner := bufio.NewScanner(os.Stdin) // ✅ Reads user input from the console

	for {
		// 📝 Prompt user for input
		fmt.Print("Enter message (or 'exit' to quit): ")
		scanner.Scan()
		message := scanner.Text()

		// 🚪 Exit condition
		if message == "exit" {
			fmt.Println("🚪 Closing client...")
			break
		}

		// 📤 Send message to the server
		_, err = conn.Write([]byte(message + "\n")) // Convert string to byte slice and send
		if err != nil {
			fmt.Println("❌ Error sending message:", err)
			break
		}

		// 📩 Receive response from server
		reply := make([]byte, 1024) // ✅ Create a buffer to store server response
		n, err := conn.Read(reply)  // ✅ Read data from the server
		if err != nil {
			fmt.Println("❌ Error receiving response:", err)
			break
		}

		// ✅ Print the response from the server
		fmt.Println("📩 Received from server:", string(reply[:n]))
	}
}

/*
✅ **Walkthrough of the Program:**

1️⃣ **Connecting to the Server**
   - Uses `net.Dial("tcp", "127.0.0.1:8888")` to **connect to a running TCP server**.
   - If the connection fails, it prints an error message and exits.
   - `defer conn.Close()` ensures that the connection is properly closed after use.

2️⃣ **Reading and Sending User Input**
   - Uses `bufio.NewScanner(os.Stdin)` to **capture user input**.
   - A loop continuously prompts the user to **enter messages**.
   - If the user enters `"exit"`, the program **terminates gracefully**.
   - Otherwise, the message is sent to the server using `conn.Write()`.

3️⃣ **Receiving and Displaying the Server's Response**
   - A **buffer** (`make([]byte, 1024)`) is created to store the server's response.
   - The program waits for a reply using `conn.Read(reply)`.
   - The received message is printed to the console.

4️⃣ **Expected Output**
   - After connecting, it will look like this:
     ```
     ✅ Connected to server!
     Enter message (or 'exit' to quit): Hello, Server!
     📩 Received from server: HELLO, SERVER!  // Example response if server echoes in uppercase
     Enter message (or 'exit' to quit):
     ```

5️⃣ **Why Use This Approach?**
   - ✅ **Real-time communication** between client and server.
   - ✅ **Handles user input dynamically** instead of hardcoded messages.
   - ✅ **Graceful exit mechanism** to close the connection properly.

📌 **Next Steps:**
- ✅ Modify the server to respond with **custom messages**.
- ✅ Implement **authentication** before sending messages.
- ✅ Upgrade to **TLS-encrypted TCP communication** for security.
*/
