package main

import (
	"fmt"
	"io/ioutil" // ⚠️ Deprecated: ioutil package functions are deprecated in Go 1.16+
)

func main() {
	// Writing to a file
	// ⚠️ Deprecated: ioutil.WriteFile() should be replaced with os.WriteFile() in newer versions of Go
	err := ioutil.WriteFile("test.txt", []byte("Hello, Go!\n"), 0644)
	if err != nil {
		fmt.Println("Error writing file:", err)
		return
	}

	// Reading from a file
	// ⚠️ Deprecated: ioutil.ReadFile() should be replaced with os.ReadFile()
	data, err := ioutil.ReadFile("test.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Print file contents
	fmt.Println("File contents:", string(data))
}

/*
✅ Let's walk through the program:

1️⃣ **Writing to a File (`ioutil.WriteFile()`)**
   - Creates `test.txt` and writes `"Hello, Go!\n"` into it.
   - **⚠️ Deprecated:** In Go 1.16+, use `os.WriteFile()` instead.

2️⃣ **Reading from a File (`ioutil.ReadFile()`)**
   - Reads the contents of `test.txt` into a byte slice.
   - Converts it to a string and prints it.
   - **⚠️ Deprecated:** In Go 1.16+, use `os.ReadFile()` instead.

3️⃣ **Expected Output (Terminal)**
	  File contents: Hello, Go!


4️⃣ **🔧 Modern Alternative (Go 1.16+)**
- Replace `ioutil.WriteFile()` with `os.WriteFile()`:
  ```go
  import "os"
  os.WriteFile("test.txt", []byte("Hello, Go!\n"), 0644)
  ```
- Replace `ioutil.ReadFile()` with `os.ReadFile()`:
  ```go
  data, err := os.ReadFile("test.txt")
  ```

5️⃣ **Why Was `ioutil` Deprecated?**
- **Better Code Organization:** Functions were moved to `os` & `io` for consistency.
- **No Need for Extra Package:** Now integrated directly in `os`.

6️⃣ **Real-World Applications**
- ✅ **Saving user-generated text data**
- ✅ **Processing small configuration files**
- ✅ **Reading and writing logs in scripts**
*/
