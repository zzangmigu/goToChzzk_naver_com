package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var count, bigger int
	sequence := make([]int, 9)
	for i := range sequence {
		fmt.Fscan(reader, &sequence[i])
		if sequence[i] > bigger {
			bigger = sequence[i]
			count = i
		}

	}

	fmt.Print(bigger, "\n", count+1)
}
