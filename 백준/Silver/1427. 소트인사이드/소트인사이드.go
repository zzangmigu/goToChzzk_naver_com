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
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()
	fmt.Fscan(r, &N)

	var numArr []int
	for i := N; i != 0; i/=10 {
		numArr = append(numArr, i%10)
	}

	sort.Sort(sort.Reverse(sort.IntSlice(numArr)))
	for i:= 0; i < len(numArr); i++ {
		fmt.Fprint(w, numArr[i])
	}
}