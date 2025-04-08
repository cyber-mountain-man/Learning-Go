package main

import (
	"fmt"       // For printing output to the terminal
	"math/rand" // For generating random numbers
	"sync"      // For synchronizing concurrent tasks using WaitGroup
	"time"      // For adding delays and working with time durations
)

// serveCustomers simulates a cashier serving 10 customers
// Each cashier runs in its own goroutine and uses a Waitgroup to signal completion.
func serveCustommers(cashierID int, wg *sync.WaitGroup) {
	defer wg.Done() // Defer marks this function as done when it finishes, so the WaitGroup counter decreases.

	for i := 1; i <= 10; i++ {
		// Print a message showing which cashier is serving which customer.
		fmt.Printf("Cashier %d is serving customer %d\n", cashierID, i)

		// Simulates variable processing time (0-249 milliseconds).
		delay := time.Duration(rand.Intn(250)) * time.Millisecond
		
		// Sleep to simulate the cashier working on the customer
		time.Sleep(delay)
	}
}

func main() {
	var wg sync.WaitGroup // Used to wait for all cashier goroutines to finish

	numCashiers := 10 // Total number of cashiers to stimulate

	// Launch each cashier in its own goroutine
	for i := 1; i <= numCashiers; i++ {
		wg.Add(1)							// Tell WaitFroup there is one more goroutine to wait for
		go serveCustommers(i, &wg) 	// Run serveCustomers concurrently for cashier i
	}

	wg.Wait() // Wait here unitl all goroutines have valled wg.Done()

	// All goroutines have finished - print a message
	fmt.Println(" All customers have been served.")
}