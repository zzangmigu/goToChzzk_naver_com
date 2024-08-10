package main

import (
    "fmt"
    "math"
)

func main() {
    var n int
    fmt.Scan(&n)

    for i := 0; i < n; i++ {
        var A1, P1, R1, P2 int
        fmt.Scan(&A1, &P1)
        fmt.Scan(&R1, &P2)

        // 피자 조각의 1달러당 면적
        sliceValue := float64(A1) / float64(P1)
        // 원형 피자의 1달러당 면적
        wholeValue := (math.Pi * float64(R1) * float64(R1)) / float64(P2)

        if sliceValue > wholeValue {
            fmt.Println("Slice of pizza")
        } else {
            fmt.Println("Whole pizza")
        }
    }
}
