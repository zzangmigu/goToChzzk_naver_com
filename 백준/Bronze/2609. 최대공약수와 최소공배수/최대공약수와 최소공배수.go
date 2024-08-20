package main

import (
	"fmt"
)

// 최대 공약수를 구하는 함수 (유클리드 호제법)
func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

// 최소 공배수를 구하는 함수
func lcm(a, b, gcd int) int {
	return a * b / gcd
}

func main() {
	var a, b int
	fmt.Scan(&a, &b)

	g := gcd(a, b)
	l := lcm(a, b, g)

	fmt.Println(g)
	fmt.Println(l)
}
