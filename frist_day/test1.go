package main

import "fmt"

func main() {
	a, sum := 1, 0
	for i := 1; i < 3; i++ {
		a = a * i
		sum += a
		fmt.Printf("第%d个值是：%d\n", i, sum)
	}

}
