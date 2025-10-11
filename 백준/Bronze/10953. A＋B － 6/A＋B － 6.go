package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var N int
	fmt.Scanln(&N)
	for i := 0; i < N; i++ {
		var testNumStr string
		fmt.Scanln(&testNumStr)
		testNumSplit := strings.Split(testNumStr, "")
		numA, _ := strconv.Atoi(testNumSplit[0])
		numB, _ := strconv.Atoi(testNumSplit[2])
		fmt.Println(numA + numB)

	}

}
