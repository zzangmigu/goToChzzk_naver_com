package main

import (
	"fmt"
)

func main() {
	var solved, union, maxLevel int
	fmt.Scan(&solved, &union, &maxLevel)
	if solved >= 1000 {
		if union >= 8000 || maxLevel >= 260 {
			fmt.Println("Very Good")
		} else {
			fmt.Println("Good")

		}

	} else {
		fmt.Println("Bad")
	}
}
