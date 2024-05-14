package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	var num int
	fmt.Fscanln(reader, &num)

	for i := 0; i < num; i++ {
		fmt.Fprintln(writer, num-i)
	}
	writer.Flush()
}
