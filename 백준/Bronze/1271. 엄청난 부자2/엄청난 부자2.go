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

	var nStr, mStr string
	fmt.Fscan(reader, &nStr, &mStr)

	n := new(big.Int)
	m := new(big.Int)
	n.SetString(nStr, 10)
	m.SetString(mStr, 10)

	quotient := new(big.Int)
	remainder := new(big.Int)

	quotient.DivMod(n, m, remainder)

	fmt.Fprintln(writer, quotient.String())
	fmt.Fprintln(writer, remainder.String())
}
