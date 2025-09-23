package main

import (
	"fmt"
)

func main() {
	var a, b, N int
	var operator string
	for {
		fmt.Scan(&a, &operator, &b)
		N++
		if operator == "E" {
			break
		}

		if operator == ">" {
			fmt.Printf("Case %d: %t\n", N, a > b)
		} else if operator == ">=" {
			fmt.Printf("Case %d: %t\n", N, a >= b)
		} else if operator == "<" {
			fmt.Printf("Case %d: %t\n", N, a < b)
		} else if operator == "<=" {
			fmt.Printf("Case %d: %t\n", N, a <= b)
		} else if operator == "==" {
			fmt.Printf("Case %d: %t\n", N, a == b)
		} else if operator == "!=" {
			fmt.Printf("Case %d: %t\n", N, a != b)
		}
	}

}
