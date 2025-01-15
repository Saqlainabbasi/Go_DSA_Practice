package question

import "fmt"

// factorial
func Factorial(val int) int {
	if val <= 1 {
		return val
	}
	return val * Factorial(val-1)
}

// Fibonacci Sequence
func Fibonacci(val int) int {
	//val is less then 1
	if val <= 1 {
		return val
	}
	fmt.Println(val)
	//val-1 will return the previous no then the val
	//val-2 will reutrn the 2 previous no then the val
	return Fibonacci(val-1) + Fibonacci(val-2)
}
