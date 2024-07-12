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

	nStr, _ := reader.ReadString('\n')
	n, _ := strconv.Atoi(strings.TrimSpace(nStr))

	arrStr, _ := reader.ReadString('\n')
	arrStr = strings.TrimSpace(arrStr)
	arrStrs := strings.Split(arrStr, " ")
	arr := make([]int, n)
	for i := range arrStrs {
		arr[i], _ = strconv.Atoi(arrStrs[i])
	}

	xStr, _ := reader.ReadString('\n')
	x, _ := strconv.Atoi(strings.TrimSpace(xStr))

	sort.Ints(arr)

	left, right := 0, n-1
	count := 0

	for left < right {
		sum := arr[left] + arr[right]
		if sum == x {
			count++
			left++
			right--
		} else if sum < x {
			left++
		} else {
			right--
		}
	}

	fmt.Fprintln(writer, count)
}
