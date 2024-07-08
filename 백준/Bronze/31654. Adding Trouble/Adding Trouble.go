package main

import (
	"fmt"
)

func main() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)
	if c == a+b {
		fmt.Println("correct!")
	} else if c != a+b {
		fmt.Println("wrong!")
	}
}
