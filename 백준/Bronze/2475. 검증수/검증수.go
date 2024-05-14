package main

import (
	"fmt"
)

func main() {
	var A, B, C, D, E, verify int
	fmt.Scan(&A, &B, &C, &D, &E)

	verify = ((A*A + B*B + C*C + D*D + E*E) % 10)
	fmt.Print(verify)

}
