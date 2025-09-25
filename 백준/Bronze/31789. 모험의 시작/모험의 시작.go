package main

import "fmt"

func main() {
	var N, X, S int
	canWin := false

	fmt.Scan(&N)
	fmt.Scan(&X, &S)

	for i := 0; i < N; i++ {
		var c, p int
		fmt.Scan(&c, &p)
		if c <= X && p > S {
			canWin = true
			break
		}
	}

	if canWin {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}

}
