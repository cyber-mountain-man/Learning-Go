package main

import (
	"flag" // ✅ Used for parsing command-line flags
	"fmt"
	"os"
)

func main() {
	// ✅ Define a boolean flag -v for verbose mode
	verbose := flag.Bool("v", false, "Enable verbose output")

	// ✅ Parse the command-line flags and arguments
	flag.Parse()

	// ✅ Retrieve any remaining arguments (non-flag)
	files := flag.Args()

	// ✅ Exit if no file arguments were provided
	if len(files) == 0 {
		fmt.Println("❗ No files provided.")
		os.Exit(1)
	}

	// ✅ Iterate over provided filenames
	for _, file := range files {
		if *verbose {
			// If verbose is enabled, print full message
			fmt.Printf("📂 Processing file: %s\n", file)
		} else {
			// Otherwise, just print the filename
			fmt.Println(file)
		}
	}
}

/*
✅ **Walkthrough of the Program:**

1️⃣ **Flag Definition and Parsing**
   - Uses `flag.Bool()` to define a `-v` (verbose) flag.
   - `flag.Parse()` parses the provided CLI flags.
   - Remaining non-flag arguments are accessed with `flag.Args()`.

2️⃣ **File Input Handling**
   - If no file arguments are passed, it exits with a message.
   - Otherwise, it iterates through each file name.

3️⃣ **Verbose Output Toggle**
   - If `-v` is passed, prints: `📂 Processing file: filename.txt`.
   - If not, it simply prints the filename.

4️⃣ **Expected Example Usage**
   ```bash
   go run main.go file1.txt file2.txt
   go run main.go -v file1.txt file2.txt

5️⃣ Example Output
📂 Processing file: file1.txt
📂 Processing file: file2.txt

✅ Why Use This Pattern?

	✅ Enables flexible user input without hardcoding filenames.
	✅ Good foundation for building CLI tools that support flags and arguments.
	✅ Separates flags from regular arguments in a clean and manageable way.

📌 Next Steps:

	✅ Extend to read and display file contents.
	✅ Add error handling for file access.
	✅ Support multiple flags (e.g., -out, -filter, etc.). */