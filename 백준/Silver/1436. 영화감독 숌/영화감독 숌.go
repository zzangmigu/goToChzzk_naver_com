package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var n int
	fmt.Scan(&n)

	count := 0
	number := 665

	for {
		number++
		if strings.Contains(strconv.Itoa(number), "666") {
			count++
		}
		if count == n {
			fmt.Println(number)
			break
		}
	}
}
