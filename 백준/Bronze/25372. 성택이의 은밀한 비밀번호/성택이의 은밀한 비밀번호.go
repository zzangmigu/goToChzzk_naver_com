package main

import "fmt"

func main() {
	var input int
	fmt.Scan(&input)
	for i := 0; i < input; i++ {
		var temp string
		fmt.Scan(&temp)
		if len(temp) >= 6 && len(temp) <= 9 {
			fmt.Println("yes")
		} else {
			fmt.Println("no")
		}

	}

}
