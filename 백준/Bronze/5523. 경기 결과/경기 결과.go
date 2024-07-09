package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)
	var aWin, bWin int
	for i := 0; i < n; i++ {
		var a, b int
		fmt.Scan(&a, &b)
		if a > b {
			aWin++
		} else if a < b {
			bWin++
		}
	}
	fmt.Println(aWin, bWin)
}
