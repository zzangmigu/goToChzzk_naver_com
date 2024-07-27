package main

import (
	"fmt"
)

func main() {
	var N, A, B int
	fmt.Scan(&N)
	fmt.Scan(&A, &B)

	maxChickenFromCola := A / 2
	maxChickenFromBeer := B

	maxChicken := min(N, maxChickenFromCola+maxChickenFromBeer)

	fmt.Println(maxChicken)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
