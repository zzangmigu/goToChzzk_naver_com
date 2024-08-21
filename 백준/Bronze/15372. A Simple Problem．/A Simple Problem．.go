package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var T int
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	fmt.Fscan(reader, &T)

	for i := 0; i < T; i++ {
		var N int
		fmt.Fscan(reader, &N)
		fmt.Fprintln(writer, N*N)
	}
}
