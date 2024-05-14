package main

import (
	
	"fmt"
	"sort"

	)
	
func main() {
	var num, insertNum int
	fmt.Scanln(&num)
	var numArr []int = make([]int, num)
	for i := 0; i < num ; i++ {
		fmt.Scanln(&insertNum)
		numArr[i] = insertNum
	}
	sort.Ints(numArr)
	for _, v := range numArr {
		fmt.Println(v)
	}
}