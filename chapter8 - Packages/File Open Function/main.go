package main

import (
	"fmt"
	"os"
)

func main() {
	// Open the file "test.txt"
	file, err := os.Open("test.txt")
	if err != nil {
		// Handle the error (e.g., file not found)
		fmt.Println("Error opening file:", err)
		return
	}
	// Ensure the file is closed when the function exits
	defer file.Close()

	// Get the file size
	stat, err := file.Stat()
	if err != nil {
		fmt.Println("Error getting file info:", err)
		return
	}

	// Read the file content
	bs := make([]byte, stat.Size()) // Create a byte slice of the exact file size
	_, err = file.Read(bs)          // Read file data into the byte slice
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Convert byte slice to a string
	str := string(bs)
	fmt.Println(str) // Print file contents
}

/*
✅ Let's walk through the program:

1️⃣ **Opening a File (`os.Open()`)**  
   - The program opens `"test.txt"` using `os.Open("test.txt")`.  
   - If the file does **not exist**, an error is returned and handled.

2️⃣ **Ensuring File Closure (`defer file.Close()`)**  
   - `defer file.Close()` ensures that the file is closed when the function exits, preventing memory leaks.

3️⃣ **Getting File Information (`file.Stat()`)**  
   - `file.Stat()` retrieves file details (size, name, etc.).
   - The **file size** is needed to create a byte slice to hold the data.

4️⃣ **Reading the File (`file.Read(bs)`)**  
   - A byte slice (`bs := make([]byte, stat.Size())`) is created to store the file contents.
   - `file.Read(bs)` reads the file data into `bs`.

5️⃣ **Converting to a String (`string(bs)`)**  
   - The **byte slice is converted to a string** and printed.

6️⃣ **Expected Output (if `test.txt` contains "Hello, World!")**  
		Hello, World!

7️⃣ **Real-World Use Cases**  
- ✅ **Reading configuration files (`.json`, `.yaml`)**.  
- ✅ **Processing log files for analytics**.  
- ✅ **Loading text files for web applications or scripts**.  
*/		
