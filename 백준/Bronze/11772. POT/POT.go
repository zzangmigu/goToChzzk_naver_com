package main

import (
	"fmt"
)

func intPow(base int, exp int) int64 {
	result := int64(1)
	base64 := int64(base)

	for i := 0; i < exp; i++ {
		result *= base64
	}
	return result
}

func main() {
	var N int
	var total int64 = 0
	fmt.Scanln(&N)

	for i := 0; i < N; i++ {
		var x int
		fmt.Scanln(&x)
		exponent := x % 10
		base := x / 10
		tempValue := intPow(base, exponent)
		total += tempValue
	}
	fmt.Println(total)

}
