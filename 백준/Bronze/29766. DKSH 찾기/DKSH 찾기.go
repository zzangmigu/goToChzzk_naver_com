package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	// "DKSH"가 몇 번 나타나는지 세기
	count := strings.Count(input, "DKSH")

	fmt.Println(count)
}
