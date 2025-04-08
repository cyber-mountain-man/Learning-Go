package main

import (
	"fmt"
	"math/rand"
	"time"
)

// ✅ f() prints a number and a counter with a random delay between prints
func f(n int) {
	for i := 0; i < 10; i++ {
		fmt.Println(n, ":", i)

		// ⏱ Introduce a small random delay to simulate work
		amt := time.Duration(rand.Intn(250)) // random 0-249 ms
		time.Sleep(time.Millisecond * amt)
	}
}

func main() {
	// 🚀 Launch 10 goroutines concurrently
	for i := 0; i < 10; i++ {
		go f(i) // Each goroutine runs function f with its own value of i
}
}