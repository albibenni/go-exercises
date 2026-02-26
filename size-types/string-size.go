package main

import "fmt"

func main_string_size() {
	s := "Hello, 世界"
	fmt.Println(len(s))         // 13 (bytes)
	fmt.Println(len([]rune(s))) // 9 (actual characters)
}
