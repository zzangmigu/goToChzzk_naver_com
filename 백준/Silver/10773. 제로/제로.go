package main

import (
	"bufio"
    "fmt"
    "os"
    "strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	K, _ := strconv.Atoi(scanner.Text())
	stack := []int{}
	for i := 0; i < K; i++ {
		scanner.Scan()
		num, _ := strconv.Atoi(scanner.Text())
		if num == 0 {
			// 스택이 비어 있지 않을 때만 팝
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			}
		} else {
			stack = append(stack, num)
		}
	}

	// 스택에 남아있는 숫자들의 합을 계산
	sum := 0
	for _, val := range stack {
		sum += val
	}
	fmt.Println(sum)
}