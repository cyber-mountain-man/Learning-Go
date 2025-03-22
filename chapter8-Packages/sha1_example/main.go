package main

import (
	"crypto/sha1" // Import crypto/sha1 for computing SHA-1 hashes
	"fmt"         // Import fmt package for formatted output
)

func main() {
	// ✅ Create a new SHA-1 hasher
	h := sha1.New()

	// ✅ Write the data ("test") into the hasher
	h.Write([]byte("test"))

	// ✅ Compute the final hash (SHA-1 checksum)
	bs := h.Sum(nil) // Sum(nil) returns the complete hashed output

	// ✅ Print the hash in raw byte format
	fmt.Println(bs) // Output: [random byte values]
}

/*
✅ **Let's walk through the program:**

1️⃣ **Creating a SHA-1 Hasher**
   - `sha1.New()` initializes a new SHA-1 hasher.
   - `h.Write([]byte("test"))` feeds the **"test"** string into the hasher.

2️⃣ **Computing the Hash**
   - `h.Sum(nil)` finalizes the hash and returns the **SHA-1 digest** in byte format.

3️⃣ **Expected Output**
   - The output is a **byte slice** representation of the SHA-1 hash, like:
     ```
     [227 176 196 66 152 252 28 20 154 251 244 200 153 111 185 36 39 174 65 228]
     ```
   - To print the hash in **hex format**, use:
     ```go
     fmt.Printf("%x\n", bs)
     ```
     This would return:
     ```
     a94a8fe5ccb19ba61c4c0873d391e987982fbbd3
     ```

4️⃣ **Why SHA-1?**
   - SHA-1 produces a **160-bit (20-byte) hash**.
   - Used in **file verification, digital signatures, and data integrity**.

5️⃣ **Important Notes on Security**
   - **SHA-1 is considered weak** due to known **collision attacks**.
   - **Modern alternatives**: Use **SHA-256** or **SHA-3** for security-critical applications.

6️⃣ **Real-World Use Cases**
   ✅ **Verifying downloaded files** (checking if data was altered).  
   ✅ **Password hashing (though SHA-256 or bcrypt is better)**.  
   ✅ **Digital signatures and authentication tokens**.  
   ✅ **Data integrity checks in databases and version control systems**.

 While SHA-1 is fast, **it should not be used for cryptographic security** due to its vulnerabilities. Instead, prefer **SHA-256 or SHA-3** in modern applications.
*/
