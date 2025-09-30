package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	for i := 0; i < 3; i++ {
		var N int
		fmt.Fscanln(reader, &N)
		sum := new(big.Int)

		for j := 0; j < N; j++ {
			var numStr string
			fmt.Fscanln(reader, &numStr)
			num := new(big.Int)
			num.SetString(numStr, 10)
			sum.Add(sum, num)

		}
		sign := sum.Sign()
		if sign < 0 {
			fmt.Println("-")
		} else if sign > 0 {
			fmt.Println("+")
		} else {
			fmt.Println("0")
		}

	}

}
