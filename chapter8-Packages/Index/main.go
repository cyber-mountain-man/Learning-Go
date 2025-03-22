package  main

import (
	"fmt"
	"strings"
)

func main() {
	// func Index(s, sep string) int
	// Finds the index of the first occurrence of "e" in the string "test"
	fmt.Println(strings.Index("test", "e")) // Expected output: 1
}

/*
✅ Let's walk through the program:

1️⃣ **Importing Packages**  
   - `fmt` → Used for printing output to the console.  
   - `strings` → Provides string manipulation functions.

2️⃣ **Using `strings.Index()`**  
   - The function `strings.Index(s, substr)` returns the index of the **first occurrence** of `substr` in `s`.  
   - In this case, `"e"` appears at index **1** in `"test"`, so it returns `1`.  
   - If the substring is **not found**, it returns `-1`.

3️⃣ **Expected Output**  
			1

4️⃣ **Real-World Use Cases**  
- ✅ Finding the position of a character in a string.  
- ✅ Searching for the first occurrence of a keyword in text.  
- ✅ Checking if a substring exists (`if strings.Index(s, "word") != -1`).  
*/
