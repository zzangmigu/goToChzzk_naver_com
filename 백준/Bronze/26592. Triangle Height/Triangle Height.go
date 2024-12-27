package main

import "fmt"

func main() {
	var T int
	fmt.Scan(&T)
	for i := 0; i < T; i++ {
		var num1, num2 float32
		fmt.Scan(&num1, &num2)
		fmt.Printf("The height of the triangle is %.2f units\n", (num1*2)/num2)
	}

}
