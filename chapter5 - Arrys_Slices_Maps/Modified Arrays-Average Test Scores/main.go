package main

import "fmt"

func main() {
    // Create an array of 5 float64 values to store test scores
    var x = [5]float64{98, 93, 77, 82, 83}

    var total float64 = 0

    // Iterate through the array using range to compute the total score
    for _, value := range x {
        total += value
    }

    // Calculate and print the average score
    fmt.Println("The average test score is: ",total / float64(len(x)))
}

/*
Let’s walk through the program:
1. First, we create an array of length 5 to hold our test scores, then we fill up each
   element with a grade.
2. Next, we set up a for loop to compute the total score.
3. Finally, we divide the total score by the number of elements to find the average.
*/
