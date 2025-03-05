package main

import "fmt"

// 'args ...int' means the function can accept multiple integer arguments
func add(args ...int) int {
    total := 0
    for _, v := range args { // Iterates over all arguments
        total += v
    }
    return total // Returns the sum of all numbers
}

func main() {
    fmt.Println(add(1, 2, 3))    // Output: 6
    fmt.Println(add(5, 10, 15))  // Output: 30
    fmt.Println(add())           // Output: 0 (no numbers passed)
}
