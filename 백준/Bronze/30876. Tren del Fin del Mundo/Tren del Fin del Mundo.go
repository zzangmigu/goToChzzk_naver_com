package main

import "fmt"

func main() {
	var N, Xsouth, Ysouth int
	fmt.Scan(&N)
	fmt.Scan(&Xsouth, &Ysouth)
	for i := 1; i < N; i++ {
		var x, y int
		fmt.Scan(&x, &y)
		if y < Ysouth {
			Xsouth = x
			Ysouth = y
		}

	}
	fmt.Println(Xsouth, Ysouth)
}
