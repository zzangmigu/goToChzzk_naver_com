package main

import (
	"fmt"
)

func findOppositeMBTI(mbti string) string {
	opposite := map[byte]byte{
		'E': 'I',
		'I': 'E',
		'S': 'N',
		'N': 'S',
		'T': 'F',
		'F': 'T',
		'J': 'P',
		'P': 'J',
	}

	result := make([]byte, 4)
	for i := 0; i < 4; i++ {
		result[i] = opposite[mbti[i]]
	}

	return string(result)
}

func main() {
	var mbti string
	fmt.Scan(&mbti)
	fmt.Println(findOppositeMBTI(mbti))
}
