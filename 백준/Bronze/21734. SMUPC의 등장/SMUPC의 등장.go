package main

import (
	"fmt"
)

func main() {
	var S string
	fmt.Scanln(&S)
	for _, char := range S {
		sum := 0
		asciiValue := int(char)
		for asciiValue > 0 {
			remain := asciiValue % 10
			sum += remain
			asciiValue /= 10
		}
		for i := 0; i < sum; i++ {
			fmt.Printf("%c", char)
		}
		fmt.Println()

	}

}
