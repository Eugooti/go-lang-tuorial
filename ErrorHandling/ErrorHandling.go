package ErrorHandling

import (
	"errors" // Standard library package for error handling
	"fmt"    // Standard library package for formatted I/O
	"os"     // Standard library package for interacting with the operating system
)

// Main function - the entry point of the program
func main() {
	// **Handling Errors from Standard Library Functions**
	// Attempt to open a non-existent file to generate an error.
	_, err := os.Open("non_existent_file.txt")
	if err != nil { // Check if an error occurred
		fmt.Println("Error opening file:", err)
	}

	// **Creating and Returning Custom Errors**
	// Calling a function that returns a custom error
	err = performCalculation(-1)
	if err != nil { // Check if an error occurred
		fmt.Println("Custom Error:", err)
	}

	// **Error Wrapping with fmt.Errorf**
	// `fmt.Errorf` can be used to wrap errors with more context.
	err = performTask("invalid_task")
	if err != nil {
		fmt.Println("Error with additional context:", err)
	}

	// **Using the `errors.Is` and `errors.As` Functions**
	// Demonstrates how to check for specific error types and values.
	err = checkEvenNumber(3)
	if err != nil {
		if errors.Is(err, ErrNotEven) { // Check if the error is of a specific type
			fmt.Println("Specific Error:", err)
		}
	}
}

// **Custom Error Creation**
// You can create custom errors using the `errors.New` function or by implementing the `error` interface.
var ErrNotEven = errors.New("number is not even")

// **Function Returning an Error**
// Function `checkEvenNumber` returns an error if the number is not even.
func checkEvenNumber(n int) error {
	if n%2 != 0 {
		return ErrNotEven // Returns the custom error if the number is not even
	}
	return nil // Returns `nil` if no error
}

// **Returning Errors from Functions**
// Function `performCalculation` checks if a value is negative and returns an error if so.
func performCalculation(value int) error {
	if value < 0 {
		// Creates and returns a custom error
		return errors.New("cannot perform calculation on a negative number")
	}
	fmt.Println("Calculation performed successfully.")
	return nil
}

// **Wrapping Errors with Context**
// `fmt.Errorf` is used to wrap an error with additional context.
func performTask(task string) error {
	if task != "valid_task" {
		// Use `fmt.Errorf` to wrap the error with more context.
		return fmt.Errorf("task %s failed: %w", task, errors.New("invalid task provided"))
	}
	fmt.Println("Task performed successfully.")
	return nil
}
