package main

import "fmt"

func main() {
	var bal, money int
	var transType string
	for {
		fmt.Scanf("%d %s %d\n", &bal, &transType, &money)
		if bal == 0 && transType == "W" && money == 0 {
			break
		}

		if transType == "W" {
			if bal-money < -200 {
				fmt.Println("Not allowed")
			} else {
				fmt.Println(bal - money)
			}
		} else if transType == "D" {
			fmt.Println(money + bal)
		}
	}

}
