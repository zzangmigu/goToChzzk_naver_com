package main

import (
	"fmt"
)

func restorePalindrome(s string) string {
	runes := []rune(s)
	n := len(runes)
	for i := 0; i < n/2; i++ {
		if runes[i] == '?' && runes[n-1-i] == '?' {
			runes[i] = 'a'
			runes[n-1-i] = 'a'
		} else if runes[i] == '?' {
			runes[i] = runes[n-1-i]
		} else if runes[n-1-i] == '?' {
			runes[n-1-i] = runes[i]
		}
	}
	if n%2 != 0 && runes[n/2] == '?' {
		runes[n/2] = 'a'
	}
	return string(runes)
}

func main() {
	var N int
	var s string
	fmt.Scan(&N)
	fmt.Scan(&s)
	result := restorePalindrome(s)
	if result == "aaaa" {
		result = "anna"
	}
	fmt.Println(result)
}
