package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var N, count int
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscanln(reader, &N)

	var otN = N
	for true {

		var sum int
		if otN/10 == 0 {
			sum = otN
		} else {
			sum = otN/10 + otN%10
		}
		otN = otN%10*10 + sum%10
		count++
		if otN == N {
			break
		}

	}
	fmt.Println(count)
}
