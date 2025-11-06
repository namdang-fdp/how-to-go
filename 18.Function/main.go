package main

import "fmt"

// syntax: func + name (parameter + param_data_types) return data_type
func mutiply(a, b int) int {
	return a * b
}

func multiply(a, b *int) int {
	*a = *a * 2
	return *a * *b
}

func main() {
	x := 5
	y := 10
	rs := mutiply(5, 10)
	result := multiply(&x, &y)
	fmt.Println(result)
	fmt.Println(rs)
}
