package main

import "fmt"

// variadic function is the function that receive the unpredictable parameter
func sum(nums ...int) int {
	total := 0
	for _, n := range nums {
		total += n
	}
	return total
}

func main() {
	fmt.Println("Sum of 1,2,3", sum(1, 2, 3))
	fmt.Println("Sum of 1,2", sum(1, 2))
}
