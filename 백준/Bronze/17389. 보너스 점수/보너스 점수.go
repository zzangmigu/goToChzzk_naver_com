package main

import (
    "bufio"
    "fmt"
    "os"
)

func main() {
    reader := bufio.NewReader(os.Stdin)
    var n int
    fmt.Fscan(reader, &n)

    var s string
    fmt.Fscan(reader, &s)

    score := 0
    bonus := 0

    for i := 0; i < n; i++ {
        if s[i] == 'O' {
            score += (i + 1) + bonus
            bonus++
        } else {
            bonus = 0
        }
    }

    fmt.Println(score)
}
