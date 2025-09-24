package main

import "fmt"

func main() {
	var K, N, T, sumT int
	var Z string
	fmt.Scan(&K)
	fmt.Scan(&N)

	for i := 0; i < N; i++ {
		fmt.Scan(&T, &Z)

		sumT += T
		if sumT >= 210 {
			fmt.Println(K)
			break

		}

		if Z == "T" {
			K++
			if K > 8 {
				K = 1
			}
		}

	}

}
