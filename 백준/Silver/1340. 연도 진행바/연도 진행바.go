package main

import (
	"fmt"
	"time"
)

func main() {
	// 입력을 받습니다.
	var monthStr string
	var day, year, hour, minute int
	fmt.Scanf("%s %d, %d %d:%d", &monthStr, &day, &year, &hour, &minute)

	// 월 이름을 숫자로 변환합니다.
	monthMap := map[string]time.Month{
		"January":   time.January,
		"February":  time.February,
		"March":     time.March,
		"April":     time.April,
		"May":       time.May,
		"June":      time.June,
		"July":      time.July,
		"August":    time.August,
		"September": time.September,
		"October":   time.October,
		"November":  time.November,
		"December":  time.December,
	}

	month := monthMap[monthStr]

	// 현재 날짜와 시간을 time.Time 타입으로 변환합니다.
	currentTime := time.Date(year, month, day, hour, minute, 0, 0, time.UTC)

	// 해당 연도의 첫 번째 날과 마지막 날을 계산합니다.
	startOfYear := time.Date(year, time.January, 1, 0, 0, 0, 0, time.UTC)
	endOfYear := time.Date(year+1, time.January, 1, 0, 0, 0, 0, time.UTC)

	// 전체 연도의 시간과 현재까지의 시간을 초 단위로 계산합니다.
	totalSeconds := endOfYear.Sub(startOfYear).Seconds()
	elapsedSeconds := currentTime.Sub(startOfYear).Seconds()

	// 퍼센트 계산
	percentage := (elapsedSeconds / totalSeconds) * 100

	// 결과 출력
	fmt.Printf("%.15f\n", percentage)
}
