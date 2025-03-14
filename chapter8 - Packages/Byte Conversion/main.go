package main

import (
	"fmt"
)

func main() {
	// Convert a string to a slice of bytes
	arr := []byte("test")
	fmt.Println(arr) // Expected output: [116 101 115 116] (ASCII values)

	// Convert a byte slice back to a string
	str := string([]byte{'t', 'e', 's', 't'})
	fmt.Println(str) // Expected output: "test"
}

/*
✅ Let's walk through the program:

1️⃣ **Converting a String to a Byte Slice**  
   - `[]byte("test")` converts `"test"` into a slice of bytes `[116 101 115 116]`.  
   - Each character is stored as its **ASCII value**.

2️⃣ **Converting a Byte Slice to a String**  
   - `string([]byte{'t', 'e', 's', 't'})` reconstructs the original string `"test"`.

3️⃣ **Expected Output**  
		[116 101 115 116] test

4️⃣ **Real-World Use Cases**  
- ✅ **Working with binary data (e.g., encoding images, files, network packets)**.  
- ✅ **Manipulating individual bytes (e.g., encryption, compression algorithms)**.  
- ✅ **Efficient data handling (Go strings are immutable, but byte slices are mutable)**.  
*/
