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

/*
✅ Let's walk through the code:

1️⃣ Test Structure (`func TestXxx(*testing.T)`)
   - All test functions in Go must start with `Test` and take a `*testing.T` argument.
   - Go automatically detects and runs these with `go test`.

2️⃣ Running the Function Under Test
   - `Add(2, 3)` and `Multiply(4, 5)` are examples of using real values to verify correct behavior.
   - These functions are assumed to be defined in `mathutils.go`.

3️⃣ Error Reporting
   - If the output doesn't match the expected result, `t.Errorf()` logs an error.
   - The test fails gracefully and reports what went wrong.

4️⃣ Running the Tests
   Use this in the terminal from the project root or test directory:

   
5️⃣ Why This Is Useful
- ✅ Confirms your code works as intended.
- ✅ Catches bugs early during development.
- ✅ Ensures updates don't break existing functionality (regression testing).
- ✅ Encourages good software engineering practices like **Test-Driven Development (TDD)**.
- ✅ Integrates smoothly with CI/CD pipelines for automation.

📌 Real-World Tip:
- Use table-driven tests to validate many input/output pairs.
- Consider separating test files from implementation using `_test.go` suffix as done here.
*/
