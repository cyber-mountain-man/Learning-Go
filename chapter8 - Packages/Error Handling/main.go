package main

import (
	"errors" // Package for creating error values
	"fmt"
)

func main() {
	// Create a custom error message
	err := errors.New("this is a custom error")

	// Print the error message
	fmt.Println("Error:", err)
}

/*
✅ Let's walk through the program:

1️⃣ **Importing the `errors` Package**
   - The `errors` package provides a function **`errors.New("message")`** to create **custom error messages**.

2️⃣ **Creating a Custom Error**
   - `err := errors.New("this is a custom error")` generates an **error object**.
   - This **error can be returned from functions** to indicate failures.

3️⃣ **Printing the Error**
   - `fmt.Println("Error:", err)` prints:  
     ```
     Error: this is a custom error
     ```

4️⃣ **When to Use `errors.New()`?**
   - ✅ **For defining simple error messages**.
   - ✅ **Returning error conditions from functions**.

5️⃣ **🚀 Next Steps: Use Custom Errors in Functions**
   - Instead of just creating an error, return it from a function:
     ```go
     func doSomething() error {
         return errors.New("something went wrong")
     }
     ```
   - The caller can then **check and handle the error**.
*/
