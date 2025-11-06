package main

import "fmt"

func main() {
	// break you stop the loops
	for i := range []int{1, 2, 3, 4, 5} {
		fmt.Println(i)
		if i == 3 {
			break
		}
	}

	// go to will tranfer to the control
	var x int = 0
Label1:
	for x < 8 {
		if x == 5 {
			x = x + 1
			goto Label1
		}
		fmt.Printf("Value is: %d\n", x)
		x++
	}

	// continue will stop that loops for 1 time (not the entire loop)
	var y int = 0
	for y < 8 {
		if y == 5 {
			y += 2
			continue
		}
		fmt.Printf("Value is: %d\n", y)
		y++
	}
}
