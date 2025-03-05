package main

import "fmt"

// zero takes an integer parameter but modifies only its local copy
func zero(x int) { 
	x = 0 // Modifies the local copy of x, not the original variable
}

func main() {
	x := 5           // Declare x and assign value 5
	zero(x)         // Call zero(x), but it passes a copy of x
	fmt.Println(x)  // x remains 5, because zero() modified a copy, not the original variable
}

/*
✅ Let’s walk through the program:

1️⃣ **Understanding How Go Passes Arguments**
   - Go **passes function arguments by value**, meaning functions receive **a copy** of the variable, not the original.
   - `zero(x int)` receives a **copy of `x`**, so changes inside the function do **not** affect the original variable.

2️⃣ **Execution Flow**
   - `x` is initialized as `5`.
   - `zero(x)` is called:
     ```
     zero(5) // Passes a copy of x (5) to the function
     ```
   - Inside `zero()`, the copy of `x` is changed to `0`, but **this does not affect the original `x` in `main()`**.
   - `fmt.Println(x)` still prints **5**, since `x` in `main()` was never changed.

3️⃣ **Expected Output**
			5
			
4️⃣ **Key Takeaways**
- **Primitive types (`int`, `float`, `bool`, etc.) are passed by value** (function receives a copy).
- **Modifications inside the function do not affect the original variable**.
- **To modify a variable inside a function, use pointers (`*int`) instead**.

5️⃣ **How to Modify the Original Variable?**
- If you need the function to **modify `x`**, use a **pointer** (`*int`):
  ```go
  func zeroPointer(x *int) {
      *x = 0 // Modify the actual variable via pointer
  }

  func main() {
      x := 5
      zeroPointer(&x) // Pass address of x
      fmt.Println(x)  // x is now 0
  }
  ```
- This will output:
  ```
  0
  ```
*/