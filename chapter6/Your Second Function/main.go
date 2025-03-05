package main

import "fmt"

// average calculates the average of a slice of float64 numbers.
func average(xs []float64) float64 {
    total := 0.0

    // Loop through each value in the slice and sum them up.
    for _, v := range xs {
        total += v
    }

    // Return the total divided by the length of the slice to get the average.
    return total / float64(len(xs))
}

func main() {
    // Define a slice of float64 numbers
    xs := []float64{98, 93, 77, 82, 83}

    // Call the average function and print the result
    fmt.Println(average(xs))
}

/*
✅ Let’s walk through the program:

1️⃣ **The `average` function is defined**:
   - It accepts a slice of `float64` numbers (`xs`).
   - It initializes `total` to `0.0`.
   - A `for` loop iterates through each value in `xs`, adding them to `total`.
   - The final average is calculated as `total / float64(len(xs))`.

2️⃣ **Inside `main`**:
   - A slice (`xs`) of five numbers is defined.
   - The `average` function is called with `xs` as an argument.
   - The result is printed.

3️⃣ **Expected Output**:
- The function calculates `(98 + 93 + 77 + 82 + 83) / 5 = 86.6`.

4️⃣ **Why This Approach?**:
- ✅ Uses a function (`average`) to make the code reusable.
- ✅ Uses a **for-loop with `_`** to iterate efficiently.
- ✅ Uses **float64 conversion** to ensure correct division.
*/