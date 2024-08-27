package main

import (
	"fmt"
)

func main() {
	var H, W int
	fmt.Scan(&H)
	fmt.Scan(&W)

	// 짧은 변의 절반을 센티미터 단위로 변환
	radius := min(H, W) * 50

	fmt.Println(radius)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
