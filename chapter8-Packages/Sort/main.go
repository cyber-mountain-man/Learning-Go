package main

import (
	"fmt"
	"sort"
)

// ✅ Define a Person struct
type Person struct {
	Name string // Stores the name of the person
	Age  int    // Stores the age of the person
}

// ✅ Define a custom type `ByName` which is a slice of `Person`
// This is required to implement Go's `sort.Interface`
type ByName []Person

// ✅ `Len` method returns the number of elements in the slice
func (ps ByName) Len() int {
	return len(ps)
}

// ✅ `Less` method defines sorting criteria (sorts by Name in ascending order)
func (ps ByName) Less(i, j int) bool {
	return ps[i].Name < ps[j].Name // Compares names alphabetically
}

// ✅ `Swap` method swaps two elements in the slice
func (ps ByName) Swap(i, j int) {
	ps[i], ps[j] = ps[j], ps[i] // Swaps elements at index i and j
}

func main() {
	// ✅ Define a slice of `Person` structs
	kids := []Person{
		{"Jill", 9},
		{"Jack", 10},
	}

	// ✅ Sort `kids` by Name using `sort.Sort` 
	// (Requires our custom implementation of `sort.Interface`)
	sort.Sort(ByName(kids))

	// ✅ Print the sorted output
	fmt.Println(kids) // Output: [{Jack 10} {Jill 9}]
}

/*
✅ Let's walk through the program:

1️⃣ **Defining the Struct (`Person`)**
   - The `Person` struct contains two fields:
     - `Name` (string) → Represents a person's name.
     - `Age` (int) → Represents the person's age.

2️⃣ **Creating a Custom Sorting Type (`ByName`)**
   - `ByName` is a **slice of Person** that enables sorting by Name.
   - Implements the `sort.Interface`:
     - `Len()`: Returns the length of the slice.
     - `Less()`: Defines sorting order (alphabetical by Name).
     - `Swap()`: Swaps elements to arrange them correctly.

3️⃣ **Sorting the Slice**
   - `sort.Sort(ByName(kids))` sorts the slice **by Name in ascending order**.

4️⃣ **Understanding the Output**
   - The slice starts as `[{"Jill", 9}, {"Jack", 10}]`
   - Since `"Jack"` comes before `"Jill"` alphabetically, the sorted result is:
     ```
     [{Jack 10} {Jill 9}]
     ```

5️⃣ **Real-World Use Cases**
   ✅ Sorting **user lists** in a web application.  
   ✅ Organizing **contacts** alphabetically.  
   ✅ Sorting **log entries by usernames** in **cybersecurity logs**.  

6️⃣ **Want to Sort by Age Instead?**
   - Define a new sorting type **`ByAge`**:
     ```go
     type ByAge []Person

     func (ps ByAge) Len() int { return len(ps) }
     func (ps ByAge) Less(i, j int) bool { return ps[i].Age < ps[j].Age }
     func (ps ByAge) Swap(i, j int) { ps[i], ps[j] = ps[j], ps[i] }
     ```
   - Then call `sort.Sort(ByAge(kids))` to sort by **age instead of name**.

7️⃣ **Why Use `sort.Interface`?**
   - ✅ **Custom sorting logic** (e.g., Name, Age, Salary, etc.).
   - ✅ **Efficient, built-in sorting algorithm**.
   - ✅ **Works on any struct with minimal changes**.
*/
