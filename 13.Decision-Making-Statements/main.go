package main

import "fmt"

func main() {
	var v int = 700

	// basic if condition
	if v < 1000 {
		fmt.Println("v is less than 1000")
	} else {
		fmt.Printf("v is greater than 1000")
	}

	// nested
	var a = 400
	var b = 500
	if a == 400 {
		if b == 500 {
			fmt.Printf("Here\n")
		}
	}

	// combine with else if
	var c = 1000
	if c == 10 {
		fmt.Println("c is 10")
	} else if c == 100 {
		fmt.Println("c is 100")
	} else if c == 1000 {
		fmt.Println("c is 1000")
	} else {
		fmt.Println("C is nether 10 nor 100 nor 1000")
	}
}
