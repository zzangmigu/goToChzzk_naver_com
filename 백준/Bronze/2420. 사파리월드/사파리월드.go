package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var N, M int

	fmt.Fscan(reader, &N, &M)
	fmt.Print(int(math.Abs((float64)(N - M))))

}
