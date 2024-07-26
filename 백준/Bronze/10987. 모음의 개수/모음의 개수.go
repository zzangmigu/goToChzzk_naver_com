package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// 표준 입력을 통해 데이터 받기
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	vowels := "aeiou"
	count := 0
	var word string
	fmt.Fscan(reader, &word)

	for _, char := range word {
		if strings.ContainsRune(vowels, char) {
			count++
		}
	}
	fmt.Println(count)
}
