package main

import (
	"fmt"
)

func main() {
	var N int
	fmt.Scan(&N)

	count := int64(N) * int64(N) * int64(N)
	degree := 3
	fmt.Println(count)
	fmt.Println(degree)

}
