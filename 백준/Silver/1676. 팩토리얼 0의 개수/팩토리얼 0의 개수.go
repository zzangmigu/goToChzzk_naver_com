package main

import "fmt"

func main() {
    var N int
    fmt.Scan(&N)

    count := 0
    for i := 5; N / i >= 1; i *= 5 {
        count += N / i
    }

    fmt.Println(count)
}
