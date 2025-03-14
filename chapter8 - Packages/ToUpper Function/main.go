package main

import (
	"fmt"
	"strings"
)

func main() {
	// func ToUpper(s string) string
	// Converts the string "test" to uppercase
	fmt.Println(strings.ToUpper("test")) // Expected output: "TEST"
}

/*
✅ Let's walk through the program:

1️⃣ **Importing Packages**  
   - `fmt` → Used for printing output to the console.  
   - `strings` → Provides string manipulation functions.

2️⃣ **Using `strings.ToUpper()`**  
   - The function `strings.ToUpper(s)` **converts all characters in `s` to uppercase**.  
   - In this case, `"test"` becomes `"TEST"`.  

3️⃣ **Expected Output**  
	"TEST"

	
4️⃣ **Real-World Use Cases**  
- ✅ **Standardizing text for comparisons** (e.g., checking case-insensitive inputs).  
- ✅ **Formatting user data (capitalizing names, acronyms, etc.)**.  
- ✅ **Pre-processing data for systems that require uppercase input**.  
*/
