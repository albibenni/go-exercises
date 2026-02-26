package main

import "fmt"

func main_map() {

	grades := make(map[string]int)

	grades["Alice"] = 95
	grades["Bob"] = 87
	grades["Charlie"] = 92

	if grade, exists := grades["Alice"]; exists {
		fmt.Printf("Alice's grade: %d\n", grade)
	}

	for student, grade := range grades {
		fmt.Printf("%s: %d\n", student, grade)
	}

	delete(grades, "Bob")
	fmt.Printf("After deletion: %v\n", grades)
}
