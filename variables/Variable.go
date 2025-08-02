package variables

import (
	"fmt"
)

// Main function - the entry point of the program
func main() {
	// **Variables in Go**

	// 1. Integer Types: int, int8, int16, int32, int64
	// The `int` type is platform-dependent; it can be either 32 or 64 bits.
	// Use specific sizes like `int32` or `int64` if you need a fixed size.
	var age int = 25
	fmt.Println("Age:", age)

	// 2. Float Types: float32, float64
	// `float32` and `float64` are used for decimal numbers.
	// Default floating-point type is `float64`.
	var pi float64 = 3.14159
	fmt.Println("Pi:", pi)

	// 3. String Type: A sequence of characters.
	// Strings are immutable in Go, meaning they cannot be modified after creation.
	var greeting string = "Hello, Golang!"
	fmt.Println("Greeting:", greeting)

	// 4. Boolean Type: `bool`
	// Booleans can be either `true` or `false`.
	var isGoFun bool = true
	fmt.Println("Is Go fun?", isGoFun)

	// **Constants in Go**
	// Constants are immutable values which are known at compile time and do not change during program execution.
	const piConstant float64 = 3.14159
	fmt.Println("Constant Pi:", piConstant)

	// Multiple constants can be declared together
	const (
		firstName string = "Eugene"
		lastName  string = "Ochieng"
	)
	fmt.Println("Full Name:", firstName, lastName)

	// **Short Variable Declaration (:=)**
	// This is shorthand for declaring and initializing a variable.
	// Note: This can only be used within functions, not at the package level.
	country := "Kenya" // Type inferred as string
	fmt.Println("Country:", country)

	// **Zero Values in Go**
	// When variables are declared without an initial value, they get a "zero value":
	var uninitializedInt int       // 0
	var uninitializedFloat float64 // 0.0
	var uninitializedString string // ""
	var uninitializedBool bool     // false

	fmt.Println("Uninitialized Int:", uninitializedInt)
	fmt.Println("Uninitialized Float:", uninitializedFloat)
	fmt.Println("Uninitialized String:", uninitializedString)
	fmt.Println("Uninitialized Bool:", uninitializedBool)

	// **Type Conversion**
	// Go is statically typed; you need to explicitly convert types.
	// Example: Convert `int` to `float64`.
	var number int = 42
	var convertedFloat float64 = float64(number) // Explicit conversion
	fmt.Println("Converted Float:", convertedFloat)

	// **Rune Type**
	// `rune` is an alias for `int32` and represents a Unicode code point.
	// Useful when working with non-ASCII characters.
	var char rune = 'A'
	fmt.Printf("Rune: %c, Unicode Code Point: %U\n", char, char)

	// **Byte Type**
	// `byte` is an alias for `uint8`. It is used to represent ASCII characters.
	// Commonly used when dealing with raw data like files or network I/O.
	var letter byte = 'B'
	fmt.Printf("Byte: %c, ASCII Code: %d\n", letter, letter)

	/**
	Insights Included in the Code
	1. Short Variable Declaration (:=): Explained with an example for declaring and
	initializing a variable in one line.
	2. Zero Values: Included an explanation and examples showing the default zero values
	for different types.
	3. Type Conversion: Explained that Go requires explicit type conversion and provided
	an example.
	4. Rune and Byte Types: Added explanations about rune and byte, their purposes, and
	how they differ from other types.
	5. Constants: Explained how constants work and provided examples of multiple constant
	declarations.
	**/
}
