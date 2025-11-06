package main

import "fmt"

func main() {

	// expression switch
	switch day := 4; day {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	default:
		fmt.Println("Invalid")
	}

	var month interface{} = 1
	switch v := month.(type) {
	case int:
		switch v {
		case 1:
			fmt.Println("January")
		case 2:
			fmt.Println("Febuary")
		default:
			fmt.Println("Invalid")
		}
	default:
		fmt.Println("Invalid type")
	}

}
