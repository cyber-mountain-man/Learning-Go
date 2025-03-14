package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	// Prompt user for input
	fmt.Println("Type something and press Enter:")

	// Copy from standard input (keyboard) to standard output (console)
	_, err := io.Copy(os.Stdout, os.Stdin)
	if err != nil {
		// Handle any potential error while reading/writing
		fmt.Println("Error:", err)
	}
}

/*
✅ Let's walk through the program:

1️⃣ **Prompting the User**  
   - The program prints `"Type something and press Enter:"` using `fmt.Println()`.

2️⃣ **Reading Input (`os.Stdin`) and Writing Output (`os.Stdout`)**  
   - `io.Copy(os.Stdout, os.Stdin)` reads **everything typed** by the user and immediately **prints it back**.
   - `os.Stdin` is a **Reader** (reads user input).
   - `os.Stdout` is a **Writer** (prints output to the console).

3️⃣ **Handling Errors**  
   - If `io.Copy()` encounters an issue, the error is caught and displayed.

4️⃣ **Expected Behavior**
   - Whatever the user **types in** is instantly **printed back**.
   - Example Run:
     ```
     Type something and press Enter:
     Hello, Go!
     Hello, Go!
     ```

5️⃣ **Real-World Use Cases**
   - ✅ **Creating an interactive shell program**.
   - ✅ **Processing user input in a CLI tool**.
   - ✅ **Redirecting input/output in Unix-style pipes** (`cat file.txt | go run main.go`).
*/
