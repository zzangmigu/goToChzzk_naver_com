package main

import (
	"fmt"
	"strconv"
)

func isSelfReplicatingNumber(n int) bool {
	square := n * n
	//제곱한 수를 문자열로
	squareStr := strconv.Itoa(square)
	//n을 문자열로
	nStr := strconv.Itoa(n)
	return squareStr[len(squareStr)-len(nStr):] == nStr
}

func main() {
	var T int
	fmt.Scan(&T)
	for i := 0; i < T; i++ {
		var N int
		fmt.Scan(&N)
		if isSelfReplicatingNumber(N) {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}
