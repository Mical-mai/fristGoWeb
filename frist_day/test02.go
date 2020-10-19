package main

import "fmt"

func main() {
	a, b, sum := 2.0, 1.0, 0.0
	for i := 1; i < 21; i++ {
		sum += a / b
		t := a
		a = a + b
		b = t
	}
	fmt.Println(sum)
}
