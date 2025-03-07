package main

import (
	"fmt"
	"strings"
)

func main() {
	// Contains(s, substr string) bool
	// Checks if the string "test" contains the substring "es"
	fmt.Println(strings.Contains("test", "es")) // Expected output: true
}

/*
✅ Let's walk through the program:

1️⃣ **Importing Packages**  
   - `fmt` → Used for printing output to the console.  
   - `strings` → Provides string manipulation functions.

2️⃣ **Using `strings.Contains()`**
   - The function `strings.Contains(s, substr)` checks if `substr` exists within `s`.  
   - In this case, `"test"` contains `"es"`, so the function returns `true`.

3️⃣ **Expected Output**  
	  	true

4️⃣ **Real-World Use Cases**  
- ✅ Checking if a word appears in user input.  
- ✅ Searching for specific substrings in a dataset.  
- ✅ Validating email domains (e.g., `strings.Contains(email, "@")`).  
*/