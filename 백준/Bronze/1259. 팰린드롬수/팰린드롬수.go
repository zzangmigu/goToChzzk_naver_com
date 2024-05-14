package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {

	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	var N int
	for true {
		fmt.Fscan(r, &N)
		if N == 0 {
			break
		}
		if isPalindrome(N) {
			fmt.Fprintln(w, "yes")
		} else {
			fmt.Fprintln(w, "no")
		}
	}
}

func isPalindrome(n int) bool {
	s := strconv.Itoa(n)
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-i-1] {
			return false
		}
	}
	return true
}