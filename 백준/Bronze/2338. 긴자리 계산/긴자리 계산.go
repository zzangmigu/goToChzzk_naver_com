package main

import (
	"fmt"
	"math/big"
	"os"
)

func main() {
	var A, B big.Int

	_, err := fmt.Fscanf(os.Stdin, "%s\n%s\n", &A, &B)
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	sum := new(big.Int).Add(&A, &B)

	diff := new(big.Int).Sub(&A, &B)

	prod := new(big.Int).Mul(&A, &B)

	fmt.Println(sum)
	fmt.Println(diff)
	fmt.Println(prod)
}
