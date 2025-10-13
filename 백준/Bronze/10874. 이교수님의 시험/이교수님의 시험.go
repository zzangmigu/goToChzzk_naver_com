package main

import (
	"fmt"
)

func main() {
	var N int
	fmt.Scanln(&N)
	correctAnswers := [10]int{}
	retestStudents := []int{}
	for i := 0; i < 10; i++ {
		correctAnswers[i] = ((i) % 5) + 1
	}

	for j := 1; j <= N; j++ {
		studentAnswers := [10]int{}
		score := 0

		for k := 0; k < 10; k++ {
			fmt.Scan(&studentAnswers[k])
			if studentAnswers[k] == correctAnswers[k] {
				score++
			}
		}
		if score == 10 {
			retestStudents = append(retestStudents, j)
		}
	}
	for _, num := range retestStudents {
		fmt.Println(num)
	}

}
