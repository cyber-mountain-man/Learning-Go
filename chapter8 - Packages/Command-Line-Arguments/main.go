package main

import (
	"flag"       // ✅ For parsing command-line flags
	"fmt"
	"math/rand"  // ✅ For generating random numbers
)

// ✅ main() - Entry point of the program
func main() {
	// 🏷️ Define a command-line flag named "max" with default value 6 and help text
	maxp := flag.Int("max", 6, "the max value")

	// 🚀 Parse the command-line flags provided by the user
	flag.Parse()

	// 🎲 Generate and print a random number between 0 and max (exclusive)
	fmt.Println(rand.Intn(*maxp))
}

/*
✅ Let's walk through the program:

1️⃣ **Defining a Flag (`flag.Int()`)**
   - The program uses `flag.Int("max", 6, "the max value")` to define a flag:
     - `-max` is the name of the flag
     - `6` is the default value
     - `"the max value"` is the help description
   - It returns a pointer to an integer (`maxp`).

2️⃣ **Parsing the Flag (`flag.Parse()`)**
   - This must be called to **populate the flag values** from the command-line input.
   - For example, running `go run main.go -max=10` would update `*maxp` to `10`.

3️⃣ **Using the Flag Value**
   - `rand.Intn(*maxp)` generates a **random number from 0 up to max - 1**.
   - `*maxp` dereferences the pointer to get the actual value.

4️⃣ **Running the Program**
   - With default: `go run main.go` → generates 0–5
   - With flag: `go run main.go -max=100` → generates 0–99

5️⃣ **Expected Output (Example)**
	   4
	   82
	   1
	- Each run will generate a different number unless seeded.

6️⃣ **Real-World Applications**
- ✅ Creating flexible command-line tools
- ✅ Dynamically adjusting program behavior (e.g., log levels, limits)
- ✅ Useful for scripting, automation, and configuration
- ✅ Provides built-in help system: `go run main.go -h`

📌 `flag.Args()` can also be used to retrieve any **non-flag arguments** passed in.
*/
