package loops

import (
	"fmt"
)

func main() {
	// **Control Structures in Go**

	// 1. IF-ELSE STATEMENTS
	// Go uses straightforward `if` and `else` statements similar to other languages.
	// However, Go does not require parentheses around conditions, but braces `{}` are mandatory.
	var number int = 10

	if number%2 == 0 { // Checks if `number` is even
		fmt.Println("Number is even")
	} else { // Executes if `number` is odd
		fmt.Println("Number is odd")
	}

	// **If with Initialization Statement**
	// In Go, you can include an initialization statement in an `if` statement.
	// This allows you to declare and initialize a variable only for that block scope.
	if num := 5; num > 0 { // Declares and initializes `num` in the if statement
		fmt.Println("Num is positive")
	} else {
		fmt.Println("Num is negative or zero")
	}

	// 2. SWITCH STATEMENTS
	// Go's `switch` is more flexible than in many languages. It does not require `break` statements;
	// it automatically breaks after a case is executed, unless `fallthrough` is used.
	day := 3

	switch day {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
		fallthrough // `fallthrough` forces the execution to continue to the next case
	case 4:
		fmt.Println("Thursday")
	default: // `default` is optional and executes if no other case matches
		fmt.Println("Invalid day")
	}

	// **Switch with Initialization**
	// Like `if`, `switch` can also have an initialization statement.
	switch today := "Sunday"; today {
	case "Saturday", "Sunday":
		fmt.Println("Weekend")
	default:
		fmt.Println("Weekday")
	}

	// 3. FOR LOOPS
	// Go only has one looping construct: the `for` loop. There is no `while` or `do-while` loop.
	// The `for` loop in Go is versatile and can be used in different forms.

	// **Basic For Loop**
	for i := 0; i < 5; i++ { // Initialization, condition, and increment/decrement in one line
		fmt.Println("Basic For Loop, iteration:", i)
	}

	// **For Loop as a While Loop**
	// If you omit the initialization and post statements, it acts like a `while` loop.
	j := 0
	for j < 3 { // Only the condition remains; acts like a `while` loop
		fmt.Println("While-like For Loop, iteration:", j)
		j++
	}

	// **Infinite Loop**
	// Omitting the condition creates an infinite loop. Use `break` to exit.
	// This loop would run forever without the `break` statement.
	counter := 0
	for { // Infinite loop
		fmt.Println("Infinite Loop, iteration:", counter)
		counter++
		if counter == 2 {
			break // Exits the loop when counter equals 2
		}
	}

	// **Range in For Loops**
	// `for` loops can iterate over slices, arrays, maps, strings, and channels using `range`.
	// `range` provides both index and value for slices/arrays and key-value pairs for maps.
	nums := []int{2, 4, 6, 8}
	for index, value := range nums { // Iterates over the slice `nums`
		fmt.Printf("Index: %d, Value: %d\n", index, value)
	}

	// You can use an underscore `_` to ignore the index or value if not needed.
	for _, value := range nums { // Ignores the index and only uses the value
		fmt.Printf("Value: %d\n", value)
	}

	// Go does not support `do-while` or `while` directly but can simulate them using
	//the `for` loop.
}
