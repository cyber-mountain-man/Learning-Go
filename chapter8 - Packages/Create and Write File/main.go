package main

import (
	"fmt"
	"os"
)

func main() {
	// Step 1: Create a new file (or overwrite if it already exists)
	file, err := os.Create("output.txt") // Creates "output.txt" in the current directory
	if err != nil {
		fmt.Println("Error creating file:", err) // Handle file creation error
		return
	}
	// Ensure the file is closed after the function exits
	defer file.Close()

	// Step 2: Write some text to the file
	_, err = file.WriteString("This is a test file written in Go!\n") // Write a string to the file
	if err != nil {
		fmt.Println("Error writing to file:", err) // Handle file write error
		return
	}

	// Step 3: Confirm successful file writing
	fmt.Println("File successfully written!") // Prints success message
}

/*
✅ Let's walk through the program:

1️⃣ **Creating a File (`os.Create()`)**
   - The program calls `os.Create("output.txt")` to create a new file.
   - If the file **already exists**, it **overwrites** it.
   - If an error occurs (e.g., **no write permission**), it is handled gracefully.

2️⃣ **Ensuring Proper File Closure (`defer file.Close()`)**
   - `defer` ensures that the file is **automatically closed** when the function exits.
   - This **prevents memory leaks** and avoids potential errors.

3️⃣ **Writing to the File (`file.WriteString()`)**
   - The program writes `"This is a test file written in Go!"` to `output.txt`.
   - If writing fails (e.g., **disk full, permission denied**), an error is handled.

4️⃣ **Output Confirmation**
   - If everything runs successfully, **"File successfully written!"** is printed.

5️⃣ **Expected Output (Terminal)**
	   File successfully written!
	- A new file `output.txt` is created, containing:
  ```
  This is a test file written in Go!
  ```

6️⃣ **Real-World Applications**
- ✅ **Generating reports/logs dynamically**
- ✅ **Saving user data or configuration settings**
- ✅ **Writing logs for debugging or audit trails**
- ✅ **Automating text-based data storage**
*/
