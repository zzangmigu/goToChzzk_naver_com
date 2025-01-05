package main

import (
	"fmt"
)

func main() {
	var input rune
	fmt.Scanf("%c", &input)

	// 한글 유니코드 범위: U+AC00 (가) ~ U+D7A3 (힣)
	base := int('가')
	index := int(input) - base + 1

	fmt.Println(index)
}
