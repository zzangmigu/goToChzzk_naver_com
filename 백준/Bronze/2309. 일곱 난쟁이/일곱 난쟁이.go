package main

import (
	"fmt"
	"sort"
)

func main() {
	heights := make([]int, 9)
	var total int

	for i := 0; i < 9; i++ {
		fmt.Scan(&heights[i])
		total += heights[i]
	}

	found := false
	for i := 0; i < 8 && !found; i++ {
		for j := i + 1; j < 9; j++ {
			if total-heights[i]-heights[j] == 100 {
				heights[i], heights[j] = 0, 0
				found = true
				break
			}
		}
	}
	
	result := make([]int, 0, 7)
	for _, height := range heights {
		if height != 0 {
			result = append(result, height)
		}
	}

	sort.Ints(result)
	for _, height := range result {
		fmt.Println(height)
	}
}
