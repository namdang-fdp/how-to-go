package main

import "fmt"

func myfunc(p, q int) (int, int, int) {
	return p - q, p + q, p * q
}

func main() {
	var val1, val2, val3 = myfunc(10, 5)

	fmt.Printf("Result 1: %d\n", val1)
	fmt.Printf("Result 2: %d\n", val2)
	fmt.Printf("Result 3: %d\n", val3)

}
