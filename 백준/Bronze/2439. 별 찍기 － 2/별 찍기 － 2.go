package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var N int
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	fmt.Fscanln(reader, &N)
	defer writer.Flush()

	for i := 0; i < N; i++ {
		for j := N - i - 1; j > 0; j-- {
			fmt.Fprint(writer, " ")
		}
		for k := 0; k < i+1; k++ {
			fmt.Fprint(writer, "*")
		}
		fmt.Fprint(writer, "\n")
	}
}
