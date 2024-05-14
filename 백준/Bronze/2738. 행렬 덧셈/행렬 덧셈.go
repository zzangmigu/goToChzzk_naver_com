package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var N, M int
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	defer writer.Flush()

	fmt.Fscan(reader, &N, &M)

	var arr1 [][]int = make([][]int, N)
	var arr2 [][]int = make([][]int, N)

	for i := 0; i < N; i++ {
		arr1[i] = make([]int, M)
		for j := 0; j < M; j++ {
			fmt.Fscan(reader, &arr1[i][j])
		}
	}

	for i := 0; i < N; i++ {
		arr2[i] = make([]int, M)
		for j := 0; j < M; j++ {
			fmt.Fscan(reader, &arr2[i][j])
			arr1[i][j] += arr2[i][j]
			fmt.Fprintf(writer, "%d ", arr1[i][j])
		}

		fmt.Fprintln(writer, "")
	}

}
