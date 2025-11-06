package main

import "fmt"

func name(i func(p, q string) string) {
	fmt.Println(i("nam", "dang"))
}

func name2() func(i, j string) string {
	myf := func(i, j string) string {
		return i + j + "namdang-fdp"
	}
	return myf
}

func main() {
	value := func(p, q string) string {
		return p + q + "fdp"
	}
	name(value)

	value2 := name2()
	fmt.Println(value2("Welcome", "to"))
}
