package main

import (
	"fmt"
	"net/http" // ✅ Import the HTTP package for handling web requests
)

// ✅ handler function that responds to all HTTP requests
func handler(w http.ResponseWriter, r *http.Request) {
	// 📝 Log the request sender's IP address and port
	fmt.Println("📩 Request received from:", r.RemoteAddr)

	// ✅ Write a response back to the client (browser, curl, etc.)
	fmt.Fprintln(w, "Hello, World!") // Sends "Hello, World!" as the response
}

func main() {
	// ✅ Assign handler function to the root URL "/"
	http.HandleFunc("/", handler) // Routes all requests to the `handler` function

	// ✅ Define the port and start the server
	port := "8080"
	fmt.Println("🚀 Web Server is running on http://127.0.0.1:" + port)

	// ✅ Start HTTP server and listen on port 8080
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		fmt.Println("❌ Server error:", err) // Logs an error if the server fails
	}
}

/*
✅ **Walkthrough of the Program:**

1️⃣ **Starting an HTTP Server (`http.ListenAndServe()`)**
   - The server listens on **port 8080**.
   - Uses `http.HandleFunc("/", handler)` to **route requests** to the `handler` function.

2️⃣ **Handling Incoming Requests (`handler()`)**
   - Logs the **IP address and port** of the client (`r.RemoteAddr`).
   - Sends a simple `"Hello, World!"` **response** back to the client.

3️⃣ **Running the Server (`main()`)**
   - Starts the web server on **http://127.0.0.1:8080**.
   - **Keeps running** indefinitely to handle multiple client requests.

---
🔍 **How to Test the Server**

**Option 1: Web Browser**
1. Open a browser.
2. Go to **http://127.0.0.1:8080/**.
3. You should see:  Hello, World!


**Option 2: cURL (Command Line)**
Run: curl http://127.0.0.1:8080/
Expected response: Hello, World!
---

✅ Why Use This Approach?
Minimal setup – Uses Go’s built-in net/http package.
Handles concurrent requests – The Go HTTP server is multithreaded by default.
Easy to extend – Can add multiple endpoints, middleware, authentication, and more.
Lightweight & fast – No external dependencies required.

🚀 Next Steps
🔹 Add multiple endpoints (/about, /contact, etc.).
🔹 Implement logging for request details.
🔹 Secure the server using HTTPS (crypto/tls)
*/