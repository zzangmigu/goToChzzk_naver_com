package main

import (
	"container/list"
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)

	// 큐 초기화
	queue := list.New()
	for i := 1; i <= n; i++ {
		queue.PushBack(i)
	}

	// 카드가 한 장 남을 때까지 반복
	for queue.Len() > 1 {
		// 제일 위에 있는 카드를 버림
		queue.Remove(queue.Front())
		// 제일 위에 있는 카드를 제일 아래로 옮김
		queue.PushBack(queue.Remove(queue.Front()))
	}

	// 마지막 남은 카드 출력
	fmt.Println(queue.Front().Value)
}
