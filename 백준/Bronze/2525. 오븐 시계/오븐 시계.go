package main

import "fmt"

func main() {
	var nowH, nowM, needTime int
	fmt.Scanf("%d %d\n%d", &nowH, &nowM, &needTime)
	if nowM+needTime >= 60 {
		if (nowM+needTime)%60 != 0 {
			nowH += (nowM + needTime) / 60
			nowM = (nowM + needTime) % 60
			if nowH >= 24 {
				nowH = nowH - 24
			}

			fmt.Println(nowH, nowM)
		} else {

			nowH += (nowM + needTime) / 60
			nowM = (nowM + needTime) % 60
			if nowH >= 24 {
				nowH = nowH - 24
			}
			fmt.Println(nowH, nowM)
		}

	} else {
		if nowH >= 24 {
			nowH = nowH - 24
		}
		fmt.Println(nowH, nowM+needTime)
	}
}
