package main

import (
	"fmt"
)

func main() {
	var score string
	fmt.Scan(&score)

	if score == "A+" {
		fmt.Print(4.3)
	} else if score == "A0" {
		fmt.Printf("%.1f", 4.0)
	} else if score == "A-" {
		fmt.Print(3.7)
	} else if score == "B+" {
		fmt.Print(3.3)
	} else if score == "B0" {
		fmt.Printf("%.1f", 3.0)
	} else if score == "B-" {
		fmt.Print(2.7)
	} else if score == "C+" {
		fmt.Print(2.3)
	} else if score == "C0" {
		fmt.Printf("%.1f", 2.0)
	} else if score == "C-" {
		fmt.Print(1.7)
	} else if score == "D+" {
		fmt.Print(1.3)
	} else if score == "D0" {
		fmt.Printf("%.1f", 1.0)
	} else if score == "D-" {
		fmt.Print(0.7)
	} else if score == "F" {
		fmt.Printf("%.1f", 0.0)
	}

}
