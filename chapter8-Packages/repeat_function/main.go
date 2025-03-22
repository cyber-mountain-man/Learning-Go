package main

import (
	"fmt"
	"strings"
)

func main() {
	// func Repeat(s string, count int) string
	// Repeats the string "a" 5 times
	fmt.Println(strings.Repeat("a", 5)) // Expected output: "aaaaa"
}

/*
✅ Let's walk through the program:

1️⃣ **Importing Packages**  
   - `fmt` → Used for printing output to the console.  
   - `strings` → Provides string manipulation functions.

2️⃣ **Using `strings.Repeat()`**  
   - The function `strings.Repeat(s, count)` **creates a new string** by repeating `s` `count` times.  
   - In this case, `"a"` is repeated **5 times**, resulting in `"aaaaa"`.

3️⃣ **Expected Output**  
		"aaaaa"

4️⃣ **Real-World Use Cases**  
- ✅ Generating repeated characters (`strings.Repeat("-", 50) → "--------------------------------------------------"`)  
- ✅ Creating padding for text formatting (`strings.Repeat(" ", 5) + "Hello" → "     Hello"`)  
- ✅ Building visual separators for logs, UI elements, or reports  
*/
