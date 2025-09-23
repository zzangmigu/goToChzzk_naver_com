package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var A, B int
	fmt.Scan(&A, &B)
	strA := strconv.Itoa(A)
	strB := strconv.Itoa(B)

	minStrA := strings.Replace(strA, "6", "5", -1)
	minStrB := strings.Replace(strB, "6", "5", -1)
	minA, _ := strconv.Atoi(minStrA)
	minB, _ := strconv.Atoi(minStrB)
	minSum := minA + minB

	maxStrA := strings.Replace(minStrA, "5", "6", -1)
	maxStrB := strings.Replace(minStrB, "5", "6", -1)
	maxA, _ := strconv.Atoi(maxStrA)
	maxB, _ := strconv.Atoi(maxStrB)
	maxSum := maxA + maxB
	fmt.Println(minSum, maxSum)

}
