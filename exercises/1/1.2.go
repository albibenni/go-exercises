package main

import (
	"fmt"
	"os"
	"strconv"
)

func main_1_2() {
	var s, sep string
	for i := range os.Args {
		s += sep + strconv.Itoa(i) + " " + os.Args[i]
		sep = "\n"
	}
	fmt.Printf("%s\n", s)
}
