package main

import (
	"fmt"
	"sort"
)

func main() {
	var n int
	fmt.Scan(&n)

	numbers := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&numbers[i])
	}

	gcd := numbers[0]
	for i := 1; i < n; i++ {
		gcd = findGCD(gcd, numbers[i])
	}

	divisors := findDivisors(gcd)
	sort.Ints(divisors)

	for _, divisor := range divisors {
		fmt.Println(divisor)
	}
}

func findGCD(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func findDivisors(n int) []int {
	divisors := []int{}
	for i := 1; i*i <= n; i++ {
		if n%i == 0 {
			divisors = append(divisors, i)
			if i != n/i {
				divisors = append(divisors, n/i)
			}
		}
	}
	return divisors
}
