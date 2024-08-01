package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var X string
	fmt.Scan(&X)

	if strings.HasPrefix(X, "0x") {
		// 16진수 처리
		number, _ := strconv.ParseInt(X[2:], 16, 64)
		fmt.Println(number)
	} else if strings.HasPrefix(X, "0") && len(X) > 1 {
		// 8진수 처리
		number, _ := strconv.ParseInt(X, 8, 64)
		fmt.Println(number)
	} else {
		// 10진수 처리
		number, _ := strconv.ParseInt(X, 10, 64)
		fmt.Println(number)
	}
}
