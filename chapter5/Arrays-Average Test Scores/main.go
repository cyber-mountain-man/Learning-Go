package main

import "fmt"

func main() {
	var x [5]float64
	x[0] = 98
	x[1] = 93
	x[2] = 77
	x[3] = 82
	x[4] = 83

	var total float64 = 0
	for i := 0; i < 5; i++ {
		total += x[i]
	}
	fmt.Println("The average test score is: ", total / 5)
}

/*
Let’s walk through the program:
1. First, we create an array of length 5 to hold our test scores, then we fill up each
element with a grade.
2. Next, we set up a for loop to compute the total score.
3. Finally, we divide the total score by the number of elements to find the average.


*/