package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var word string
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscan(reader, &word)

	var letters = make(map[uint8]int)
	for i := 0; i < 26; i++ {
		letters[uint8(i)+65] = 0
	}

	for i := 0; i < len(word); i++ {
		ascii := word[i]
		if ascii > 90 {
			ascii -= 32
		}
		letters[ascii]++
	}

	var maxVal = -2
	var maxKey string
	for key, val := range letters {
		if val > maxVal {
			maxVal = val
			maxKey = string(key)
		} else if val == maxVal {
			maxKey = "?"
		}
	}

	fmt.Println(maxKey)
}
