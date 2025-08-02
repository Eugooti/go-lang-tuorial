package pointers

import "fmt"

// Main function - the entry point of the program
func main() {
	// **Declaring a Pointer**
	// Declaring an integer variable and a pointer to that integer.
	var x int = 10
	var p *int // `*int` denotes that `p` is a pointer to an integer

	// Assigning the address of `x` to `p`
	p = &x // The `&` operator is used to get the address of a variable
	fmt.Println("Value of x:", x)
	fmt.Println("Address of x:", p)

	// **Dereferencing a Pointer**
	// Dereferencing `p` gives us the value stored at the address.
	fmt.Println("Value pointed to by p:", *p) // The `*` operator is used to dereference a pointer

	// **Modifying the Value via a Pointer**
	// Modifying `x` through the pointer `p`
	*p = 20 // This changes the value of `x` to 20
	fmt.Println("New value of x:", x)

	// **Pointers with Functions**
	// Pointers are often used with functions to modify the argument's value.
	y := 5
	fmt.Println("Value of y before calling changeValue:", y)
	changeValue(&y) // Passing the address of `y` to the function
	fmt.Println("Value of y after calling changeValue:", y)

	// **Nil Pointers**
	// Pointers can be nil, meaning they don't point to any memory location.
	var ptr *int
	fmt.Println("Value of nil pointer:", ptr)
	if ptr == nil {
		fmt.Println("Pointer is nil, no memory location is assigned yet.")
	}
}

// **Function that Accepts a Pointer**
// Function `changeValue` takes a pointer to an integer and modifies the value at that address.
func changeValue(val *int) {
	*val = 100 // Changes the value at the memory location pointed to by `val`
}
