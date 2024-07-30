package main

import (
	"fmt"
)

func main() {
	var name string
	var age, weight int
	for true {
		fmt.Scan(&name, &age, &weight)
		if name == "#" && age == 0 && weight == 0 {
			break
		}
		if age > 17 || weight >= 80 {
			fmt.Println(name, "Senior")
		} else {
			fmt.Println(name, "Junior")
		}
	}
}
