package main

import "fmt"

// go do not allow you to define the value but not use
// if you declare a value but not need to call it
// use _ to make the compiler understand that is the init value

func divmod(a, b int) (int, int) {
	return a / b, a % b
}

func main() {
	quotient, _ := divmod(10, 3)
	fmt.Printf("The quotient is: %d", quotient)
}
