package main

import (
	
	"fmt"
	"sort"
	)
	


func main() {
	var insertNum, sum int
	var numArr []int = make([]int, 5)
	for i := 0; i < 5; i++ {
		fmt.Scanln(&insertNum)
		numArr[i] = insertNum
		sum += insertNum
	}
	
	fmt.Println(sum/5)
	sort.Ints(numArr)
	fmt.Println(numArr[2])

}