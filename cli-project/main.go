package main

import (
	"eugooti.com/cli-project/calculator"
	"eugooti.com/cli-project/tenants"
	"fmt"
)

func main() {
	factorial, _ := calculator.Factorial(0)
	fmt.Println(factorial)

	Tenants := tenants.GetTenants()
	var (
		firstName   string
		lastName    string
		gender      int
		houseNumber int
	)

	fmt.Println("First Name: ")
	fmt.Scan(&firstName)
	fmt.Println("Last Name: ")
	fmt.Scan(&lastName)
	fmt.Println("Gender (1-male, 2-female): ")
	fmt.Scan(&gender)
	fmt.Println("House Number: ")
	fmt.Scan(&houseNumber)

	var Gender string
	switch gender {
	case 1:
		Gender = "male"
	case 2:
		Gender = "female"
	default:
		Gender = "Wrong gender"
	}

	var newTenant tenants.Tenant = tenants.Tenant{
		Firstname:   firstName,
		Lastname:    lastName,
		Gender:      Gender,
		HouseNumber: houseNumber,
	}

	addTenant := tenants.AddTenant(newTenant)

	fmt.Println(addTenant)
	fmt.Println(Tenants)

	for _, tenant := range Tenants {
		if tenant.Gender == "male" {
			fmt.Println(tenant.Firstname)
		}
	}
}
