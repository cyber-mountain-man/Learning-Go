package main

import (
	"fmt"
	"os"
)

func main() {
	// Step 1: Open the current directory (".")
	dir, err := os.Open(".") // "." refers to the current working directory
	if err != nil {
		fmt.Println("Error opening directory:", err)
		return // Exit if any error occurs
	}
	// Ensure the directory is closed after we are done.
	defer dir.Close()

	// Step 2: Read all directory entries (-1 means read all)
	fileInfos, err := dir.Readdir(-1) // Retrieves information about all files/folders in the directory
	if err != nil {
		fmt.Println("Error reading directory:", err)
		return // Exit if an error occurs
	}

	// Step 3: Loop through the directory contents and print file/folder names
	for _, fi := range fileInfos {
		fmt.Println(fi.Name()) // Print the name of each file or directory
	}
}

/*
✅ Let's walk through the program:

1️⃣ **Opening the Current Directory (`os.Open(".")`)**  
   - `os.Open(".")` opens the **current directory**.
   - If there’s an error (e.g., permission issues), it prints an error message and exits.

2️⃣ **Ensuring the Directory is Closed (`defer dir.Close()`)**  
   - `defer` ensures that the directory is **closed automatically** when `main()` finishes execution.
   - This **prevents resource leaks**.

3️⃣ **Reading Directory Contents (`dir.Readdir(-1)`)**  
   - `dir.Readdir(-1)` retrieves **all entries** (files and folders) inside the directory.
   - Returns a **slice of `os.FileInfo`** objects, containing metadata about each file.

4️⃣ **Looping Through Entries and Printing Names (`fi.Name()`)**  
   - The program **iterates over the directory entries** and prints their names.

5️⃣ **Expected Output (Example)**
   If the current directory contains:
		main.go notes.txt my_folder
   The output will be:
		main.go notes.txt my_folder


6️⃣ **Why Use This Approach?**  
- ✅ **Lists all files & folders dynamically** (no hardcoded file names).  
- ✅ **Works in any directory**—it scans based on where the program runs.  
- ✅ **Useful for building file explorers, log scanners, or security monitoring tools.**  

7️⃣ **Real-World Applications**  
- 🔹 **File listing CLI tools**  
- 🔹 **Security monitoring (detecting new/unexpected files)**  
- 🔹 **Log file management scripts**  
- 🔹 **Backup and file processing automation**  
*/
