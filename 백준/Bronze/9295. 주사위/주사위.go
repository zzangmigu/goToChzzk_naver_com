package main

import (
	"fmt"
)

func main() {
	var n, a, b int
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a, &b)
		fmt.Print("Case ", i+1, ": ", a+b)
		fmt.Println()
	}
}
