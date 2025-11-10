package main

import "fmt"

// defer keyword for the func use to read the value of
// params in the compile time, but will not execute the func
// it will push that func to stack, and pop it when all the statement cover it
// successfully execute. Try to use it whenever you need to clean up, unlock, cancel something
// remember: defer func use the LIFO principles

func main() {
	fmt.Println("Start")
	defer fmt.Println("1st defer")
	defer fmt.Println("2nd defer")
	fmt.Println("Doing something")
	defer fmt.Println("3rd defer")
}
