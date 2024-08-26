package main

import (
	"fmt"
)

func main() {
	var N int
	fmt.Scan(&N)

	if N == 1 {
		fmt.Println(1)
		return
	}

	layer := 1
	count := 1

	for count < N {
		count += 6 * layer
		layer++
	}

	fmt.Println(layer)
}
