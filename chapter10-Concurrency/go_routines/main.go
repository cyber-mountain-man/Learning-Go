package main

import "fmt"

// ✅ f prints a number (n) and its counter (0 to 9)
func f(n int) {
	for i := 0; i < 10; i++ {
		fmt.Println(n, ":", i)
	}
}

func main() {
	// ✅ Start function f(0) in a new goroutine
	go f(0)

	// ✅ Prevent main() from exiting immediately
	var input string
	fmt.Scanln(&input)
}
/*
✅ Let's walk through the program:

1️⃣ **Function Definition (`f`)**
   - A simple loop that prints the input `n` and a counter `i` from 0 to 9.
   - Helps visualize which goroutine is running when.

2️⃣ **Launching a Goroutine (`go f(0)`)**
   - Uses the `go` keyword to run `f(0)` **concurrently**.
   - This does **not block** the main thread — it continues running immediately.
   - Goroutines are lightweight threads managed by the Go runtime.

3️⃣ **Blocking Main from Exiting (`fmt.Scanln`)**
   - Without this, `main()` would exit before the goroutine has a chance to run.
   - This line pauses the program, letting the goroutine complete its work.

4️⃣ **Expected Output**
		0 : 0
		0 : 0
		0 : 1
		0 : 2
		...
		0 : 9

		- Output order is predictable here because only one goroutine is running (besides main).
- With **multiple goroutines**, output may interleave in a non-deterministic order.

---

✅ **Why Use Goroutines?**
- ✅ For **concurrent processing** — like handling web requests, file I/O, or background tasks.
- ✅ Goroutines are **much lighter than OS threads**.
- ✅ Easily scale to **thousands of concurrent tasks**.

📌 **Next Steps**
- ✅ Try spawning multiple goroutines like `go f(1)`, `go f(2)`.
- ✅ Introduce `time.Sleep()` to simulate delays.
- ✅ Later, explore **channels** to manage communication between goroutines.

This is a foundational step toward learning **concurrency** in Go!
*/
