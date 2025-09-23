package main

import "fmt"

func main() {
	var N int
	fmt.Scan(&N)
	for i := 0; i < N; i++ {
		var x float64
		fmt.Scan(&x)
		fmt.Printf("$%.2f\n", x-x*0.2)
	}
}
