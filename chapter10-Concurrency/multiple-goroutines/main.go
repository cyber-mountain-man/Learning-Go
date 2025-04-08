package main

import (
	"fmt"
	"sync"
)

// ✅ f prints numbers for a given ID
func f(id int, wg *sync.WaitGroup) {
	defer wg.Done() // ✅ Marks this goroutine as done when it finishes
	for i := 0; i < 3; i++ {
		fmt.Printf("Goroutine %d → %d\n", id, i)
	}
}

func main() {
	var wg sync.WaitGroup // ✅ WaitGroup to wait for all goroutines

	// ✅ Launch 5 goroutines
	for i := 0; i < 5; i++ {
		wg.Add(1)           // 🚀 Register a new goroutine
		go f(i, &wg)        // 🔁 Launch goroutine
	}

	wg.Wait() // 🕒 Wait for all goroutines to complete
	fmt.Println("✅ All goroutines completed.")
}

/*
✅ Let's walk through the program:

1️⃣ **What Is a Goroutine?**
   - A lightweight thread managed by the Go runtime.
   - Created with the `go` keyword (e.g., `go f(i)`).

2️⃣ **What Does This Program Do?**
   - Starts 5 goroutines using a loop.
   - Each goroutine prints 3 messages with its own identifier.

3️⃣ **Why Use sync.WaitGroup?**
   - Prevents `main()` from exiting before all goroutines finish.
   - `wg.Add(1)` registers a new goroutine.
   - `wg.Done()` marks a goroutine as finished.
   - `wg.Wait()` blocks until all registered goroutines are done.

4️⃣ **Expected Output Example**
	 	Goroutine 1 → 0
		Goroutine 1 → 1
		Goroutine 1 → 2
		Goroutine 3 → 0
		Goroutine 3 → 1
		Goroutine 3 → 2
		Goroutine 4 → 0
		Goroutine 4 → 1
		Goroutine 4 → 2
		Goroutine 2 → 0
		Goroutine 2 → 1
		Goroutine 2 → 2
		Goroutine 0 → 0
		Goroutine 0 → 1
		Goroutine 0 → 2
		✅ All goroutines completed.
...
⚠ Output order may vary due to concurrency.

5️⃣ **Why Use Goroutines?**
- ✅ Efficient concurrency.
- ✅ Great for parallel tasks, I/O handling, background jobs, etc.
- ✅ Scales better than OS threads.

📌 Real-World Use Cases:
- Web servers handling multiple requests concurrently.
- Workers processing jobs from a queue.
- Polling APIs or services in the background.
*/
