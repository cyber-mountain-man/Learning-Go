package main

import (
	"container/list" // ✅ Import Go's built-in doubly linked list package
	"fmt"
	"os/exec" // ✅ Allows execution of system commands
	"time"     // ✅ Used for adding delays between task executions
)

// ✅ Define Task type as an alias for function with no parameters
type Task func()

func main() {
	// ✅ Create a new task queue using a linked list
	taskQueue := list.New()

	// ✅ Define real cybersecurity tasks using Task type

	// 🔍 Task 1: Run an Nmap scan
	task1 := Task(func() {
		fmt.Println("🔍 Running Nmap scan...")
		// Execute Nmap to scan ports 22, 80, and 443 on scanme.nmap.org
		cmd := exec.Command("nmap", "-p", "22,80,443", "scanme.nmap.org")
		output, _ := cmd.Output() // Capture command output
		fmt.Println(string(output))
	})

	// 📊 Task 2: Check system logs
	task2 := Task(func() {
		fmt.Println("📊 Checking system logs...")
		// Windows alternative for `cat /var/log/syslog`
		cmd := exec.Command("type", "C:\\Windows\\System32\\LogFiles\\some-log.log")
		output, _ := cmd.Output()
		fmt.Println(string(output))
	})

	// 🛡 Task 3: Perform a WHOIS lookup
	task3 := Task(func() {
		fmt.Println("🛡 Running WHOIS lookup...")
		cmd := exec.Command("whois", "example.com") // Lookup domain information
		output, _ := cmd.Output()
		fmt.Println(string(output))
	})

	// ✅ Store functions correctly in the linked list
	taskQueue.PushBack(task1)
	taskQueue.PushBack(task2)
	taskQueue.PushBack(task3)

	fmt.Println("📌 Executing Cybersecurity Task Queue...")

	// ✅ Process and execute each function dynamically
	for e := taskQueue.Front(); e != nil; e = e.Next() {
		task, ok := e.Value.(Task) // ✅ Ensure correct type conversion
		if !ok {
			fmt.Println("Error: Unable to convert to Task")
			continue
		}
		task()                      // ✅ Execute the function
		time.Sleep(2 * time.Second)  // Simulate delay between tasks
	}

	// ✅ Clear queue after execution
	taskQueue.Init()
	fmt.Println("✅ All cybersecurity tasks completed, queue cleared!")
}

/* 📌 Walkthrough of the Cybersecurity Task Queue
1️⃣ Using container/list for Task Management

A linked list is used to dynamically queue tasks (instead of hardcoding sequential execution).
Each task is a function that gets executed dynamically.
2️⃣ Defining Cybersecurity Tasks as Functions

Task 1: Nmap Scan 🔍
Runs nmap to scan for open ports.
Task 2: Check System Logs 📊
Uses type (Windows) to read log files.
Task 3: WHOIS Lookup 🛡
Queries domain information.
3️⃣ Adding Tasks to the Queue

Tasks are pushed into the linked list using PushBack().
4️⃣ Processing the Queue

The list is traversed using taskQueue.Front(), executing each stored function dynamically.
task() // Executes function calls each function stored in the list.
5️⃣ Handling Type Assertions for Safety

task, ok := e.Value.(Task) ensures the stored value is a valid function before executing it.
6️⃣ Adding Execution Delays (time.Sleep)

Simulates real-world cybersecurity scan delays.
7️⃣ Clearing the Queue After Execution

taskQueue.Init() resets the linked list.


📌 Expected Output
📌 Executing Cybersecurity Task Queue...
🔍 Running Nmap scan...
(Displays Nmap scan results)
📊 Checking system logs...
(Displays contents of log file)
🛡 Running WHOIS lookup...
(Displays WHOIS domain information)
✅ All cybersecurity tasks completed, queue cleared!

📌 Why Use a Linked List for Cybersecurity Automation?
✔ Dynamically Execute Tasks → No need to hardcode a task order.
✔ Modular & Expandable → New tasks can be added without changing core logic.
✔ Real-World Application → Can be extended for network scanning, intrusion detection, and log analysis.

This linked list-based cybersecurity task execution framework provides a flexible way to automate and manage security tasks dynamically.
*/