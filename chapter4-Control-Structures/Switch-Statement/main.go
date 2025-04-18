package main

import "fmt"

func main() {
	for i:= 0; i <= 10; i++ {
		if i % 2 == 0 {
			fmt.Println(i, "even")
		}else {
			fmt.Println(i, "odd")
		}

		// Number word representation
		switch i {
		case 0:
			fmt.Println("Zero")
		case 1:
			fmt.Println("One")
		case 2:
			fmt.Println("Two")
		case 3:
			fmt.Println("Three")
		case 4:
			fmt.Println("Four")
		case 5:
			fmt.Println("Five")
		case 6:
			fmt.Println("Six")
		case 7:
			fmt.Println("Seven")
		case 8:
			fmt.Println("Eight")
		case 9:
			fmt.Println("Nine")
		case 10:
			fmt.Println("Ten")
		default: 
			fmt.Println("Unknown Number")
		}
		fmt.Println() // Blank line for readability
	}
}