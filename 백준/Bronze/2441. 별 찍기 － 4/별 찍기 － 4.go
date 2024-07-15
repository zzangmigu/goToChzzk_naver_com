package main

import (
	"fmt"
	"strings"
)

func main() {
	var N int
	fmt.Scan(&N)

	for i := 0; i < N; i++ {
		fmt.Print(strings.Repeat(" ", i))
		fmt.Println(strings.Repeat("*", N-i))
	}
}
