package main

import (
    "fmt"
)

func main() {
    var T int
    fmt.Scan(&T)

    for i := 0; i < T; i++ {
        var N, M int
        fmt.Scan(&N, &M)

        // L열의 네 번째 자리는 12번째 자리
        row := 12
        col := 4

        if row > N || col > M {
            fmt.Println(-1)
        } else {
            seatNumber := (row-1)*M + col
            fmt.Println(seatNumber)
        }
    }
}
