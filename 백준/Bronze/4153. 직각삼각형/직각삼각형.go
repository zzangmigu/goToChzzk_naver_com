package main

import (
	"bufio"
	"fmt"
	"os"
	"math"
)

func main() {
	var a, b, c float64
	reader := bufio.NewReader(os.Stdin)

	for true {
		fmt.Fscanln(reader, &a, &b, &c)

		if a == 0 && b == 0 && c == 0 {
			break
		}
		aPow := math.Pow(a, 2)
		bPow := math.Pow(b, 2)
		cPow := math.Pow(c, 2)

		if aPow > bPow && aPow > cPow{
			if bPow + cPow == aPow {
				fmt.Println("right")
			} else {
				fmt.Println("wrong")
			}
		} else if bPow > aPow && bPow > cPow {
			if aPow + cPow == bPow {
				fmt.Println("right")
			} else {
				fmt.Println("wrong")
			}
		} else if cPow > aPow && cPow > bPow {
			if aPow + bPow == cPow {
				fmt.Println("right")
			} else {
				fmt.Println("wrong")
			}
		}
	}
}