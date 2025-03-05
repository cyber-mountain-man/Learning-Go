package main

import "fmt"

func main() {
    slice1 := []int{1, 2, 3}  // Source slice
    slice2 := make([]int, 2)  // Destination slice with length 2

    copy(slice2, slice1) // Copy elements from slice1 → slice2 copy(src,dst) src is source
						 // and dst is destination

    fmt.Println("slice1:", slice1) // Output: [1 2 3]
    fmt.Println("slice2:", slice2) // Output: [1 2]
}
