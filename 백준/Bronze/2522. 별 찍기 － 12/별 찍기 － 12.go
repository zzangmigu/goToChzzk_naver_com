package main

import (
	"fmt"
	"strings"
)

func main() {
	var N int
	fmt.Scan(&N)

	// 상반부: 1 to N
	for i := 1; i <= N; i++ {
		fmt.Printf("%s%s\n", strings.Repeat(" ", N-i), strings.Repeat("*", i))
	}

	// 하반부: N-1 to 1
	for i := N-1; i >= 1; i-- {
		fmt.Printf("%s%s\n", strings.Repeat(" ", N-i), strings.Repeat("*", i))
	}
}
