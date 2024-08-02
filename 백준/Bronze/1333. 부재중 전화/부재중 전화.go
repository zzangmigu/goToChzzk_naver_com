package main

import (
	"fmt"
)

func main() {
	var N, L, D int
	fmt.Scan(&N, &L, &D)

	// 총 앨범 재생 시간 (마지막 노래 후에는 조용한 구간이 없으므로 5초 추가하지 않음)
	totalAlbumTime := N*L + (N-1)*5

	// 전화벨이 울리는 시간 확인
	for t := 0; t <= totalAlbumTime; t += D {
		// t가 노래 재생 시간과 겹치지 않는지 확인
		if !isPlaying(t, N, L) {
			fmt.Println(t)
			return
		}
	}

	// 앨범 재생이 끝난 후에 울리는 가장 첫 번째 전화벨
	fmt.Println(totalAlbumTime + D - (totalAlbumTime % D))
}

// isPlaying 함수는 t 초가 노래 재생 중인지 확인
func isPlaying(t, N, L int) bool {
	for i := 0; i < N; i++ {
		start := i * (L + 5)
		end := start + L
		if t >= start && t < end {
			return true
		}
	}
	return false
}
