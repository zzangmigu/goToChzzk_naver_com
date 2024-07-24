package main

import (
	"fmt"
)

func main() {
	var A, P, C int
	fmt.Scan(&A, &P, &C)

	maxPrize := max(A+C, P)
	
	fmt.Println(maxPrize)
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
