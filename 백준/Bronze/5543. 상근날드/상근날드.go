package main

import (
	"fmt"
)

func main() {
	// 햄버거 가격 입력
	var sangduk, jungduk, haduk int
	fmt.Scan(&sangduk)
	fmt.Scan(&jungduk)
	fmt.Scan(&haduk)

	// 음료 가격 입력
	var cola, cider int
	fmt.Scan(&cola)
	fmt.Scan(&cider)

	// 모든 가능한 세트 메뉴 가격 계산
	minBurger := threeMin(sangduk, jungduk, haduk)
	minDrink := twoMin(cola, cider)
	minPrice := minBurger + minDrink - 50

	// 결과 출력
	fmt.Println(minPrice)
}

// 두 값의 최소값을 찾는 함수
func twoMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// 세 값의 최소값을 찾는 함수
func threeMin(a, b, c int) int {
	return twoMin(twoMin(a, b), c)
}
