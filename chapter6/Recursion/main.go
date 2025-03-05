package main

import "fmt"

// Recursive function to calculate the factorial of a number
func factorial(x uint) uint {
    if x == 0 { // Base case: factorial(0) is defined as 1
        return 1
    }
    return x * factorial(x-1) // Recursive case: x * factorial(x-1)
}

func main() {
    fmt.Println(factorial(2)) // Output: 2
}

/*
✅ Let’s walk through the program:

1️⃣ **What is Factorial?**  
   - The **factorial** of a number (`n!`) is the product of all positive integers **from 1 to `n`**.
   - **Formula:** `n! = n * (n-1) * (n-2) * ... * 1`
   - **Examples:**  
     ```
     2! = 2 * 1 = 2
     3! = 3 * 2 * 1 = 6
     5! = 5 * 4 * 3 * 2 * 1 = 120
     ```

2️⃣ **Recursive Function Explanation (`factorial(x)`)**:
   - **Base Case:** If `x == 0`, return `1` (since `0! = 1`).
   - **Recursive Case:** Multiply `x` by `factorial(x-1)`, calling itself **until `x` reaches 0**.

3️⃣ **What Happens When We Call `factorial(2)`?**
   - The function keeps calling itself recursively:
     ```
     factorial(2) → 2 * factorial(1)
     factorial(1) → 1 * factorial(0)
     factorial(0) → 1 (Base case reached)
     ```
     - Resolving:
       ```
       factorial(1) = 1 * 1 = 1
       factorial(2) = 2 * 1 = 2
       ```

4️⃣ **Expected Output**:
- The lookup is successful because `2! = 2 * 1 = 2`.

5️⃣ **Why Use Recursion Here?**
- ✅ Recursion makes the function **short and easy to understand**.
- ✅ Useful for mathematical calculations like **factorials, Fibonacci sequences, tree traversal**.
- ❌ **Not efficient for large numbers** (may cause **stack overflow** if `x` is too large).

📌 **Next Steps:**
- ✅ Modify this program to compute **larger factorials efficiently** (e.g., `10!`).  
- ✅ Replace recursion with **iteration** to avoid **stack overflow** for large numbers.
- ✅ Use **`math/big`** to handle **very large factorials** (e.g., `100!`).
*/