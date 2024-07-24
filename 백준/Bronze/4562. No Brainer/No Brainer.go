package main

import (
	"fmt"
)

func main() {
	var n, eats, requires int
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		fmt.Scan(&eats, &requires)
		if eats < requires {
			fmt.Println("NO BRAINS")
		} else {
			fmt.Println("MMM BRAINS")
		}
	}
}
