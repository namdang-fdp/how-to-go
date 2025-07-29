package main

import "fmt"

func main() {
	// int type and arimethic
	var X uint8 = 255
	var Y int16 = 32767
	fmt.Println(X, X-3)
	fmt.Println(Y+2, Y-2)

	var a int16 = 170
	var b int16 = 83
	fmt.Printf("Addition: %d + %d = %d\n", a, b, a+b)
	fmt.Printf("Subtraction: %d - %d = %d\n", a, b, a-b)
	fmt.Printf("Mutiplication: %d * %d = %d\n", a, b, a*b)
	fmt.Printf("Division: %d + %d = %d\n", a, b, a/b)
	fmt.Printf("Remainder: %d + %d = %d\n", a, b, a%b)

	// floating point
	c := 20.45
	d := 34.89
	e := d - c

	fmt.Printf("Result of d-c: %f\n", e)
	fmt.Printf("Data type of d-c: %T", e)

	var g float32 = 5.00
	var h float32 = 2.25
	fmt.Printf("Addition: %g + %g = %g\n", g, h, g+h)
	fmt.Printf("Substraction: %g - %g = %g\n", g, h, g-h)
	fmt.Printf("Multiplication: %g * %g = %g\n", g, h, g*h)
	fmt.Printf("Division: %g / %g = %g\n", g, h, g/h)

	// complex number
	var j complex64 = complex(6, 2)
	var k complex128 = complex(9, 2)
	fmt.Println(j)
	fmt.Println(k)
	realNum := real(j)
	fmt.Println(realNum)
	imaginary := imag(k)
	fmt.Println(imaginary)

	// boolean
	str1 := "NamDang"
	str2 := "namdang"
	str3 := "namdangs"
	rs1 := str1 == str2
	rs2 := str1 == str3
	fmt.Println(rs1)
	fmt.Println(rs2)

	// string type
	str := "namdangcoder"
	fmt.Printf("Length of the string is: %d\n", len(str))
	strName := "Name"
	fmt.Printf("%s", strName+str)
}
