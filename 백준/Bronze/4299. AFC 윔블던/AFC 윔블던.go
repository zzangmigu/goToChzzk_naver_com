package main

import (
	"fmt"
)

func main() {
	var s, d int
	fmt.Scan(&s, &d)

	if (s+d)%2 != 0 || (s-d)%2 != 0 {
		fmt.Println(-1)
		return
	}

	x := (s + d) / 2
	y := (s - d) / 2

	if x < 0 || y < 0 {
		fmt.Println(-1)
	} else {
		if x >= y {
			fmt.Println(x, y)
		} else {
			fmt.Println(y, x)
		}
	}
}
