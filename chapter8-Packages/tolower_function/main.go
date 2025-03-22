package main

import (
	"fmt"
	"strings"
)

func main() {
	// func ToLower(s string) string
	// Converts the string "TEST" to lowercase
	fmt.Println(strings.ToLower("TEST")) // Expected output: "test"
}

/*
✅ Let's walk through the program:

1️⃣ **Importing Packages**  
   - `fmt` → Used for printing output to the console.  
   - `strings` → Provides string manipulation functions.

2️⃣ **Using `strings.ToLower()`**  
   - The function `strings.ToLower(s)` **converts all characters in `s` to lowercase**.  
   - In this case, `"TEST"` becomes `"test"`.  

3️⃣ **Expected Output**  
	"test"


4️⃣ **Real-World Use Cases**  
- ✅ **Case-insensitive user input handling** (e.g., usernames, emails).  
- ✅ **Normalizing text for database storage or comparison**.  
- ✅ **Pre-processing search terms** (ensuring case does not affect search results).  
*/
