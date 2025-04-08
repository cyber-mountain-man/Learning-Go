# Concurrent Cashiers – A Go Concurrency Demo

This Go program demonstrates **concurrency using goroutines and sync.WaitGroup** by simulating a real-world scenario: **10 grocery store cashiers** each serving **10 customers** concurrently.

---

## 🧠 What This Program Teaches

- How to run multiple tasks at the same time using `goroutines`
- How to wait for all tasks to finish using `sync.WaitGroup`
- How to simulate real-world timing delays using `math/rand` and `time.Sleep`
- How Go handles lightweight concurrency without using threads manually

---

## 🛒 Real-World Analogy

Imagine 10 cashiers in a grocery store. Each cashier helps 10 customers. The time it takes to help each customer varies. All cashiers are working **at the same time**, and the program waits until **everyone is finished** before closing the store.

---

## 🧬 How It Works

- `main()` starts 10 cashier goroutines.
- Each goroutine runs `serveCustomers()`, simulating 10 customers.
- Random delays are used to simulate different service durations.
- A `sync.WaitGroup` ensures the program **waits for all goroutines to finish** before exiting.

---

## 🏃‍♂️ How to Run

Make sure you have Go installed (version 1.20 or newer recommended).

1. Clone the repository or download the `.go` file.
2. Open your terminal and navigate to the project folder.
3. Run the program:

```bash
go run main.go
```

You should see output like:

```
Cashier 1 is serving customer 1
Cashier 2 is serving customer 1
...
🛒 All customers have been served.
```

---

## 📂 File Structure

```
concurrent_cashiers/
├── main.go       # The Go source file
└── README.md     # This explanation file
```

---

## 📘 Key Go Concepts Used

| Feature            | Description |
|--------------------|-------------|
| `goroutines`       | Lightweight, concurrent execution of functions |
| `sync.WaitGroup`   | Tracks and waits for completion of goroutines |
| `defer`            | Ensures `wg.Done()` is called even if there's an error |
| `rand.Intn()`      | Generates random delays to mimic work time |
| `time.Sleep()`     | Simulates each cashier working on a customer |

---

## ✅ Next Steps (Optional Enhancements)

- Log customer start/end service times
- Add input to set number of cashiers/customers
- Visualize activity using a simple web interface or terminal animation

---

## 🔐 License

This project is open for learning purposes. No license restrictions.
```