package pointers

import "fmt"

// Main function - the entry point of the program
func main() {
	// **Basic Pointers**
	// Declare a variable
	x := 10
	// Print the value and the memory address
	fmt.Println("Value of x:", x)
	fmt.Println("Address of x:", &x)

	// Create a pointer that stores the address of x
	var p *int = &x
	fmt.Println("Pointer p stores the address of x:", p)

	// **Dereferencing Pointers**
	// Dereferencing the pointer to get the value at the memory address
	fmt.Println("Value at the address stored in pointer p:", *p)

	// **Modifying Values Using Pointers**
	// Modifying the value of x through the pointer p
	*p = 20
	fmt.Println("New value of x after modification through pointer p:", x)

	// **Pointers and Functions**
	// Passing pointers to functions
	num := 5
	fmt.Println("Value before doubling:", num)
	doubleValue(&num) // Pass the address of num to the function
	fmt.Println("Value after doubling:", num)

	// **Nil Pointers**
	// A nil pointer does not point to any memory location
	var ptr *int
	fmt.Println("Nil pointer:", ptr)

	// **Pointer to Pointer**
	// A pointer can also point to another pointer
	ptr1 := &x
	ptr2 := &ptr1
	fmt.Println("Value of x through double pointer:", **ptr2)
}

// **Function that takes a pointer as a parameter**
// Function `doubleValue` takes a pointer to an integer and modifies the value at that address.
func doubleValue(n *int) {
	*n = *n * 2 // Dereference the pointer and modify the value
}
