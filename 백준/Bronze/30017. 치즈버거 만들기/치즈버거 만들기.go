package main

import (
	"fmt"
)

func main() {
	var A, B int
	fmt.Scan(&A, &B)

	maxCheese := B
	if A-1 < B {
		maxCheese = A - 1
	}

	maxBurgerSize := 2*maxCheese + 1
	fmt.Println(maxBurgerSize)
}
