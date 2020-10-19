package main

import "fmt"

func main() {

	for i := 'x'; i <= 'z'; i++ {
		for j := 'x'; j <= 'z'; j++ {
			if i != j {
				for k := 'x'; k <= 'z'; k++ {
					if j != k && i != k {
						if i != 'x' && k != 'x' && k != 'z' {
							fmt.Printf("a和%s\n", string(i))
							fmt.Printf("b和%s\n", string(j))
							fmt.Printf("c和%s\n", string(k))
						}
					}

				}
			}
		}
	}
}
