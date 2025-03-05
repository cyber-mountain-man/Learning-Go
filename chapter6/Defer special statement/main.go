package main

import "fmt"

// Function that prints "1st"
func first() {
    fmt.Println("1st")
}

// Function that prints "2nd"
func second() {
    fmt.Println("2nd")
}

// Function that prints "3rd"
func third() {
    fmt.Println("3rd")
}

func main() {
    defer second() // Defer execution of second() until the end of main()
    first()        // Executes immediately
    third()        // Executes immediately
}

/*
✅ Let’s walk through the program:

1️⃣ **Understanding `defer`**  
   - The `defer` keyword postpones the execution of a function **until the surrounding function (`main()`) exits**.
   - `defer second()` **does not execute immediately**, but it is scheduled to run **just before `main()` returns**.

2️⃣ **Execution Order**  
   - `first()` runs **immediately** → prints `"1st"`.
   - `third()` runs **immediately** → prints `"3rd"`.
   - `second()` (which was deferred) runs **at the end** of `main()`.

3️⃣ **Expected Output**:
- Even though `second()` is **declared before `first()` in `main()`**, it runs **last** because it was deferred.

4️⃣ **Why Use `defer`?**
- ✅ Ensures that cleanup operations (like **closing files, releasing locks, or disconnecting from databases**) **run at the end** of a function.
- ✅ Useful for maintaining **readable and predictable control flow**.

📌 **Next Steps:**
- ✅ Modify this program to use multiple `defer` calls and see how they execute in **LIFO (Last-In, First-Out) order**.
- ✅ Use `defer` to safely **close a file or database connection** in a real-world scenario.
*/