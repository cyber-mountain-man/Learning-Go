package main

import "fmt"

// one takes a pointer to an integer and sets the value to 1
func one(xPtr *int) {
	*xPtr = 1 // Dereferencing the pointer to modify the actual value in memory
}

func main() {
	xPtr := new(int) // Allocates memory for integer and returns a pointer to it (default value is 0)
	one(xPtr)		 // Passes the pointer to the function, which modifies the value
	fmt.Println(*xPtr)  // Prints 1, because the function modified the value at the memory address
}

/*
✅ Let’s walk through the program:

1️⃣ **Using `new()` to Create a Pointer**
   - `new(int)` dynamically allocates memory for an integer and initializes it to `0`.
   - It returns a pointer (`*int`) to the allocated memory.

2️⃣ **Passing a Pointer to a Function**
   - `one(xPtr)` receives the pointer, allowing it to modify the actual value.

3️⃣ **Dereferencing with `*`**
   - Inside `one()`, `*xPtr = 1` changes the value **at the memory address**, not a copy.

4️⃣ **Execution Flow**
   - `xPtr := new(int)` → Allocates an integer, defaulting to `0`.
   - `one(xPtr)` → Modifies the actual value at the memory address.
   - `fmt.Println(*xPtr)` → Prints `1` because `xPtr` now holds `1`.

5️⃣ **Expected Output**
			1
6️⃣ **Why Use Pointers Here?**
- ✅ Prevents unnecessary copying of values.
- ✅ Allows modifying variables inside functions without returning a new value.
- ✅ Useful for working with dynamically allocated memory.
*/