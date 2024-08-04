package main

import (
	"fmt"
	"math"
)

func main() {
	var N, S, M, L, XL, XXL, XXXL, T, P int
	fmt.Scan(&N)
	fmt.Scan(&S, &M, &L, &XL, &XXL, &XXXL)
	fmt.Scan(&T, &P)

	// 티셔츠 묶음 계산
	totalBundles := int(math.Ceil(float64(S)/float64(T))) +
		int(math.Ceil(float64(M)/float64(T))) +
		int(math.Ceil(float64(L)/float64(T))) +
		int(math.Ceil(float64(XL)/float64(T))) +
		int(math.Ceil(float64(XXL)/float64(T))) +
		int(math.Ceil(float64(XXXL)/float64(T)))

	// 펜 묶음 계산
	penBundles := N / P
	singlePens := N % P

	// 결과 출력
	fmt.Println(totalBundles)
	fmt.Println(penBundles, singlePens)
}
