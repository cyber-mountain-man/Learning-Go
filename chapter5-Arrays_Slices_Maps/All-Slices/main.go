package main

import "fmt"

func main() {
    // 1. Basic Slice Declaration
    basicSlice := []int{1, 2, 3, 4, 5} // Slice with direct initialization
    fmt.Println("Basic Slice:", basicSlice)

    // 2. Slicing from an Array
    arr := [...]int{10, 20, 30, 40, 50, 60}
    sliced := arr[1:4] // Includes indices 1,2,3 (not 4)
    fmt.Println("Sliced from Array:", sliced) // Output: [20 30 40]

    // 3. Using Default Low/High Values
    fmt.Println("Slice from start:", arr[:3]) // Equivalent to arr[0:3] → [10 20 30]
    fmt.Println("Slice to end:", arr[2:])  // Equivalent to arr[2:len(arr)] → [30 40 50 60]
    fmt.Println("Full Slice:", arr[:])     // Equivalent to arr[0:len(arr)] → [10 20 30 40 50 60]

    // 4. Creating a Slice with `make()`
    madeSlice := make([]int, 3, 5) // Length = 3, Capacity = 5
    fmt.Println("Slice with make():", madeSlice, "Length:", len(madeSlice), "Capacity:", cap(madeSlice))

    // 5. Appending to a Slice
    madeSlice = append(madeSlice, 100, 200) // Appending values
    fmt.Println("After Append:", madeSlice, "New Length:", len(madeSlice), "New Capacity:", cap(madeSlice))

    // 6. Copying a Slice
    sourceSlice := []int{7, 8, 9}
    destinationSlice := make([]int, len(sourceSlice)) // Allocate memory
    copy(destinationSlice, sourceSlice) // Copy elements
    fmt.Println("Copied Slice:", destinationSlice) // Output: [7 8 9]
}

/*
Let’s walk through the program:

1. First, we create a basic slice using direct initialization.
2. Next, we slice an existing array using the `[low:high]` syntax:
   - The `low` index is **included**, but the `high` index is **excluded**.
   - Example: `arr[1:4]` includes elements at indices `1, 2, 3` (but not `4`).
3. We use **default slicing**:
   - `arr[:3]` means "from the start up to index 3".
   - `arr[2:]` means "from index 2 to the end".
   - `arr[:]` creates a slice of the full array.
4. We use `make()` to create a slice with an initial length and capacity.
5. We demonstrate how `append()` dynamically adds elements to a slice.
6. Finally, we use `copy()` to copy one slice into another.

This program demonstrates how slices provide flexibility compared to arrays,
making them one of the most commonly used data structures in Go.
*/
