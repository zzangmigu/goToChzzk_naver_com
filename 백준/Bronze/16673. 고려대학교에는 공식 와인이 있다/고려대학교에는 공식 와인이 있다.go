package main

import "fmt"

func main() {
	var C, K, P int
	fmt.Scan(&C, &K, &P)
	fmt.Println(calcWines(C, K, P))

}

func calcWines(C, K, P int) int {
	sumWines := 0
	n := 1
	for i := 0; i < C; i++ {
		sumWines += K*n + P*n*n
		n++

	}
	return sumWines

}
