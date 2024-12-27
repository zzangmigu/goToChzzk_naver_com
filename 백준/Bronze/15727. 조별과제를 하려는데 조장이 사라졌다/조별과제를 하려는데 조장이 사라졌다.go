package main

import "fmt"

func main() {
	var input int
	fmt.Scan(&input)
	if input%5 == 0 {
		fmt.Print(input / 5)
	} else {
		fmt.Print(input/5 + 1)
	}

}
