# Multi-threaded Data Processing System (Go)

This project contains an implementation of a multi-threaded data processing system in **Go**. The system simulates a producer-consumer model where multiple worker goroutines retrieve tasks from a shared queue, process them, and store the results.

### Overview

The Go implementation follows an idiomatic approach to concurrency, utilizing goroutines and channels as its fundamental building blocks. This model emphasizes communication over shared memory, which often leads to simpler and more robust concurrent code.

### Key Concepts

* **Concurrency:**
  * **`goroutines`**: Lightweight, concurrently executing functions are used for each worker, running on a minimal stack.
  * **`channels`**: The primary means of communication and synchronization. A buffered channel acts as the thread-safe task queue, and another is used for collecting results. Channels provide built-in synchronization, eliminating the need for explicit locks in many cases.
  * **`sync.WaitGroup`**: The main function uses a `WaitGroup` to launch the worker goroutines and then waits for all of them to complete their execution, ensuring a clean program exit.

* **Error Handling:**
  * Go functions that can fail typically return a second value of type `error`. The code checks for a non-`nil` error value to handle failures gracefully.
  * The `defer` keyword is used to ensure that resources, such as file handles, are properly closed after their use, regardless of whether an error occurs.

### How to Run on macOS

1. Ensure you have Go installed. You can check by running `go version` in your Terminal.
2. Save the Go code as `main.go`.
3. Open Terminal and navigate to the directory where the file is saved.
4. Run the program directly: `go run main.go`
