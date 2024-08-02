package main

import (
	"fmt"
	"sort"
)

type Problem struct {
	sub1 int
	sub2 int
}

func main() {
	var N, L, K int
	fmt.Scan(&N, &L, &K)
	
	problems := make([]Problem, N)
	for i := 0; i < N; i++ {
		fmt.Scan(&problems[i].sub1, &problems[i].sub2)
	}
	
	easyProblems := []Problem{}
	hardProblems := []Problem{}
	
	for _, p := range problems {
		if p.sub2 <= L {
			hardProblems = append(hardProblems, p)
		} else if p.sub1 <= L {
			easyProblems = append(easyProblems, p)
		}
	}
	
	sort.Slice(hardProblems, func(i, j int) bool {
		return hardProblems[i].sub2 < hardProblems[j].sub2
	})
	sort.Slice(easyProblems, func(i, j int) bool {
		return easyProblems[i].sub1 < easyProblems[j].sub1
	})
	
	score := 0
	problemsSolved := 0
	
	for i := 0; i < len(hardProblems) && problemsSolved < K; i++ {
		score += 140
		problemsSolved++
	}
	
	for i := 0; i < len(easyProblems) && problemsSolved < K; i++ {
		score += 100
		problemsSolved++
	}
	
	fmt.Println(score)
}
