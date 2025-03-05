package main

import "fmt"

func main() {
    // Creating a map with string keys (element symbols) and string values (element names)
    elements := make(map[string]string)

    // Populating the map with element symbols and their corresponding names
    elements["H"] = "Hydrogen"
    elements["He"] = "Helium"
    elements["Li"] = "Lithium"
    elements["Be"] = "Beryllium"
    elements["B"] = "Boron"
    elements["C"] = "Carbon"
    elements["N"] = "Nitrogen"
    elements["O"] = "Oxygen"
    elements["F"] = "Fluorine"
    elements["Ne"] = "Neon"

    // Retrieving and printing the value associated with the key "Li"
    fmt.Println(elements["Li"]) // Output: Lithium
}


