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

	var jinho string
	var n, count int
	fmt.Fscanln(reader, &jinho)
	fmt.Fscanln(reader, &n)
	mbti := make([]string, n)
	for i := 0; i < n; i++ {
		fmt.Fscanln(reader, &mbti[i])
		if jinho == mbti[i] {
			count++
		}
	}
	fmt.Fprintln(writer, count)
}
