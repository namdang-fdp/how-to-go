package main

import "fmt"

func calculator(a, b int) (mul int, div int) {
	mul = a * b
	div = a / b

	return // -> go compiler will return the param init in the func
}

func main() {
	m, d := calculator(10, 5)
	fmt.Printf("Rs 1: %d\n", m)
	fmt.Printf("Rs 2: %d\n", d)
}
