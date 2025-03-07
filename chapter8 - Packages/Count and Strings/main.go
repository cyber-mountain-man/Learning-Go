package main

import (
	"fmt"
	"strings"
)

func main() {
	// func Count(s, sep string) int
	// Counts the occurrences of "t" in the string "test"
	fmt.Println(strings.Count("test", "t")) // Expected output: 2
}

/*
✅ Let's walk through the program:

1️⃣ **Importing Packages**  
   - `fmt` → Used for printing output to the console.  
   - `strings` → Provides string manipulation functions.

2️⃣ **Using `strings.Count()`**  
   - The function `strings.Count(s, substr)` counts how many times `substr` appears in `s`.  
   - In this case, `"test"` contains `"t"` **twice**, so it returns `2`.

3️⃣ **Expected Output**  
		2

4️⃣ **Real-World Use Cases**  
- ✅ Counting occurrences of a letter in a word.  
- ✅ Checking how many times a keyword appears in a text.  
- ✅ Analyzing character frequency in passwords.  
*/
