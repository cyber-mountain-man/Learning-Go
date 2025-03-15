package main

import (
	"fmt"
	"os"
)

func main() {
	// Step 1: Open the file in append mode
	// - "os.O_APPEND" -> Opens the file for appending (keeps existing content)
	// - "os.O_WRONLY" -> Allows only writing, not reading
	// - "0644" -> Sets the file permissions (owner can read/write, others can read)
	file, err := os.OpenFile("output.txt", os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error opening file:", err) // Handle file open error
		return
	}
	// Ensure the file is closed when function exits
	defer file.Close()

	// Step 2: Append new content to the file
	_, err = file.WriteString("Appending new content!\n") // Adds a new line at the end of the file
	if err != nil {
		fmt.Println("Error writing to file:", err) // Handle file write error
		return
	}

	// Step 3: Print success message
	fmt.Println("Successfully appended data!")
}

/*
✅ Let's walk through the program:

1️⃣ **Opening the File in Append Mode (`os.OpenFile()`)**
   - Unlike `os.Create()`, which overwrites the file, `os.OpenFile()` with `os.O_APPEND` **keeps existing content**.
   - `os.O_WRONLY` allows writing but prevents reading.
   - **Permissions (`0644`)**:
     - Owner can read/write.
     - Others can only read.

2️⃣ **Appending Data (`file.WriteString()`)**
   - Appends `"Appending new content!\n"` to `output.txt` without erasing previous content.
   - If the file **does not exist**, this program will fail.

3️⃣ **Ensuring Proper File Closure (`defer file.Close()`)**
   - **Prevents resource leaks** by ensuring the file closes after writing is complete.

4️⃣ **Output Confirmation**
   - If the program runs successfully, it prints:
     ```
     Successfully appended data!
     ```

5️⃣ **Expected File Contents (output.txt)**
	This is a test file written in Go! Appending new content!


6️⃣ **Real-World Applications**
- ✅ **Logging System:** Continuously appends logs instead of overwriting them.
- ✅ **Transaction Records:** Append new transactions without modifying old ones.
- ✅ **Data Tracking:** Keep track of user activity by appending new entries.
*/