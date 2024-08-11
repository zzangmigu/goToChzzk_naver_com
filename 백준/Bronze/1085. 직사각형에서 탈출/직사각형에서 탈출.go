package main

import (
    "fmt"
    "math"
)

func main() {
    var x, y, w, h int
    fmt.Scan(&x, &y, &w, &h)

    // x, y로부터 가장 가까운 직사각형 경계까지의 거리 계산
    minDistanceX := math.Min(float64(x), float64(w-x))
    minDistanceY := math.Min(float64(y), float64(h-y))

    // 두 거리 중 더 작은 값을 결과로 출력
    result := math.Min(minDistanceX, minDistanceY)
    fmt.Printf("%d\n", int(result))
}
