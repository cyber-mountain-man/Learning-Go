package main

import (
	"fmt"
	"strings"
)

func main() {
	// func Split(s, sep string) []string
	// Splits the string "a-b-c-d-e" into a slice of substrings using "-" as the separator
	fmt.Println(strings.Split("a-b-c-d-e", "-")) // Expected output: []string{"a", "b", "c", "d", "e"}
}

/*
✅ Let's walk through the program:

1️⃣ **Importing Packages**  
   - `fmt` → Used for printing output to the console.  
   - `strings` → Provides string manipulation functions.

2️⃣ **Using `strings.Split()`**  
   - The function `strings.Split(s, sep)` **splits** `s` into substrings **based on the separator `sep`**.  
   - In this case, `"a-b-c-d-e"` is split at every `"-"`, resulting in `[]string{"a", "b", "c", "d", "e"}`.

3️⃣ **Expected Output**  
	["a" "b" "c" "d" "e"]

(Go prints slices with spaces instead of commas)

4️⃣ **Real-World Use Cases**  
- ✅ **Parsing CSV data** (`strings.Split("John,Doe,30", ",")`)  
- ✅ **Breaking apart URLs** (`strings.Split("https://example.com/home", "/")`)  
- ✅ **Tokenizing user input** (`strings.Split("hello world!", " ")`)  
*/
