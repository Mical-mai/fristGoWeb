package main

import "fmt"

func main() {
	a := 1
	for i := 9; i > 0; i-- {
		x := (a + 1) * 2
		a = x
	}
	fmt.Println(a)
	fmt.Println(a/10)
}
