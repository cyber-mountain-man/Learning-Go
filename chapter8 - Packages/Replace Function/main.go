package main

import (
	"fmt"
	"strings"
)

func main() {
	// func Replace(s, old, new string, n int) string
	// Replaces the first 2 occurrences of "a" with "b" in the string "aaaa"
	fmt.Println(strings.Replace("aaaa", "a", "b", 2)) // Expected output: "bbaa"
}

/*
✅ Let's walk through the program:

1️⃣ **Importing Packages**  
   - `fmt` → Used for printing output to the console.  
   - `strings` → Provides string manipulation functions.

2️⃣ **Using `strings.Replace()`**  
   - The function `strings.Replace(s, old, new, n)` **replaces** up to `n` occurrences of `old` with `new` in `s`.  
   - `"aaaa"` → Replace **first 2 occurrences** of `"a"` with `"b"` → `"bbaa"`  
   - If `n` is `-1`, **all occurrences** will be replaced.

3️⃣ **Expected Output**  
		"bbaa"

4️⃣ **Real-World Use Cases**  
- ✅ Censoring words in text (`strings.Replace(text, "badword", "****", -1)`)  
- ✅ Formatting user input (`strings.Replace(phone, "-", "", -1) → remove dashes from phone numbers`)  
- ✅ Replacing specific characters in a dataset  
*/
