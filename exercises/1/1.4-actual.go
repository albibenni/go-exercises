package main

import (
	"bufio"
	"fmt"
	"os"
)

func main_1_4_actual() {
	counts := make(map[string]int)
	files := os.Args[1:]
	fmt.Println(files)
	if len(files) == 0 {
		countLinesDuplicated(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLinesDuplicated(f, counts)
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
func countLinesDuplicated(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
	// NOTE: ignoring potential errors from input.Err()
}
