package functions

import (
	"errors"
	"fmt"
)

// Main function - the entry point of the program
func main() {
	// **Basic Functions**
	// Calling a function that adds two numbers
	sum := add(5, 3)
	fmt.Println("Sum:", sum)

	// **Functions with Multiple Return Values**
	// Calling a function that returns multiple values (result and error)
	result, err := divide(10, 2)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result of division:", result)
	}

	// **Named Return Values**
	// Calling a function that uses named return values
	area := rectangleArea(5, 7)
	fmt.Println("Area of rectangle:", area)

	// **Variadic Functions**
	// Calling a variadic function that sums any number of integers
	total := sumAll(1, 2, 3, 4, 5)
	fmt.Println("Total Sum of all numbers:", total)

	// You can also pass a slice to a variadic function using the `...` operator
	nums := []int{10, 20, 30, 40, 50}
	totalFromSlice := sumAll(nums...)
	fmt.Println("Total Sum from slice:", totalFromSlice)
}

// **Basic Function Declaration**
// Function `add` takes two integers and returns their sum.
func add(a int, b int) int {
	return a + b
}

// **Functions with Multiple Return Values**
// Function `divide` takes two floats and returns the result and an error if any.
func divide(a, b float64) (float64, error) {
	// Check if the divisor is zero to avoid division by zero error.
	if b == 0 {
		return 0, errors.New("cannot divide by zero")
	}
	// If no error, return the result and nil for error.
	return a / b, nil
}

// **Named Return Values**
// Function `rectangleArea` calculates the area of a rectangle.
// It uses named return values which are initialized to zero and can be returned without explicitly using `return` values.
func rectangleArea(length, width float64) (area float64) {
	area = length * width // `area` is already declared as the return value
	return                // Returns the named return variable `area`
}

// **Variadic Functions**
// A variadic function is a function that can accept a variable number of arguments.
// Function `sumAll` takes any number of integers and returns their sum.
func sumAll(numbers ...int) int {
	sum := 0
	for _, num := range numbers { // `range` iterates over all arguments passed to the function
		sum += num
	}
	return sum
}
