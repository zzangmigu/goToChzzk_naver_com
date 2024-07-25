package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	// 표준 입력을 통해 데이터 받기
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	// 단어의 개수 입력 받기
	var n int
	fmt.Fscanln(reader, &n)

	// 단어를 저장할 맵과 슬라이스
	uniqueWords := make(map[string]bool)
	words := make([]string, 0, n)

	// 단어 입력받기 및 중복 제거
	for i := 0; i < n; i++ {
		var word string
		fmt.Fscanln(reader, &word)
		if !uniqueWords[word] {
			uniqueWords[word] = true
			words = append(words, word)
		}
	}
	// 정렬: 길이가 짧은 것부터, 길이가 같으면 사전 순으로
	sort.Slice(words, func(i, j int) bool {
		if len(words[i]) == len(words[j]) {
			return words[i] < words[j]
		}
		return len(words[i]) < len(words[j])
	})

	// 결과 출력
	for _, word := range words {
		fmt.Fprintln(writer, word)
	}
}
