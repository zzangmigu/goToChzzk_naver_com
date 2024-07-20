package main

import (
	"fmt"
	"strconv"
)

func reverseNumber(n int) int {
	str := strconv.Itoa(n)
	runes := []rune(str)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	reversed, _ := strconv.Atoi(string(runes))
	return reversed
}

func revAdd(x int, y int) int {
	revX := reverseNumber(x)
	revY := reverseNumber(y)
	sumRev := revX + revY
	result := reverseNumber(sumRev)
	return result
}

func main() {
	var x, y int
	fmt.Scanf("%d %d", &x, &y)
	fmt.Println(revAdd(x, y))
}
