package main

import (
	"fmt"
)

func main() {
	var N int
	fmt.Scan(&N)

	result := 0

	for i := 1; i < N; i++ {
		sum := i
		temp := i

		for temp != 0 {
			sum += temp % 10
			temp /= 10
		}
		
		if sum == N {
			result = i
			break
		}
	}

	fmt.Println(result)
}
