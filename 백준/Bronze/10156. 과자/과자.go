package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	var snack, number, money, minus int
	fmt.Fscan(reader, &snack, &number, &money)
	minus = snack*number - money
	if minus < 0 {
		fmt.Fprint(writer, 0)
	} else {
		fmt.Fprint(writer, minus)

	}
}
