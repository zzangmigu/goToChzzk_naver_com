package main

import (
	"fmt"
	"strconv"
)

func main() {
	var N string
	fmt.Scan(&N)

	// 문자열의 길이를 반으로 나누어 두 부분으로 분리
	mid := len(N) / 2
	firstPart := N[:mid]
	secondPart := N[mid:]

	// 두 부분을 정수로 변환
	firstNum, _ := strconv.Atoi(firstPart)
	secondNum, _ := strconv.Atoi(secondPart)

	// 결과 출력
	fmt.Println(firstNum, secondNum)
}
