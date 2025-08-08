// main.go
package main

import (
	"fmt"
	"log"
	"math/rand"
	"sync"
	"time"
)

// Task represents a unit of work.
type Task struct {
	ID int
}

// Worker function for a goroutine.
func worker(id int, tasks <-chan Task, results chan<- string, wg *sync.WaitGroup) {
	// The defer statement ensures wg.Done() is called when the function returns.
	defer wg.Done()
	log.Printf("Worker %d starting...", id)

	// The for-range loop on the channel will automatically terminate when the channel is closed.
	for task := range tasks {
		// Simulate computational work with a random delay.
		delay := time.Duration(rand.Intn(500)) * time.Millisecond
		time.Sleep(delay)

		// Process the task and send the result to the results channel.
		result := fmt.Sprintf("Worker %d processed Task-%d in %dms.", id, task.ID, delay.Milliseconds())
		log.Println(result)
		results <- result
	}

	log.Printf("Worker %d finished.", id)
}

func main() {
	const (
		numWorkers = 4
		numTasks   = 20
	)

	// Create a buffered channel for tasks. This acts as our shared queue.
	tasks := make(chan Task, numTasks)

	// Create a buffered channel for results. This is another form of a shared resource.
	results := make(chan string, numTasks)

	// WaitGroup to wait for all worker goroutines to finish.
	var wg sync.WaitGroup

	log.Println("Main: Starting", numWorkers, "worker goroutines.")
	// Launch worker goroutines.
	for i := 1; i <= numWorkers; i++ {
		wg.Add(1) // Increment the WaitGroup counter for each worker.
		go worker(i, tasks, results, &wg)
	}

	log.Println("Main: Populating the task channel with", numTasks, "tasks.")
	// Send tasks to the channel.
	for i := 0; i < numTasks; i++ {
		tasks <- Task{ID: i}
	}
	
	// Close the tasks channel to signal that no more tasks will be sent.
	// This is crucial for the for-range loops in the workers to terminate.
	close(tasks)

	log.Println("Main: Waiting for all workers to finish...")
	// Wait for all worker goroutines to complete.
	wg.Wait()

	// Since we know all workers have finished, we can now safely close the results channel.
	close(results)

	// Collect and print all results.
	fmt.Println("\nMain: All workers have finished. Final results:")
	for result := range results {
		fmt.Println(result)
	}

	fmt.Println("\nMain: Program terminated successfully.")
}

