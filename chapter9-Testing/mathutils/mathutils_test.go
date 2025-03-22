package mathutils

import "testing"

// ✅ TestAdd verifies that the Add function returns the correct result.
func TestAdd(t *testing.T) {
	result := Add(2, 3)      // 🔢 Call the Add function with input 2 and 3
	expected := 5            // 🎯 Expected result

	if result != expected {
		t.Errorf("Add(2, 3) = %d; want %d", result, expected) // ❌ Report if result is wrong
	}
}

// ✅ TestMultiply checks if Multiply returns correct multiplication result.
func TestMultiply(t *testing.T) {
	result := Multiply(4, 5)  // 🔢 Call the Multiply function with 4 and 5
	expected := 20            // 🎯 Expected result

	if result != expected {
		t.Errorf("Multiply(4, 5) = %d; want %d", result, expected) // ❌ Report mismatch
	}
}

// 🧪 Benchmark Test for Add
func BenchmarkAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Add(10, 20)
	}
}

// 🧪 Benchmark Test for Multiply
func BenchmarkMultiply(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Multiply(10, 20)
	}
}

/*
✅ Let's walk through the test suite:

1️⃣ **Test Structure (`func TestXxx(*testing.T)`)**
   - All test functions start with `Test` and take a `*testing.T` parameter.
   - Go uses this naming convention to auto-discover and run the tests with `go test`.

2️⃣ **Running the Function Under Test**
   - `Add(2, 3)` and `Multiply(4, 5)` validate real inputs against expected outputs.
   - These functions are defined in the `mathutils` package.

3️⃣ **Error Reporting**
   - `t.Errorf()` logs a failure if the result doesn't match what's expected.
   - Makes it easy to identify what failed and why.

4️⃣ **Benchmark Tests (`func BenchmarkXxx(*testing.B)`)**
   - Benchmarks measure performance using the built-in `testing.B` type.
   - `b.N` is dynamically set by Go to run enough iterations for reliable measurements.
   - Example:
     ```go
     for i := 0; i < b.N; i++ {
         Add(10, 20)
     }
     ```

5️⃣ **How to Run the Tests and Benchmarks**
   - To run unit tests:
     ```bash
     go test
     ```
   - To run benchmarks:
     ```bash
     go test -bench=.
     ```

6️⃣ **Expected Benchmark Output**

	goos: windows
	goarch: amd64 
	BenchmarkAdd-8 500000000 2.50 ns/op 
	BenchmarkMultiply-8 300000000 4.00 ns/op 
	PASS 
	ok mathutils 2.345s

	
7️⃣ **Why This Is Useful**
- ✅ Confirms code correctness and performance.
- ✅ Catches bugs early and prevents regressions.
- ✅ Encourages **Test-Driven Development (TDD)**.
- ✅ Helps tune performance with benchmarks.
- ✅ Seamlessly integrates into CI/CD pipelines.

📌 **Real-World Tips**
- Use **table-driven tests** to test multiple input/output cases cleanly.
- Keep tests in `_test.go` files to separate them from production code.
- Add benchmarking for computational or frequently called code to detect performance regressions over time.
*/
