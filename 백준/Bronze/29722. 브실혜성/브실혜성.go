package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	today := scanner.Text()
	scanner.Scan()
	strN := scanner.Text()

	yy, MM, DD := strings.Split(today, "-")[0], strings.Split(today, "-")[1], strings.Split(today, "-")[2]

	intY, _ := strconv.Atoi(yy)
	intM, _ := strconv.Atoi(MM)
	intD, _ := strconv.Atoi(DD)
	intN, _ := strconv.Atoi(strN)

	convertedToday := (360 * (intY - 1)) + (30 * (intM - 1)) + (intD - 1)
	dreamDays := convertedToday + intN

	calY := dreamDays/360 + 1
	calM := (dreamDays%360)/30 + 1
	calD := (dreamDays%360)%30 + 1

	var newM, newD string

	if calM < 10 {
		newM = "0" + strconv.Itoa(calM)
	} else {
		newM = strconv.Itoa(calM)
	}
	if calD < 10 {
		newD = "0" + strconv.Itoa(calD)
	} else {
		newD = strconv.Itoa(calD)
	}

	fmt.Printf("%d-%s-%s", calY, newM, newD)

}
