package main

import (
	"fmt"
)

func main() {
	var T int
	fmt.Scanln(&T)
	for i := 0; i < T; i++ {
		var N int
		fmt.Scanln(&N)
		currentChange := 0
		minChange := 0
		for j := 0; j < N; j++ {
			var p1, p2 int
			fmt.Scanln(&p1, &p2)
			currentChange += p1 - p2
			if currentChange < minChange {
				minChange = currentChange
			}
		}
		fmt.Println(-minChange)
	}

}
