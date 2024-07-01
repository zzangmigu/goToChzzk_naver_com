package main

import (
	"bufio"
	"fmt"
	"os"
)

func isVPS(s string) bool {
	stack := []rune{}

	for _, char := range s {
		if char == '(' {
			stack = append(stack, char)
		} else if char == ')' {
			if len(stack) == 0 {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}

	return len(stack) == 0
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var num int
	fmt.Fscanln(reader, &num)

	for i := 0; i < num; i++ {
		var input string
		fmt.Fscanln(reader, &input)

		if isVPS(input) {
			fmt.Fprintln(writer, "YES")
		} else {
			fmt.Fprintln(writer, "NO")
		}
	}
}
