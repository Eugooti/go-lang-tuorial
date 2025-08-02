package concurrency

import (
	"fmt"
	"time"
)

// Concurrency in Go: Goroutines and Channels

// 1. Goroutines:
// A Goroutine is a lightweight thread managed by the Go runtime.
// Goroutines are created using the `go` keyword followed by a function call.

func printNumbers() {
	for i := 1; i <= 5; i++ {
		fmt.Printf("Number: %d\n", i)
		time.Sleep(500 * time.Millisecond) // Simulate some work with sleep
	}
}

func printLetters() {
	for c := 'A'; c <= 'E'; c++ {
		fmt.Printf("Letter: %c\n", c)
		time.Sleep(700 * time.Millisecond) // Simulate some work with sleep
	}
}

func main() {
	// 2. Starting Goroutines:
	// By using the `go` keyword, both functions will run concurrently (in separate Goroutines).
	go printNumbers()
	go printLetters()

	// Main function continues to execute immediately after starting Goroutines.
	// We add a sleep here to allow Goroutines to complete their execution.
	time.Sleep(4 * time.Second)

	fmt.Println("Main function ends.")
	// Note: In a real-world application, `sync.WaitGroup` or channels would be used to wait for Goroutines to complete.

	// 3. Channels:
	// Channels are a way for Goroutines to communicate with each other and synchronize their execution.

	// Creating a channel of type 'int'
	numbers := make(chan int)

	// Function to generate numbers and send them to the channel
	go func() {
		for i := 1; i <= 5; i++ {
			numbers <- i // Sending data to the channel
			time.Sleep(500 * time.Millisecond)
		}
		close(numbers) // Close the channel when done
	}()

	// Receiving data from the channel
	for num := range numbers {
		fmt.Printf("Received number from channel: %d\n", num)
	}

	// 4. Using Channels for Synchronization:
	// Channels can also be used to signal completion.

	done := make(chan bool) // Create a 'done' channel

	// Start a Goroutine to perform a task
	go func() {
		fmt.Println("Performing a task...")
		time.Sleep(2 * time.Second) // Simulate a task
		done <- true                // Signal completion
	}()

	<-done // Wait for the signal before proceeding
	fmt.Println("Task completed. Exiting program.")
}
