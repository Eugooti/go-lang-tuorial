package main

import (
	"fmt"
	"goTutorial/mathutil"
)

func main() {
	// Using functions from the imported package
	sum := mathutil.Add(10, 5)
	fmt.Printf("Sum: %d\n", sum)

	diff := mathutil.Subtract(10, 5)
	fmt.Printf("Difference: %d\n", diff)
}
