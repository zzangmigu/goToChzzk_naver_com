package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

// Point 구조체 정의
type Point struct {
	x int
	y int
}

// Points를 정렬하기 위한 인터페이스 정의
type Points []Point

func (p Points) Len() int {
	return len(p)
}

func (p Points) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p Points) Less(i, j int) bool {
	if p[i].y == p[j].y {
		return p[i].x < p[j].x
	}
	return p[i].y < p[j].y
}

func main() {
	// 입력을 받기 위한 bufio 스캐너 사용
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	// 점의 개수 입력 받기
	var n int
	fmt.Fscanf(reader, "%d\n", &n)

	// 점들을 저장할 슬라이스 생성
	points := make(Points, n)

	// 각 점의 좌표를 입력 받아 슬라이스에 저장
	for i := 0; i < n; i++ {
		fmt.Fscanf(reader, "%d %d\n", &points[i].x, &points[i].y)
	}

	// 사용자 정의 정렬 함수 사용하여 정렬
	sort.Sort(points)

	// 정렬된 결과 출력
	for _, point := range points {
		fmt.Fprintf(writer, "%d %d\n", point.x, point.y)
	}
}
