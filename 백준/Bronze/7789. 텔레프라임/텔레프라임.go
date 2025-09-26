package main

import (
	"fmt"
	"math/big"
)

func main() {
	var oriNumber, appendN int
	fmt.Scan(&oriNumber, &appendN)

	newN := appendN*1000000 + oriNumber

	BigoriNumber := big.NewInt(int64(oriNumber))
	BigappendN := big.NewInt(int64(newN))

	isPrimeOri := BigoriNumber.ProbablyPrime(0)
	isPrimeAppendN := BigappendN.ProbablyPrime(0)
	if isPrimeOri == true && isPrimeAppendN == true {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
