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
		for j := 0; j < i+1; j++ {
			fmt.Fprint(writer, "*")
		}
		fmt.Fprint(writer, "\n")
	}
}
