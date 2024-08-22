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
	var n int
	fmt.Fscan(reader, &n)

	volumes := make([]float64, n)
	var totalVolume float64

	for i := 0; i < n; i++ {
		var c float64
		fmt.Fscan(reader, &c)
		volume := math.Pow(c, 3)
		volumes[i] = volume
		totalVolume += volume
	}

	sideLength := math.Pow(totalVolume, 1.0/3.0)
	fmt.Fprintf(writer, "%.9f\n", sideLength)
}
