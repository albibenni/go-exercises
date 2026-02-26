package main

import (
	"fmt"
	"unsafe"
)

func main_array_and_slices() {
	s := make([]int, 8)
	ar := [...]int{1, 2, 3, 4} // infer

	fmt.Printf("size of slice %d\n", unsafe.Sizeof(s))
	fmt.Printf("type of slice s: %T, len: %d, cap: %d\n", s, len(s), cap(s))
	s = append(s, 10)

	fmt.Printf("size of slice %d\n", unsafe.Sizeof(s))
	fmt.Printf("type of slice s: %T, len: %d, cap: %d\n", s, len(s), cap(s))
	fmt.Printf("type of array ar: %T\n", ar)
}
