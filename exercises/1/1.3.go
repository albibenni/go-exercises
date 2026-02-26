package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func v1(str []string) string {
	var s, sep string
	for i := range str {
		s += sep + " " + os.Args[i]
		sep = " "
	}
	return s
}
func v2OPT(str []string) string {
	return strings.Join(str[0:], " ")
}
func main_1_3() {
	args := os.Args
	start := time.Now()
	res := v1(args)
	elapsed := time.Since(start).Microseconds()
	fmt.Printf("%s\n", res)
	fmt.Printf("%d\n", elapsed)
	start2 := time.Now()
	res2 := v2OPT(args)
	elapsed2 := time.Since(start2).Microseconds()

	fmt.Printf("%s\n", res2)
	fmt.Printf("%d\n", elapsed2)
}
