package main

import "fmt"

var x string = "Hello, World"

func main() {
	fmt.Println(x)
}

// The f function now has access to the x variable
func f() {
	fmt.Println(x)
}