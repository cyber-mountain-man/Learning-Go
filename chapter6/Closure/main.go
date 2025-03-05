package main

import "fmt"

func main() {
    // Define an anonymous function that adds two integers
    add := func(x, y int) int {
        return x + y // Returns the sum of x and y
    }

    // Call the anonymous function and print the result
    fmt.Println(add(1, 1)) // Output: 2
}

/*
✅ Let’s walk through the program:

1️⃣ **We define an anonymous function**:
   - The function `func(x, y int) int` **takes two integers** as input.
   - It **returns their sum** (`x + y`).
   - This function is assigned to the variable `add`, making it callable.

2️⃣ **Why Use an Anonymous Function?**
   - ✅ **Useful for inline operations** without defining a separate named function.
   - ✅ **Great for one-time-use logic** inside `main()` or other functions.
   - ✅ **Stores the function in a variable**, allowing it to be reused.

3️⃣ **Calling the Anonymous Function**:
   ```go
   fmt.Println(add(1, 1))
*/