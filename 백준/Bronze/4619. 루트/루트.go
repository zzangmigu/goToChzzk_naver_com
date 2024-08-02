package main

import (
    "fmt"
    "math"
)

func main() {
    for {
        var B, N int
        fmt.Scan(&B, &N)
        if B == 0 && N == 0 {
            break
        }

        // A의 초기값을 설정
        A := int(math.Pow(float64(B), 1.0/float64(N)))
        // A와 A+1의 N제곱을 비교하여 더 가까운 값을 선택
        if math.Abs(math.Pow(float64(A), float64(N))-float64(B)) > math.Abs(math.Pow(float64(A+1), float64(N))-float64(B)) {
            A++
        }

        fmt.Println(A)
    }
}
