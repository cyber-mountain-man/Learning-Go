package main

import "fmt"

// zero takes a pointer to an integer and modifies the original value
func zero(xPtr *int) {
	*xPtr = 0 // Dereferencing the pointer to modify the original variable
}

func main() {
	x := 5 			// Declare x and assign value 5
	zero(&x)		// Pass the memory address of x to zero()
	fmt.Println(x)  // x is now 0, because the function modified the actual variable
}

/*
✅ Let’s walk through the program:

1️⃣ **Understanding Pointers in Go**
   - **Pointers allow functions to modify the original variable** instead of a copy.
   - `*int` means "pointer to an integer" (stores a memory address).
   - `*xPtr = 0` means "change the value at the memory address xPtr points to."

2️⃣ **Execution Flow**
   - `x` is initialized as `5`.
   - `zero(&x)` is called, passing the **address of `x`**.
   - Inside `zero()`, `*xPtr = 0` modifies **the actual value of `x` in memory**.
   - `fmt.Println(x)` prints `0` because `x` was changed.

3️⃣ **Expected Output**
			0


4️⃣ **Key Takeaways**
- **Pass-by-value (default behavior in Go)** → Functions receive **copies** of variables.
- **Pass-by-reference (using pointers)** → Functions can **modify the original variable**.
- **Use pointers when you need to modify a variable inside a function**.

5️⃣ **Comparison With Non-Pointer Version**
- **Non-pointer version (previous example)**: `x` remains **unchanged** because `zero(x)` only modifies a copy.
- **Pointer version (this example)**: `x` **changes** because `zero(&x)` modifies the original variable in memory.

6️⃣ **When to Use Pointers?**
- ✅ When modifying a variable inside a function.
- ✅ When passing **large data structures** to avoid unnecessary copying.
- ✅ When working with **structs and slices** to manipulate original data.

*/
