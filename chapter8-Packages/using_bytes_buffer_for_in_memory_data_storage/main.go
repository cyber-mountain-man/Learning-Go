package main

import (
	"bytes"
	"fmt"
)

func main() {
	// Declare a Buffer (no need to manually initialize)
	var buf bytes.Buffer 

	// Write a string to the buffer
	buf.Write([]byte("Hello, Buffer!")) // Stores the string in memory as bytes

	// Convert buffer to a string and print it
	fmt.Println(buf.String()) // Output: Hello, Buffer!
}

/*
✅ Let's walk through the program:

1️⃣ **What is `bytes.Buffer`?**  
   - `bytes.Buffer` is an in-memory buffer that **stores byte data**.
   - It can be used as both **a Reader and a Writer**, making it flexible for handling data.

2️⃣ **Declaring a Buffer**  
   - `var buf bytes.Buffer` → Declares a buffer without initialization.  
   - Go **automatically initializes it**, unlike slices or arrays.

3️⃣ **Writing Data to the Buffer**  
   - `buf.Write([]byte("Hello, Buffer!"))` stores `"Hello, Buffer!"` inside the buffer as bytes.

4️⃣ **Reading Data from the Buffer**  
   - `buf.String()` converts the **byte slice to a string** and prints it.

5️⃣ **Expected Output:**  
		Hello, Buffer!

		
6️⃣ **Why Use `bytes.Buffer`?**  
- ✅ **Faster than writing to a file for temporary data storage.**  
- ✅ **Doesn’t require manual memory allocation.**  
- ✅ **Useful for processing strings before sending them to files, networks, or logs.**  

7️⃣ **Real-World Use Cases**  
- **Efficiently manipulating strings in memory** before writing them elsewhere.  
- **Building dynamic data before writing to a file.**  
- **Buffering API responses before sending to clients.**  
*/
