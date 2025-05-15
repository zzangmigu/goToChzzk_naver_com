package main

import (
	"fmt"
)

func main() {
	var R, C int
	fmt.Scan(&R, &C)

	parkingLot := make([][]rune, R)
	for i := range parkingLot {
		var row string
		fmt.Scan(&row)
		parkingLot[i] = []rune(row)
	}

	counts := make([]int, 5)

	for i := 0; i < R-1; i++ {
		for j := 0; j < C-1; j++ {
			cars := 0
			building := false
			positions := [4][2]int{{i, j}, {i, j + 1}, {i + 1, j}, {i + 1, j + 1}}

			for _, pos := range positions {
				r, c := pos[0], pos[1]
				if parkingLot[r][c] == '#' {
					building = true
					break
				} else if parkingLot[r][c] == 'X' {
					cars++
				}
			}

			if !building {
				counts[cars]++
			}
		}
	}
	
	for _, count := range counts {
		fmt.Println(count)
	}
}
