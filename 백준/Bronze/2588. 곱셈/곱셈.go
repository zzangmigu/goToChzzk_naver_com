package main

import "fmt"

func main() {
	var a, b, s1, s2, s3, t1, t2, t3, total int
	fmt.Scanf("%d\n %d", &a, &b)
	s1 = b / 100
	s2 = (b % 100) / 10
	s3 = (b % 10)
	t1, t2, t3 = a*s3, a*s2, a*s1
	total = t1 + t2*10 + t3*100
	fmt.Printf("%d\n%d \n%d \n%d", t1, t2, t3, total)
}
