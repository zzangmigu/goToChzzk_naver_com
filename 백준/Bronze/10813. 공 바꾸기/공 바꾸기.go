package main

import (
	"fmt"
)

func main() {
	var N, M int
	fmt.Scan(&N, &M)

	// 초기 바구니 상태 설정: basket[i]는 i번 바구니에 있는 공의 번호를 의미
	basket := make([]int, N)
	for i := 0; i < N; i++ {
		basket[i] = i + 1
	}

	// M번의 교환 명령 수행
	for i := 0; i < M; i++ {
		var a, b int
		fmt.Scan(&a, &b)
		// 배열 인덱스는 0부터 시작하므로 a, b에서 1을 뺌
		basket[a-1], basket[b-1] = basket[b-1], basket[a-1]
	}

	// 결과 출력
	for _, ball := range basket {
		fmt.Print(ball, " ")
	}
}
