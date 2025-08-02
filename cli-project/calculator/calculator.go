package calculator

import "errors"

func Sum(numbers ...int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func Minus(a int, b int) int {
	return a - b
}

func Multiplication(numbers ...int) int {
	multiple := 0
	for _, number := range numbers {
		multiple *= number
	}
	return multiple

}

func Division(a int, b int) float64 {
	return float64(a / b)
}

func Factorial(a int) (int, error) {
	if a == 0 {
		return 0, errors.New("you need a number greater than 0")
	}
	factorial := 1
	for i := 1; i <= a; i++ {
		factorial *= i
	}
	return factorial, nil
}
