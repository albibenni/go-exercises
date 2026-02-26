package main

import "fmt"

func main_make() {
	// make([]Type, length, capacity)
	s1 := make([]int, 5)       // Creates slice with length 5, capacity 5, zero values
	s2 := make([]int, 3, 10)   // Creates slice with length 3, capacity 10
	s3 := make([]string, 0, 5) // Empty slice with capacity 5
	fmt.Printf("Creates slice with length 5, capacity 5, zero values: %v\n", s1)
	fmt.Printf("Creates slice with length 3, capacity 10: %v\n", s2)
	s2 = append(s2, 10)
	fmt.Printf("Creates slice with length now of 4 with append, capacity 10: %v\n", s2)
	fmt.Printf("Creates slice with length now of %d with append, capacity -func cap() %d: %v\n", len(s2), cap(s2), s2)
	fmt.Printf("Empty slice with capacity 5: %v\n", s3)
	fmt.Printf("Empty slice with capacity 5: %v\n", s3)

	// make([]type, length, capacity)
	newSlice := make([]int, 3, 5) // length=3, capacity=5
	fmt.Println(len(newSlice))    // prints: 3
	fmt.Println(cap(newSlice))    // prints: 5
	fmt.Println(newSlice)         // prints: [0 0 0]

	// You can only access indices 0, 1, 2 initially
	// newSlice[3] would cause a panic, even though capacity allows it

	// But you can append without reallocating until you exceed capacity
	newSlice = append(newSlice, 4) // now length=4, capacity still=5
	newSlice = append(newSlice, 5) // now length=5, capacity=5
	newSlice = append(newSlice, 6) // now length=6, capacity grows (typically doubles)

	fmt.Printf("Slice with length now of %d with append, capacity -func cap() %d: %v\n", len(newSlice), cap(newSlice), newSlice)

	// make(map[KeyType]ValueType)
	m1 := make(map[string]int)      // Creates empty map
	m2 := make(map[string]int, 100) // Creates map with initial space for ~100 elements (hint)

	fmt.Printf("maps %v \n", m1)
	fmt.Printf("maps %v \n", m2)

	// make(chan Type) - unbuffered channel
	ch1 := make(chan int)

	// make(chan Type, buffer_size) - buffered channel
	ch2 := make(chan string, 10) // Buffer size of 10

	fmt.Printf("channels %v \n", ch1)
	fmt.Printf("channels %v \n", ch2)
}
