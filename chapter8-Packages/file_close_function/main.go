package main

import (
	"fmt"
	"os" // Modern file handling package
)

func main() {
	// ✅ Modern Approach (Go 1.16+)
	// Reads the entire file using os.ReadFile()
	bs, err := os.ReadFile("test.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	// Convert byte slice to string and print
	fmt.Println("File contents:\n", string(bs))

	/*
		⚠️ Deprecated Approach (Before Go 1.16)
		import "io/ioutil"

		bs, err := ioutil.ReadFile("test.txt")
		if err != nil {
			fmt.Println("Error reading file:", err)
			return
		}
		fmt.Println("File contents:\n", string(bs))
	*/
}

/*
✅ Let's walk through the program:

1️⃣ **Modern Approach: `os.ReadFile()` (Go 1.16+)**
   - Reads the entire file into memory **without needing `ioutil`**.
   - This is **recommended** since `ioutil.ReadFile()` is **deprecated**.

2️⃣ **Deprecated Approach: `ioutil.ReadFile()` (Before Go 1.16)**
   - Used to perform the same function, but `ioutil` is now obsolete.
   - Functions from `ioutil` were moved to `os` for better consistency.

3️⃣ **Error Handling**
   - If `test.txt` does not exist or cannot be read, an **error message** is displayed.

4️⃣ **Expected Output (if `test.txt` contains "Hello, Go!")**
		File contents: Hello, Go!


5️⃣ **Why the Change?**
- ✅ **Better code organization** → `os.ReadFile()` keeps file operations in one package.
- ✅ **Reduces extra imports** → No need for `ioutil`.

6️⃣ **Real-World Use Cases**
- ✅ **Reading configuration files** (`config.json`, `.env`).
- ✅ **Processing logs or report files**.
- ✅ **Loading text-based data for applications**
*/

