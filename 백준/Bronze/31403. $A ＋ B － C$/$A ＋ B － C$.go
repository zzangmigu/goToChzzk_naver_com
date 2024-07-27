package main

import (
	"fmt"
	"strconv"
)

func main() {
	var A, B, C int
	fmt.Scan(&A)
	fmt.Scan(&B)
	fmt.Scan(&C)

	numResult := A + B - C
	fmt.Println(numResult)

	strA := strconv.Itoa(A)
	strB := strconv.Itoa(B)
	strC := strconv.Itoa(C)
	concatenatedAB := strA + strB
	numConcatenatedAB, _ := strconv.Atoi(concatenatedAB)
	numC, _ := strconv.Atoi(strC)

	strResult := numConcatenatedAB - numC
	fmt.Println(strResult)
}
