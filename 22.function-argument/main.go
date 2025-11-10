package main

import "fmt"

// pass by value --> go copy the value and modify. The real value still unchange
func modify(num int) {
	num = 50
}

// pass by ref --> go change the address of the real value
func modify2(num *int) {
	*num = 50
}

func main() {
	num := 20
	modify(num)
	fmt.Printf("Using pass by value: %d\n", num)
	modify2(&num)
	fmt.Printf("Using pass by ref: %d\n", num)
}
