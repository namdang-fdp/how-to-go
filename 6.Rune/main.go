package main

import (
	"fmt"
	"reflect"
)

func main() {
	rune1 := 'B'
	fmt.Printf("Rune 1: %c; Unicode: %U; Type: %s", rune1, rune1, reflect.TypeOf(rune1))
}
