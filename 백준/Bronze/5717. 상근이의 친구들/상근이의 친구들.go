package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// 표준 입력을 통해 데이터 받기
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var a, b int
	for true {
		fmt.Fscanln(reader, &a, &b)
		if a == 0 && b == 0 {
			break
		} else {
			fmt.Fprintln(writer, a+b)
		}
	}
}
