package main

import (
	"fmt"
)

func main() {
	var N int
	fmt.Scan(&N)
	for i := 0; i < N; i++ {
		for j := 0; j < i; j++ {
			fmt.Print(" ")
		}
		for j := 0; j < 2*(N-i)-1; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}

	for i := N - 2; i >= 0; i-- {
		for j := 0; j < i; j++ {
			fmt.Print(" ")
		}
		for j := 0; j < 2*(N-i)-1; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}
