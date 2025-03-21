package main

import (
	"fmt"
	"net/rpc" // ✅ Import the `net/rpc` package to enable RPC communication
)

func main() {
	// 🔌 Connect to the RPC server running on 127.0.0.1:8888
	client, err := rpc.Dial("tcp", "127.0.0.1:8888") // Establish a connection to the RPC server
	if err != nil {
		fmt.Println("❌ Connection failed:", err)
		return
	}
	defer client.Close() // ✅ Ensure the connection is closed when done

	// 📤 Call the remote method "Server.Negate" with input 999
	var result int64
	err = client.Call("Server.Negate", int64(999), &result) // Call the remote procedure with an integer argument
	if err != nil {
		fmt.Println("❌ RPC Call error:", err)
		return
	}

	// ✅ Print the result received from the RPC server
	fmt.Println("Server.Negate(999) =", result) // Expected output: Server.Negate(999) = -999
}

/*
✅ **Walkthrough of the Program:**

1️⃣ **Establishing a Connection (`rpc.Dial()`)**
   - `rpc.Dial("tcp", "127.0.0.1:8888")` creates a **TCP connection** to the **RPC server**.
   - If the connection **fails**, the program prints an error message and exits.
   - `defer client.Close()` ensures that the connection is properly **closed** after execution.

2️⃣ **Calling a Remote Procedure (`client.Call()`)**
   - The client calls **`Server.Negate`** on the **remote RPC server**.
   - The method takes an **integer argument (999)** and returns its **negated value (-999)**.
   - The result is stored in the `result` variable.
   - If the RPC call fails, an error message is displayed.

3️⃣ **Printing the Response**
   - The result of `Server.Negate(999)` is **printed to the console**.

4️⃣ **Expected Output**
	Server.Negate(999) = -999


---

### ✅ **Why Use RPC (Remote Procedure Call)?**
- **Allows executing functions on a remote machine** as if they were local.
- **Encapsulates function calls over the network**, reducing complexity.
- **Useful for microservices, distributed computing, and inter-process communication**.

---

### 📌 **Next Steps**
- ✅ Extend `Server` with more **remote functions** (e.g., addition, multiplication).
- ✅ Implement **error handling** to handle invalid inputs.
- ✅ Secure the RPC calls using **TLS encryption** (`crypto/tls`).
- ✅ Explore **`net/rpc/jsonrpc`** to enable **JSON-based RPC communication**.

This program demonstrates a **basic RPC client** in Go! 🚀
*/
