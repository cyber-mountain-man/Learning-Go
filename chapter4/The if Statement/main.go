/*
Specifies whether or not the number is even or odd:
1 odd
2 even
3 odd
4 even
5 odd
6 even
7 odd
8 even
9 odd
10 even
*/

package main

import "fmt"

func main() {
	for i := 1; i <= 10; i++ {
		if i % 2 == 0 {
			fmt.Println(i, "even")
		} else{
			fmt.Println(i, "odd")
		}
	}
}

/*
Let’s walk through this program:
1. Create a variable i of type int and give it the value 1.
2. Is i less than or equal to 10? Yes: jump to the if block.
3. Is the remainder of i ÷ 2 equal to 0? No: jump to the else block.
4. Print i followed by odd.
5. Increment i (the statement after the condition).
6. Is i less than or equal to 10? Yes: jump to the if block.
7. Is the remainder of i ÷ 2 equal to 0? Yes: jump to the if block.
8. Print i followed by even, and so on until i is equal to 11.
9. …
*/
