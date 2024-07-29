package main

import (
	"fmt"
)

type Person struct {
	weight int
	height int
}

func main() {
	var n int
	fmt.Scan(&n)

	people := make([]Person, n)
	rankings := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&people[i].weight, &people[i].height)
	}

	for i := 0; i < n; i++ {
		rank := 1
		for j := 0; j < n; j++ {
			if i != j && people[i].weight < people[j].weight && people[i].height < people[j].height {
				rank++
			}
		}
		rankings[i] = rank
	}

	for _, rank := range rankings {
		fmt.Print(rank, " ")
	}
}
