package main

import "fmt"

func main_string() {
	s := "my string"
	bytes := []byte(s)
	run := []rune(s)
	sall := string(run)

	fmt.Printf("string %s \n", s)
	fmt.Printf("bytes %v \n", bytes)
	fmt.Printf("bytes %v \n", run)
	fmt.Printf("bytes %v \n", sall)
}
