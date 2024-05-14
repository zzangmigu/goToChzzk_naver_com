package main

import (
	"fmt"
)

func main() {
	var testTime, repeatTime int
	var S string

	fmt.Scan(&testTime)
	for i := 0; i < testTime; i++ {
		fmt.Scan(&repeatTime, &S)

		for j := 0; j < len(S); j++ {
			for k := 0; k < repeatTime; k++ {
				fmt.Print(string(S[j]))
			}
		}
		fmt.Println()

	}
}
