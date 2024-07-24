package main

import (
	"fmt"
)

func main() {
	var sum int
	arr1 := make([]int, 5)
	for i := 0; i < 5; i++ {
		fmt.Scan(&arr1[i])
		if arr1[i] < 40 {
			arr1[i] = 40
		}
		sum += arr1[i]
	}
	fmt.Println(sum / 5)
}
