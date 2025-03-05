package main

import "fmt"

// safeDivide attempts to divide two numbers and handles division by zero gracefully
func safeDivide(a, b int) {
	defer func() {
		// Recover from panic if it occurs
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r) // Print the panic message
		}
	}()

	// If denominator is zero, trigger a panic
	if b == 0 {
		panic("division by zero is not allowed")
	}

	// Perform division if no panic occurs
	fmt.Println("Result:", a/b)
}

func main() {
	safeDivide(10, 2) // Works fine, prints "Result: 5"
	safeDivide(5, 0)  // Causes panic, but it's caught and handled
	fmt.Println("Program continues after panic recovery") // Ensures program execution continues
}

/*
✅ Let’s walk through the program:

1️⃣ **Understanding `panic` and `recover`**
   - `panic("division by zero is not allowed")` stops execution **immediately**.
   - `recover()` catches the panic inside a **`defer` function**, allowing the program to continue.

2️⃣ **How the `safeDivide` function works**
   - The function **attempts to divide `a` by `b`**.
   - If `b == 0`, it **triggers a panic**.
   - A **deferred function runs `recover()`**, which **catches the panic and prints an error message** instead of crashing.

3️⃣ **Execution Flow**
   - `safeDivide(10, 2)` → Works fine, prints `Result: 5`.
   - `safeDivide(5, 0)` → Causes **panic**, but `recover()` **handles it** and prints `"Recovered from panic: division by zero is not allowed"`.
   - The program **continues running**, instead of crashing.

4️⃣ **Expected Output**
   - Result: 5 Recovered from panic: division by zero is not allowed Program continues after panic recovery

5️⃣ **Why Use `panic` and `recover`?**
- ✅ **Prevents program crashes** when encountering critical errors.
- ✅ **Allows graceful error handling** in functions that could fail unexpectedly.
- ✅ **Useful for error handling in web servers, middleware, and CLI tools**.
*/