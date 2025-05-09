package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func reverse(arr []int, start, end int) {
	for i, j := start, end; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	// N과 M 입력 받기
	line, _ := reader.ReadString('\n')
	line = strings.TrimSpace(line)
	parts := strings.Split(line, " ")
	n, _ := strconv.Atoi(parts[0])
	m, _ := strconv.Atoi(parts[1])

	buckets := make([]int, n)
	for i := 0; i < n; i++ {
		buckets[i] = i + 1
	}

	for i := 0; i < m; i++ {
		line, _ = reader.ReadString('\n')
		line = strings.TrimSpace(line)
		parts := strings.Split(line, " ")
		start, _ := strconv.Atoi(parts[0])
		end, _ := strconv.Atoi(parts[1])

		reverse(buckets, start-1, end-1)
	}
	
	for i, ball := range buckets {
		fmt.Print(ball)
		if i < n-1 {
			fmt.Print(" ")
		}
	}
	fmt.Println()
}
