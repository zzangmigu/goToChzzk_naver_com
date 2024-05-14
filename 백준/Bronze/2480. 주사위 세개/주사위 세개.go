package main

import "fmt"

func main() {
	var d1, d2, d3, prize int
	fmt.Scanln(&d1, &d2, &d3)
	if d1 == d2 && d2 == d3 && d1 == d3 {
		prize = 10000 + d1*1000
		fmt.Println(prize)
	} else if d1 == d2 {
		prize = 1000 + d1*100
		fmt.Println(prize)
	} else if d2 == d3 {
		prize = 1000 + d2*100
		fmt.Println(prize)
	} else if d1 == d3 {
		prize = 1000 + d3*100
		fmt.Println(prize)
	} else if d1 > d2 && d1 > d3 {
		prize = d1 * 100
		fmt.Println(prize)
	} else if d2 > d3 && d2 > d1 {
		prize = d2 * 100
		fmt.Println(prize)
	} else if d3 > d1 && d3 > d2 {
		prize = d3 * 100
		fmt.Println(prize)
	}
}
