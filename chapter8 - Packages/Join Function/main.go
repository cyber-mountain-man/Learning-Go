package main

import (
	"fmt"
	"strings"
)

func main() {
	// func Join(a []string, sep string) string
	// Joins elements of the slice ["a", "b"] with "-" as a separator
	fmt.Println(strings.Join([]string{"a", "b"}, "-")) // Expected output: "a-b"
}

/*
✅ Let's walk through the program:

1️⃣ **Importing Packages**  
   - `fmt` → Used for printing output to the console.  
   - `strings` → Provides string manipulation functions.

2️⃣ **Using `strings.Join()`**  
   - The function `strings.Join(a []string, sep string)` **concatenates** elements of a slice (`a`) using `sep`.  
   - In this case, `[]string{"a", "b"}` is joined with `"-"`, resulting in `"a-b"`.

3️⃣ **Expected Output**  
			a-b

4️⃣ **Real-World Use Cases**  
- ✅ Constructing CSV or log entries (`strings.Join([]string{"John", "Doe", "30"}, ",") → "John,Doe,30"`)  
- ✅ Creating URLs or file paths (`strings.Join([]string{"home", "user", "docs"}, "/") → "home/user/docs"`)  
- ✅ Formatting sentences (`strings.Join(words, " ")` to combine words into a sentence)  
*/
