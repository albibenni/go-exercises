package main

import "fmt"

func main_defer() {
	fmt.Println("Starting main function")
	fmt.Println("-------")
	multipleDefers()
	fmt.Println("-------")

	multipleDefersFor()
	fmt.Println("-------")
	fmt.Printf("defer modifying return value from 5 to: %d\n", example())
	fmt.Println("-------")
	safeFunction()
	fmt.Println("Ending main function")
}
func multipleDefers() {
	defer fmt.Println("First defer")
	defer fmt.Println("Second defer")
	defer fmt.Println("Third defer")
}

// Output:
// Third defer
// Second defer
// First defer
func multipleDefersFor() {
	for i := 0; i < 10; i++ {
		defer fmt.Printf("defer %d\n", i)
	}
}

func example() (result int) {
	defer func() {
		result++ // Modifies the return value
	}()
	return 5 // Actually returns 6
}

func safeFunction() {

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()
	panic(fmt.Sprintf("panic but not too much because we recover()"))

	// Code that might panic
}
