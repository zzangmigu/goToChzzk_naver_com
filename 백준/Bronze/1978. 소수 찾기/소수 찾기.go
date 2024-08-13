package main

import (
    "fmt"
    "math"
)

// Function to check if a number is prime
func isPrime(n int) bool {
    if n <= 1 {
        return false
    }
    for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
        if n%i == 0 {
            return false
        }
    }
    return true
}

func main() {
    var N int
    fmt.Scan(&N)

    count := 0
    for i := 0; i < N; i++ {
        var num int
        fmt.Scan(&num)
        if isPrime(num) {
            count++
        }
    }

    fmt.Println(count)
}
