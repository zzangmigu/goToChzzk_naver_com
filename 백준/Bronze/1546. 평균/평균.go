package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var N int
	var max, sum int
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscanln(reader, &N)

	var scores = make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Fscan(reader, &scores[i])
		sum += scores[i]
		if max < scores[i] {
			max = scores[i]
		}
	}

	fmt.Println(float64(sum) / float64(N) / float64(max) * 100.0)
}
