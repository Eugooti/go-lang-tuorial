package arrays

import (
	"fmt"
)

// Arrays, Slices, and Maps in Go

func main() {
	// 1. Arrays in Go:
	// An array is a fixed-size collection of elements of the same type.
	// The size of an array is part of its type and cannot be changed once declared.

	var arr [5]int // Declaring an array of integers with a length of 5
	arr[0] = 10    // Assigning a value to an index
	arr[1] = 20
	fmt.Println("Array:", arr) // Output: [10 20 0 0 0]

	// Key Insights about Arrays:
	// - Fixed size: Arrays have a predefined size that cannot be modified.
	// - Memory efficient: They are stored contiguously in memory, making them fast to access.
	// - Limited flexibility: Due to the fixed size, they are less flexible compared to slices.

	// 2. Slices in Go:
	// A slice is a dynamically-sized, flexible view into the elements of an array.
	// Slices are much more common than arrays in Go due to their flexibility.

	var slice []int                   // Declaring a slice of integers
	slice = append(slice, 10, 20, 30) // Appending values to the slice
	fmt.Println("Slice:", slice)      // Output: [10 20 30]

	// Creating a slice with predefined values
	slice2 := []string{"Go", "Python", "JavaScript"}
	fmt.Println("Slice 2:", slice2)

	// Slicing an Array to Create a Slice:
	subSlice := arr[1:3]                           // Slicing creates a view from the array 'arr'
	fmt.Println("Sub-slice from Array:", subSlice) // Output: [20 0]

	// Key Insights about Slices:
	// - Dynamic size: Slices can grow and shrink as needed using built-in functions like `append()`.
	// - Reference type: Slices are references to arrays. Modifying the slice affects the original array.
	// - Slicing syntax: Slices use a `low:high` syntax to select a range of elements.
	// - Capacity and length: Slices have both `len` (length) and `cap` (capacity) properties.

	// 3. Maps in Go:
	// A map is a collection of key-value pairs, where each key is unique.
	// Maps provide fast lookups, additions, and deletions based on keys.

	// Declaring and Initializing a Map:
	grades := make(map[string]int) // Creating an empty map using `make()`
	grades["Alice"] = 90           // Adding key-value pairs to the map
	grades["Bob"] = 85
	grades["Charlie"] = 92
	fmt.Println("Grades Map:", grades) // Output: map[Alice:90 Bob:85 Charlie:92]

	// Accessing Map Elements:
	aliceGrade := grades["Alice"] // Retrieving value by key
	fmt.Println("Alice's Grade:", aliceGrade)

	// Checking if a Key Exists in a Map:
	_, exists := grades["Dan"] // Checking for non-existing key
	if !exists {
		fmt.Println("Dan's grade is not found in the map.")
	}

	// Deleting a Key-Value Pair from a Map:
	delete(grades, "Bob")                      // Removing "Bob" from the map
	fmt.Println("Updated Grades Map:", grades) // Output: map[Alice:90 Charlie:92]

	// Key Insights about Maps:
	// - Unordered: Maps do not maintain the order of elements.
	// - Fast lookups: Maps provide average O(1) time complexity for lookups, inserts, and deletes.
	// - Keys and values: Keys must be of a type that supports equality (e.g., integers, strings, etc.).
	// - Not thread-safe: Maps are not safe for concurrent use. Synchronization is needed if using across goroutines.

	// Choosing Between Arrays, Slices, and Maps:
	// - Use **arrays** when the size is fixed and known at compile time for memory efficiency.
	// - Use **slices** when you need dynamic size and flexibility in data structures.
	// - Use **maps** when you need fast lookups, inserts, and deletes with unique keys.
}
