package main

import "fmt"

func main_as_value() {
	fmt.Printf("------\n")
	var asValue func(int, int) int
	asValue = sumi

	result1 := calculate(add, 5, 3)          // 8
	result2 := calculate(multiply, 5, 3)     // 15
	fmt.Printf("------%d\n", asValue(10, 2)) // 12
	fmt.Printf("------%d\n", result1)
	fmt.Printf("------%d\n", result2)
}

func sumi(x, y int) (total int) {
	total = x + y
	return
}

// Define a function type
type MathOperation func(int, int) int

func calculate(op MathOperation, a, b int) int {
	return op(a, b)
}

func add(a, b int) int {
	return a + b
}

func multiply(a, b int) int {
	return a * b
}
