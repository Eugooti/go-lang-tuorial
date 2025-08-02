package goRuntime

import (
	"fmt"
	"sync"
	"time"
)

// Worker function that simulates some processing
func worker(id int, jobs <-chan int, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for job := range jobs {
		fmt.Printf("Worker %d started job %d\n", id, job)
		time.Sleep(time.Second) // Simulate time-consuming work
		fmt.Printf("Worker %d finished job %d\n", id, job)
		results <- job * 2 // Send the result back
	}
}

func main() {
	// Create channels for jobs and results
	const numJobs = 5
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	// WaitGroup to wait for all goroutines to finish
	var wg sync.WaitGroup

	// Create a pool of 3 workers
	for w := 1; w <= 3; w++ {
		wg.Add(1)
		go worker(w, jobs, results, &wg)
	}

	// Send jobs to workers
	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs) // Close the jobs channel to signal no more jobs

	// Wait for all workers to finish
	wg.Wait()

	// Collect the results
	close(results)
	for result := range results {
		fmt.Println("Result:", result)
	}
}
