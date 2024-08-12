package main

import (
	"fmt"
	"strconv"
)

func toBase9(n int) string {
	if n == 0 {
		return "0"
	}
	base9 := ""
	for n > 0 {
		remainder := n % 9
		base9 = strconv.Itoa(remainder) + base9
		n = n / 9
	}
	return base9
}

func main() {
	var T int
	fmt.Scan(&T)
	fmt.Println(toBase9(T))
}
