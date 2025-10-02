package main

import "fmt"

func main() {
	var T, yTele, mTele int
	fmt.Scanln(&T)
	for i := 0; i < T; i++ {
		var duration int
		fmt.Scan(&duration)

		yTele += (duration/30 + 1) * 10
		mTele += (duration/60 + 1) * 15
	}
	if yTele < mTele {
		fmt.Println("Y", yTele)
	} else if yTele > mTele {
		fmt.Println("M", mTele)
	} else {
		fmt.Println("Y M", yTele)
	}
}
