package main

import (
	"fmt"
)

func main() {
	var N int
	fmt.Scan(&N)
	for i := 0; i < N; i++ {
		var a, b int
		fmt.Scan(&a, &b)
		a = a % 10
		if b == 0 {
			fmt.Println(1)
			continue
		}

		b = b % 4
		if b == 0 {
			b = 4
		}
		result := 1
		for j := 0; j < b; j++ {
			result = (result * a) % 10
		}

		if result == 0 {
			fmt.Println(10)
		} else {
			fmt.Println(result)
		}
	}

}
