package main

import (
	
	"fmt"
	"sort"	
	"bufio"
	"os"
	)
	


func main() {
	var N int
	r := bufio.NewReader(os.Stdin)
	fmt.Fscan(r, &N)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	var numArr []int = make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Fscan(r, &numArr[i])
	}

	sort.Ints(numArr)
	for i := 0; i < N; i++ {
		fmt.Fprintln(w, numArr[i])
	}
}