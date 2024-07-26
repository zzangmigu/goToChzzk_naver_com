package main

import (
    "fmt"
)

func main() {
    var N, M int
    fmt.Scan(&N, &M)

    cards := make([]int, N)
    for i := 0; i < N; i++ {
        fmt.Scan(&cards[i])
    }

    closestSum := 0

    // 모든 3장의 카드 조합을 탐색
    for i := 0; i < N-2; i++ {
        for j := i + 1; j < N-1; j++ {
            for k := j + 1; k < N; k++ {
                sum := cards[i] + cards[j] + cards[k]
                if sum <= M && sum > closestSum {
                    closestSum = sum
                }
            }
        }
    }

    fmt.Println(closestSum)
}
