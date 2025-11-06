package main

import "fmt"

func main() {
	c := make(chan int)
	select {
	case <-c:
	default:
		fmt.Println("Default case")
	}
}
