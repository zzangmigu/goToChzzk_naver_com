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

	var n int
	fmt.Fscanln(reader, &n)

	var cards = map[int]int{}
	for i := 0; i < n; i++ {
		var input int
		if i == n-1 {
			fmt.Fscanln(reader, &input)
		} else {
			fmt.Fscan(reader, &input)
		}
		cards[input]++
	}

	var m int
	fmt.Fscanln(reader, &m)

	for i := 0; i < m; i++ {
		var number int
		fmt.Fscanf(reader, "%d ", &number)
		fmt.Fprintf(writer, "%s ", hasCard(cards, number))
	}
}

func hasCard(cards map[int]int, number int) string {
	if cards[number] != 0 {
		return "1"
	}
	return "0"
}
