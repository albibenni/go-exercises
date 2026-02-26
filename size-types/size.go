package main

import (
	"fmt"
	"unsafe"
)

func main_size() {
	// Basic types
	fmt.Printf("bool:       %d bytes\n", unsafe.Sizeof(bool(true)))
	fmt.Printf("int8:       %d bytes\n", unsafe.Sizeof(int8(0)))
	fmt.Printf("uint8:      %d bytes\n", unsafe.Sizeof(uint8(0)))
	fmt.Printf("int16:      %d bytes\n", unsafe.Sizeof(int16(0)))
	fmt.Printf("uint16:     %d bytes\n", unsafe.Sizeof(uint16(0)))
	fmt.Printf("int32:      %d bytes\n", unsafe.Sizeof(int32(0)))
	fmt.Printf("uint32:     %d bytes\n", unsafe.Sizeof(uint32(0)))
	fmt.Printf("int64:      %d bytes\n", unsafe.Sizeof(int64(0)))
	fmt.Printf("uint64:     %d bytes\n", unsafe.Sizeof(uint64(0)))
	fmt.Printf("int:        %d bytes\n", unsafe.Sizeof(int(0)))
	fmt.Printf("uint:       %d bytes\n", unsafe.Sizeof(uint(0)))
	fmt.Printf("uintptr:    %d bytes\n", unsafe.Sizeof(uintptr(0)))
	fmt.Printf("float32:    %d bytes\n", unsafe.Sizeof(float32(0)))
	fmt.Printf("float64:    %d bytes\n", unsafe.Sizeof(float64(0)))
	fmt.Printf("complex64:  %d bytes\n", unsafe.Sizeof(complex64(0)))
	fmt.Printf("complex128: %d bytes\n", unsafe.Sizeof(complex128(0)))
	fmt.Printf("byte:       %d bytes\n", unsafe.Sizeof(byte(0)))
	fmt.Printf("rune:       %d bytes\n", unsafe.Sizeof(rune(0)))
	fmt.Printf("string:     %d bytes\n", unsafe.Sizeof(""))

	// Composite types
	fmt.Printf("[]int:      %d bytes\n", unsafe.Sizeof([]int{}))
	fmt.Printf("map[string]int: %d bytes\n", unsafe.Sizeof(map[string]int{}))
	fmt.Printf("chan int:   %d bytes\n", unsafe.Sizeof(make(chan int)))
	fmt.Printf("*int:       %d bytes\n", unsafe.Sizeof((*int)(nil)))
	fmt.Printf("func():     %d bytes\n", unsafe.Sizeof(func() {}))
	fmt.Printf("interface{}: %d bytes\n", unsafe.Sizeof(interface{}(nil)))

	// Array examples
	fmt.Printf("[5]int:     %d bytes\n", unsafe.Sizeof([5]int{}))
	fmt.Printf("[10]byte:   %d bytes\n", unsafe.Sizeof([10]byte{}))

	// Struct example
	type Person struct {
		Name string
		Age  int
	}
	fmt.Printf("Person struct: %d bytes\n", unsafe.Sizeof(Person{}))
}
