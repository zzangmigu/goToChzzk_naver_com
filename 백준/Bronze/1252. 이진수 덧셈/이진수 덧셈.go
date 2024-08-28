package main

import (
	"fmt"
	"math/big"
)

func main() {
	var bin1, bin2 string
	fmt.Scan(&bin1, &bin2)

	// 이진수를 큰 정수로 변환
	num1 := new(big.Int)
	num1.SetString(bin1, 2)

	num2 := new(big.Int)
	num2.SetString(bin2, 2)

	// 두 이진수를 더함
	sum := new(big.Int)
	sum.Add(num1, num2)

	// 결과를 이진수 문자열로 변환하여 출력
	fmt.Println(sum.Text(2))
}
