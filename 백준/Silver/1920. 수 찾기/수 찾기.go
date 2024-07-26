package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	// 첫 번째 줄에서 N을 읽습니다.
	nLine, _ := reader.ReadString('\n')
	n, _ := strconv.Atoi(strings.TrimSpace(nLine))

	// 다음 줄에서 N개의 정수를 읽습니다.
	aLine, _ := reader.ReadString('\n')
	aStr := strings.Fields(aLine)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		a[i], _ = strconv.Atoi(aStr[i])
	}

	// 세 번째 줄에서 M을 읽습니다.
	mLine, _ := reader.ReadString('\n')
	m, _ := strconv.Atoi(strings.TrimSpace(mLine))

	// 마지막 줄에서 M개의 정수를 읽습니다.
	mLine, _ = reader.ReadString('\n')
	mStr := strings.Fields(mLine)
	mNumbers := make([]int, m)
	for i := 0; i < m; i++ {
		mNumbers[i], _ = strconv.Atoi(mStr[i])
	}

	// 배열 A를 정렬합니다.
	sort.Ints(a)

	// 각 숫자에 대해 이진 탐색을 수행합니다.
	for _, num := range mNumbers {
		if exists(a, num) {
			fmt.Fprintln(writer, "1")
		} else {
			fmt.Fprintln(writer, "0")
		}
	}
}

// 존재하는지 여부를 찾는 함수
func exists(arr []int, target int) bool {
	index := sort.SearchInts(arr, target)
	return index < len(arr) && arr[index] == target
}
