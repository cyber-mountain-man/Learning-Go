package main

import (
	"fmt"        // Import fmt package for printing output
	"hash/crc32" // Import hash/crc32 for computing CRC32 hashes
	"io"         // Import io package for copying file data
	"os"         // Import os package for file operations
)

// getHash calculates the CRC32 hash of a given file
func getHash(filename string) (uint32, error) {
	// Open the file in read mode
	f, err := os.Open(filename)
	if err != nil {
		return 0, err // Return error if file can't be opened
	}
	// Ensure the file is closed when the function exits
	defer f.Close()

	// Create a CRC32 hasher
	h := crc32.NewIEEE()

	// Copy file content into the hasher
	// io.Copy(dst, src) reads data from src (file) and writes it to dst (hasher)
	_, err = io.Copy(h, f)
	if err != nil {
		return 0, err // Return error if file can't be read
	}

	// Return the computed checksum
	return h.Sum32(), nil
}

func main() {
	// Compute hash for "test1.txt"
	h1, err := getHash("test1.txt")
	if err != nil {
		fmt.Println("Error computing hash for test1.txt:", err)
		return
	}

	// Compute hash for "test2.txt"
	h2, err := getHash("test2.txt")
	if err != nil {
		fmt.Println("Error computing hash for test2.txt:", err)
		return
	}

	// Print calculated hashes
	fmt.Println("Hash 1:", h1)
	fmt.Println("Hash 2:", h2)

	// Compare hashes and print whether the files are identical
	fmt.Println("Are the hashes equal?", h1 == h2)
}

/*
✅ **Let's walk through the program:**

1️⃣ **Opening a File and Calculating a CRC32 Hash**
   - `os.Open(filename)`: Opens a file in read mode.
   - `crc32.NewIEEE()`: Creates a CRC32 hasher.
   - `io.Copy(h, f)`: Reads file content and feeds it into the hasher.
   - `h.Sum32()`: Computes the final hash value.

2️⃣ **Why CRC32?**
   - CRC32 is a **fast and lightweight** hashing algorithm.
   - It’s commonly used for **file integrity verification**.

3️⃣ **Handling Errors Properly**
   - If a file **does not exist**, `os.Open` returns an error.
   - If a file **cannot be read**, `io.Copy` will return an error.

4️⃣ **Comparing File Hashes**
   - If two files have the **same hash**, they are **identical**.
   - If two files have **different hashes**, their contents have changed.

5️⃣ **Expected Output**
   - If `test1.txt` and `test2.txt` are identical:
     ```
     Hash 1: 123456789
     Hash 2: 123456789
     Are the hashes equal? true
     ```
   - If they are different:
     ```
     Hash 1: 987654321
     Hash 2: 123456789
     Are the hashes equal? false
     ```

6️⃣ **Real-World Use Cases**
   ✅ **File Integrity Checks** – Verify if files were modified after download.  
   ✅ **Data Deduplication** – Identify duplicate files quickly.  
   ✅ **Cybersecurity Forensics** – Detect file tampering in malware investigations.  
   ✅ **Backup Validation** – Ensure backup files match original files.

This program provides a **quick, efficient, and reliable** way to detect file differences without needing a full content comparison.
*/ 
