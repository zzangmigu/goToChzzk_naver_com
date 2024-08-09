package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
)

func main() {
	var n int
	fmt.Scan(&n)

	// 난이도 의견이 없다면 0을 출력
	if n == 0 {
		fmt.Println(0)
		return
	}

	// 난이도 의견을 담을 slice
	opinions := make([]int, n)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	for i := 0; i < n; i++ {
		scanner.Scan()
		fmt.Sscan(scanner.Text(), &opinions[i])
	}

	// 의견을 정렬
	sort.Ints(opinions)

	// 절사할 갯수 계산
	trimCount := int(math.Round(float64(n) * 0.15))

	// 절사평균 구하기
	sum := 0
	for i := trimCount; i < n-trimCount; i++ {
		sum += opinions[i]
	}

	// 절사평균 계산 후 반올림
	average := float64(sum) / float64(n-2*trimCount)
	finalDifficulty := int(math.Round(average))

	// 결과 출력
	fmt.Println(finalDifficulty)
}
