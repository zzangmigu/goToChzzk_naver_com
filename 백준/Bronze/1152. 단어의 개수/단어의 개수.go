package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

func main() {

	input := bufio.NewReader(os.Stdin)
	word, _ := input.ReadString('\n')

	fmt.Print(WordCount(word))

}

func WordCount(value string) int {
	re := regexp.MustCompile(`[\S]+`)
	results := re.FindAllString(value, -1)
	return len(results)

}
