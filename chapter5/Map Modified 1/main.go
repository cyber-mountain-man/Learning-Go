package main

import "fmt"

func main() {
    // Define a nested map to store chemical elements and their states
    elements := map[string]map[string]string{
        "H": {  // Hydrogen
            "name":  "Hydrogen",
            "state": "gas",
        },
        "He": { // Helium
            "name":  "Helium",
            "state": "gas",
        },
        "Li": { // Lithium
            "name":  "Lithium",
            "state": "solid",
        },
        "Be": { // Beryllium
            "name":  "Beryllium",
            "state": "solid",
        },
        "B": { // Boron
            "name":  "Boron",
            "state": "solid",
        },
        "C": { // Carbon
            "name":  "Carbon",
            "state": "solid",
        },
        "N": { // Nitrogen
            "name":  "Nitrogen",
            "state": "gas",
        },
        "O": { // Oxygen
            "name":  "Oxygen",
            "state": "gas",
        },
        "F": { // Fluorine
            "name":  "Fluorine",
            "state": "gas",
        },
        "Ne": { // Neon
            "name":  "Neon",
            "state": "gas",
        },
    }

    // Lookup an element (Lithium) from the map
    if el, ok := elements["Li"]; ok { // Check if "Li" exists in the map
        fmt.Println(el["name"], el["state"]) // Output: Lithium solid
    }
}

/*
✅ Let’s walk through the program:

1️⃣ **We define a nested map** where:
   - The **keys** (e.g., "H", "Li", "O") represent **element symbols**.
   - The **values** are another **map**, containing:
     - `"name"` → The full name of the element.
     - `"state"` → The state of the element at room temperature.

2️⃣ **We attempt to look up an element in the map**:
   - The program checks if `"Li"` (Lithium) exists in the `elements` map.
   - If it **exists**, it prints both the **name** and **state**.

3️⃣ **Key Benefits of Using a Nested Map**:
   - ✅ **Fast Lookup** → O(1) time complexity for fetching elements.
   - ✅ **Well-Organized** → Keeps related properties together.
   - ✅ **Scalable** → Easily extendable to include more elements or properties (e.g., atomic number).

4️⃣ **Expected Output**:
- The lookup is successful because "Li" **exists in the map**.
- The program prints `"Lithium solid"`, indicating **the name and its standard state**.
*/
