package main

import (
	"container/list" // Importing the package for doubly linked lists
	"fmt"
)

func main() {
	// ✅ Create a new linked list
	var x list.List // Zero value is an empty list

	// ✅ Push values to the back of the list
	x.PushBack(1)
	x.PushBack(2)
	x.PushBack(3)

	// ✅ Loop through the linked list from front to back
	for e := x.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value.(int)) // Type assertion to extract int value
	}
}

/*
✅ Let's walk through the program:

1️⃣ **Importing the `container/list` Package**
   - The `list` package provides a **doubly linked list** implementation.

2️⃣ **Creating a List**
   - `var x list.List` initializes an **empty linked list**.
   - Alternatively, we could use `x := list.New()`.

3️⃣ **Adding Elements**
   - `x.PushBack(1)`, `x.PushBack(2)`, `x.PushBack(3)`
   - Adds elements **to the back of the list**.

4️⃣ **Iterating Over the List**
   - `x.Front()` gets the **first element**.
   - The loop **follows `Next()` pointers** until `nil` is reached.

5️⃣ **Type Assertion (`.(int)`)**
   - Since **`e.Value` is of type `interface{}`**, we use **`e.Value.(int)`** to retrieve the actual integer value.

6️⃣ **Expected Output**
		1 2 3


7️⃣ **Why Use a Linked List?**
- ✅ **Efficient insertions and deletions** → No need to shift elements like slices.
- ✅ **Useful when the size of the list is unknown or changing dynamically**.
- ✅ **Doubly linked** → Can navigate **both forward and backward**.
*/
