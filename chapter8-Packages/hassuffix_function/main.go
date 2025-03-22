package main

import (
	"fmt"
	"strings"
)

func main() {
	// func HasSuffix(s, suffix string) bool
	// Checks if the string "test" ends with the suffix "st"
	fmt.Println(strings.HasSuffix("test", "st")) // Expected output: true
}

/*
✅ Let's walk through the program:

1️⃣ **Importing Packages**  
   - `fmt` → Used for printing output to the console.  
   - `strings` → Provides string manipulation functions.

2️⃣ **Using `strings.HasSuffix()`**  
   - The function `strings.HasSuffix(s, suffix)` checks if `s` ends with `suffix`.  
   - In this case, `"test"` **ends with** `"st"`, so it returns `true`.

3️⃣ **Expected Output**  
		true

4️⃣ **Real-World Use Cases**  
- ✅ Checking if a filename ends with a specific extension (`.txt`, `.pdf`).  
- ✅ Validating domain suffixes (e.g., `"example.com"`, `"example.org"`).  
- ✅ Ensuring a sentence ends with punctuation (`".", "!"`).  
*/
