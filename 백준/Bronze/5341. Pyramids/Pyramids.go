package main

import (
	"fmt"
)

func main() {
	var n int
	for true {
		fmt.Scan(&n)
		if n == 0 {
			break
		}
		fmt.Println(n * (n + 1) / 2)
	}
}
