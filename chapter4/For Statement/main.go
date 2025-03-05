package main

import "fmt"

func main() {
	i := 1
	for i <= 10 {
		fmt.Println(i)
		i = i + 1
	}
}

/*
As an exercise, let’s walk through the program like a computer would:
1. Create a variable named i with the value 1.
2. Is i <= 10? Yes.
3. Print i.
4. Set i to i + 1 (i now equals 2).
5. Is i <= 10? Yes.
6. Print i.
7. Set i to i + 1 (i now equals 3).
8. …
9. Set i to i + 1 (i now equals 11).
10. Is i <= 10? No.
11. Nothing left to do, so exit.
*/