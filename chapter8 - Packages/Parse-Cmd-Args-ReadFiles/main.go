package main

import (
	"flag"
	"fmt"
	"os"
)

// ✅ main() - Parses flags and prints contents of each provided file
func main() {
	// ✅ Define a flag to enable verbose output
	verbose := flag.Bool("v", false, "Enable verbose output")

	// ✅ Parse CLI flags (e.g., -v)
	flag.Parse()

	// ✅ Retrieve remaining non-flag arguments (e.g., file names)
	files := flag.Args()

	if len(files) == 0 {
		fmt.Println("❗ No files provided. Usage: go run main.go [-v] file1.txt file2.txt ...")
		os.Exit(1)
	}

	// ✅ Iterate over each file and read its content
	for _, file := range files {
		if *verbose {
			fmt.Printf("📂 Processing file: %s\n", file)
		}

		// ✅ Open the file
		data, err := os.ReadFile(file)
		if err != nil {
			fmt.Printf("❌ Error reading %s: %v\n", file, err)
			continue
		}

		// ✅ Print file contents
		fmt.Printf("📄 Content of %s:\n%s\n", file, data)
	}
}

/*
✅ Let's walk through the program:

1️⃣ **Flag Definition**
   - `verbose := flag.Bool("v", false, ...)` defines a `-v` flag for verbose output.
   - You can toggle additional info like which file is currently being processed.

2️⃣ **Parsing CLI Arguments**
   - `flag.Parse()` processes all flags like `-v`.
   - `flag.Args()` retrieves any remaining positional arguments like `file1.txt`.

3️⃣ **Handling Missing Arguments**
   - If no files are passed, it prints a helpful message and exits.

4️⃣ **Reading and Printing File Contents**
   - Uses `os.ReadFile(file)` to load each file's content.
   - If verbose mode is on, prints the name of each file being processed.
   - Then it prints the **entire content** of each file.

5️⃣ **Expected Terminal Command**
   ```bash
   go run main.go -v file1.txt file2.txt

6️⃣ Sample Output

	📂 Processing file: file1.txt
	📄 Content of file1.txt:
	This is file 1.
	It contains sample content...

	📂 Processing file: file2.txt
	📄 Content of file2.txt:
	Hello from file 2!
...

7️⃣ Real-World Use Cases

	✅ Reading log files, config files, or input data.
	✅ Building tools to process multiple input files (e.g., reports).
	✅ CLI tools that support additional flags and options. */
