package structs

import (
	"fmt"
)

// Structs and Methods in Go

// 1. Defining a Struct:
// A `struct` is a collection of fields (variables) that are grouped together under a single name.
// Structs are useful for representing data with multiple properties.

type Person struct {
	FirstName string  // Field of type string
	LastName  string  // Field of type string
	Age       int     // Field of type int
	Height    float64 // Field of type float64
}

// 2. Creating a Custom Type:
// You can create a custom type in Go based on existing types, enhancing code readability and creating more meaningful abstractions.

type Meters float64 // Custom type 'Meters' based on float64

// 3. Methods in Go:
// Methods are functions with a special receiver argument. The receiver is a variable that is passed before the function
//name, like a method in OOP languages.

func (p Person) FullName() string { // Method attached to the 'Person' struct
	// Returns the full name by concatenating FirstName and LastName
	return p.FirstName + " " + p.LastName
}

// 4. Method with a Pointer Receiver:
// When we use a pointer receiver (*Person), the method can modify the struct it is called upon.

func (p *Person) IncrementAge() { // Method to increment age, changes will reflect on the original object
	p.Age++ // Increase the age of the person by 1
}

// 5. Function vs. Method:
// Functions are declared outside any struct, while methods are tied to a specific struct type.

func CreatePerson(firstName, lastName string, age int, height Meters) Person { // Function to create a new Person instance
	// Convert custom type 'Meters' to 'float64'
	return Person{
		FirstName: firstName,
		LastName:  lastName,
		Age:       age,
		Height:    float64(height), // Convert custom type to float64
	}
}

func main() {
	// 6. Creating an Instance of a Struct:
	// You can create an instance of a struct by providing values for all fields.
	person1 := Person{
		FirstName: "John",
		LastName:  "Doe",
		Age:       30,
		Height:    1.75,
	}

	// 7. Accessing Struct Fields:
	// You can access and modify struct fields using the dot (.) operator.
	fmt.Println("First Name:", person1.FirstName) // Accessing a struct field

	// 8. Using Struct Methods:
	// Methods can be called on struct instances to perform actions or computations related to that struct.
	fullName := person1.FullName() // Calling a method
	fmt.Println("Full Name:", fullName)

	// 9. Pointer Receivers and Methods:
	// Pointer receivers are used when the method needs to modify the original struct.
	fmt.Println("Age before increment:", person1.Age)
	person1.IncrementAge() // Modifies the original 'person1' struct
	fmt.Println("Age after increment:", person1.Age)

	// 10. Using a Custom Type with Structs:
	// Demonstrating a function that uses a custom type to create a new struct instance.
	person2 := CreatePerson("Jane", "Smith", 25, 1.65)
	fmt.Println("New Person:", person2.FirstName, person2.LastName, "Age:", person2.Age, "Height:", person2.Height)
}
