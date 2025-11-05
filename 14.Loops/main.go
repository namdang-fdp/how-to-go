package main

import "fmt"

func main() {
	// basic loops
	// for i := 0; i < 5; i++ {
	// 	fmt.Printf("namdang-fdp\n")
	// }
	// we also can refactor this to range variable
	// intRVariable := []int{1, 2, 3, 4, 5}
	// for i := range intRVariable {
	// 	fmt.Println(i)
	// }

	// infinite loops
	// for {
	// 	fmt.Printf("namdang-fdp loops\n")
	// }

	// while loops implement by for
	i1 := 0
	for i1 < 3 {
		i1 += 2
		fmt.Printf("Here\n")
	}
	fmt.Println(i1)

	// range in for loop
	rvariable := []string{"nam", "namdang", "namdang-fdp"}
	for i, j := range rvariable {
		fmt.Println(i, j)
	}

	// special, if the character is unicode --> i is increase 1
	for i, r := range "Nàm" {
		fmt.Printf("index=%d, rune=%U, char=%c\n", i, r, r)
	}
	// index=0, rune=U+004E, char=N
	// index=1, rune=U+00E0, char=à
	// index=3, rune=U+006D, char=m

	// using map
	rmap := map[int]string{
		22: "nam",
		33: "namdang",
		44: "namdang-fdp",
	}
	for key, value := range rmap {
		fmt.Println(key, value)
	}

	//using with channel
	chnl := make(chan int)

	go func() {
		chnl <- 100
		chnl <- 1000
		chnl <- 10000
		close(chnl)
	}()

	for i := range chnl {
		fmt.Println(i)
	}
}
