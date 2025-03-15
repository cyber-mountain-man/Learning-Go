package main

import (
	"bufio" // Package for buffered input handling
	"fmt"   // Package for formatted I/O
	"os"    // Package for OS-related functions, including Stdin for input
)

func main() {
	// Prompt user for input
	fmt.Print("Enter some text: ")

	// Create a new Scanner to read input from the standard input (keyboard)
	scanner := bufio.NewScanner(os.Stdin)

	// Read user input
	if scanner.Scan() {
		// Print the entered text
		fmt.Println("You entered:", scanner.Text())
	}

	// Check for errors during scanning
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading input:", err)
	}
}

/*
✅ Let's walk through the program:

1️⃣ **Importing Packages (`import "bufio", "fmt", "os"`)**
   - **`bufio`** → Provides buffered I/O operations for efficient input reading.
   - **`fmt`** → Handles formatted input/output functions like `Println()`.
   - **`os`** → Used for interacting with the operating system (`os.Stdin` for input).

2️⃣ **Buffered Input with `bufio.Scanner`**
   - `bufio.NewScanner(os.Stdin)` creates a scanner to **read user input line by line**.
   - `scanner.Scan()` waits for the user to **enter text** and press **Enter**.
   - `scanner.Text()` retrieves the input as a **string**.

3️⃣ **Error Handling (`scanner.Err()`)**
   - Checks if an error occurred **while scanning user input**.
   - If an error occurs, it prints `"Error reading input"`.

4️⃣ **Expected Output (Terminal Interaction)**
	Enter some text: Hello, Go! You entered: Hello, Go!


5️⃣ **Real-World Applications**
- ✅ **Command-line tools** → Accept user commands.
- ✅ **Interactive scripts** → Get user input dynamically.
- ✅ **Processing logs from `stdin`** → Automate data processing.
*/
