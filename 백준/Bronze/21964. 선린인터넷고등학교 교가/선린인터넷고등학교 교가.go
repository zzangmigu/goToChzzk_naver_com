package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	var n int
	fmt.Fscan(reader, &n)
	var text string
	fmt.Fscan(reader, &text)
	if n <= 5 {
		fmt.Println(text)
	} else {
		fmt.Println(text[len(text)-5:])
	}
}
