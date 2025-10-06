package main

import (
	"fmt"
)

func main() {
	var T, result int
	fmt.Scanln(&T)
	for i := 0; i < T; i++ {
		var multiStrip int
		fmt.Scanln(&multiStrip)
		result += multiStrip
	}
	fmt.Println(result - T + 1)
}
