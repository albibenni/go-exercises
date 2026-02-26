package main

import "fmt"

func main_variadic() {
	fmt.Printf("------\n")

	fmt.Printf("sum - %d\n", sum(1, 2, 5, 10))

	myNums := []int{1, 2, 51, 2, 2}
	fmt.Printf("sum variadic op. - %d\n", sum(myNums...))
}
func sum(numbers ...int) (total int) {
	total = 0
	for _, num := range numbers {
		total += num
	}
	return total
}
