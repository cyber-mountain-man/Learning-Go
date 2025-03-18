package main

import (
	"fmt"
	"hash/crc32" // Import the CRC32 hashing package
)

func main() {
	// ✅ Create a new CRC32 hasher using IEEE polynomial
	h := crc32.NewIEEE()

	// ✅ Write data to the hasher (converts string to bytes)
	h.Write([]byte("test"))

	// ✅ Compute the CRC32 checksum from the written data
	v := h.Sum32()

	// ✅ Print the checksum (output varies based on input)
	fmt.Println(v) // Example output: 3632233996
}

/*
✅ Let's walk through the program:

1️⃣ **Importing the `crc32` Package**
   - The `hash/crc32` package provides the **Cyclic Redundancy Check (CRC32) hashing algorithm**.
   - CRC32 is commonly used for **checksums, error detection, and data integrity verification**.

2️⃣ **Creating a Hasher**
   - `crc32.NewIEEE()` initializes a **CRC32 hasher** using the **IEEE polynomial**.
   - The IEEE variant is commonly used for network checksums.

3️⃣ **Writing Data to the Hasher**
   - `h.Write([]byte("test"))` converts the string `"test"` into bytes and writes it to the hasher.
   - The hasher processes the input data **incrementally**.

4️⃣ **Generating the CRC32 Checksum**
   - `h.Sum32()` computes the **final CRC32 checksum**.
   - The checksum is a **fixed 32-bit unsigned integer**.

5️⃣ **Printing the Hash**
   - The output (e.g., `3632233996`) represents the **CRC32 checksum** of `"test"`.

6️⃣ **Expected Output**
		3632233996

- The **output remains consistent** for the same input.

7️⃣ **Real-World Use Cases**
- ✅ **File integrity verification** (checking for corruption).
- ✅ **Fast hash-based indexing** (hash tables, data deduplication).
- ✅ **Network transmission checks** (verifying packet correctness).
- ✅ **Error detection** in data storage and communications.

🔹 **Note:** CRC32 is **not cryptographically secure**, so it's **not recommended for security-sensitive applications**.
*/
