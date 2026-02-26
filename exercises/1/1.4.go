package main

import (
	"bufio"
	"fmt"
	"os"
)

func main_1_4() {
	counts := make(map[string]int)
	count := 0
	files := os.Args[1:]
	fmt.Println(files)
	if len(files) == 0 {
		count = 1
		countLines(os.Stdin, counts, &count)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			count = 1
			countLines(f, counts, &count)
			f.Close()
		}
	}
	for line, n := range counts {
		if n >= 1 {
			fmt.Printf("%d\t%s\n", n, line)
		} else {

			fmt.Println("n<=1")
		}
	}
}
func countLines(f *os.File, counts map[string]int, count *int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()] = *count
		*count++
	}
	// NOTE: ignoring potential errors from input.Err()
}
