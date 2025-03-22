package main

import (
	"fmt"
	"io"
	"strings"
)

func main() {
	// Create a Reader from a string
	reader := strings.NewReader("Go is awesome!!!") // Treats the string as a readable stream

	// Create a byte slice (buffer) to store read data
	buf := make([]byte, 5) // Buffer size is 5 bytes

	// Read data into the byte slice
	n, err := reader.Read(buf)
	if err != nil && err != io.EOF { // Handle errors (except for End-Of-File)
		fmt.Println("Error reading", err)
	}

	// Print the bytes that were read
	fmt.Println("Read bytes:", string(buf[:n])) // Output: Go is
}

/*
✅ Let's walk through the program:

1️⃣ **Using `strings.NewReader()`**  
   - `strings.NewReader("Go is awesome!!!")` creates **a Reader from a string**.
   - Unlike `bytes.Buffer`, **this is read-only**, meaning you can't modify it.

2️⃣ **Setting Up a Buffer for Reading**  
   - `buf := make([]byte, 5)` creates a **byte slice** to store **5 bytes at a time**.

3️⃣ **Reading Data from the Reader**  
   - `reader.Read(buf)` reads **5 bytes** into `buf` and returns how many bytes were read (`n`).

4️⃣ **Handling Errors**  
   - If an error occurs, it is checked (`err != io.EOF` ensures we ignore the "End-Of-File" error).

5️⃣ **Printing the Read Data**  
   - `string(buf[:n])` converts the **read bytes into a string** before printing.

6️⃣ **Expected Output:**  
		Read bytes: Go is

- Since we **read only 5 bytes**, `"Go is awesome!!!"` gets **truncated to "Go is"**.

7️⃣ **Why Use `strings.NewReader()`?**  
- ✅ **More efficient than `bytes.Buffer` for read-only string data.**  
- ✅ **Works with any function expecting an `io.Reader`.**  
- ✅ **Perfect for streaming data in chunks instead of loading it all at once.**  

8️⃣ **Real-World Use Cases**  
- **Processing text data line by line** (e.g., logs, configuration files).  
- **Streaming large JSON files** without loading everything into memory.  
- **Sending chunks of text over a network connection.**  
*/
