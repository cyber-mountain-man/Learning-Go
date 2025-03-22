package mathutils

import "testing"

func TestAdd(t *testing.T) {
	tests := []struct {
		a, b     int
		expected int
	}{
		{2, 3, 5},
		{-1, -1, -2},
		{0, 0, 0},
		{-5, 5, 0},
		{100, 200, 300},
	}

	for _, tt := range tests {
		result := Add(tt.a, tt.b)
		if result != tt.expected {
			t.Errorf("Add(%d, %d) = %d; want %d", tt.a, tt.b, result, tt.expected)
		}
	}
}

func TestMultiply(t *testing.T) {
	tests := []struct {
		a, b     int
		expected int
	}{
		{2, 3, 6},
		{-1, -1, 1},
		{0, 10, 0},
		{-5, 5, -25},
		{10, 10, 100},
	}

	for _, tt := range tests {
		result := Multiply(tt.a, tt.b)
		if result != tt.expected {
			t.Errorf("Multiply(%d, %d) = %d; want %d", tt.a, tt.b, result, tt.expected)
		}
	}
}

/*
✅ Let's walk through the table-driven tests:

1️⃣ **What Are Table-Driven Tests?**
   - A testing pattern where input and expected output are stored in a list (slice of structs).
   - Allows testing **multiple cases** in a **compact and readable** way.

2️⃣ **Test Table Definition**
   - Each test case is a struct with input fields (`a`, `b`) and an `expected` result.
   - Examples: `{2, 3, 5}`, `{0, 0, 0}`, etc.

3️⃣ **Looping Over the Table**
   - A `for` loop runs each test case by calling the function with the given inputs.
   - If the result doesn't match the expected output, the test fails with a helpful message.

4️⃣ **Why This Is Great**
   - ✅ Easy to scale up with more cases.
   - ✅ Reduces repetitive code.
   - ✅ Makes your test suite clean, concise, and professional.

📌 Use this structure especially for:
   - Math operations
   - Validation functions
   - Business rules and conditional logic
   - Anything with lots of input/output scenarios
*/
