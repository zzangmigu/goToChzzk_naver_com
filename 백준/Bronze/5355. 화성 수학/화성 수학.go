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
	t, _ := strconv.Atoi(scanner.Text())

	for i := 0; i < t; i++ {
		scanner.Scan()
		expression := scanner.Text()
		result := calculate(expression)
		fmt.Printf("%.2f\n", result)
	}

}

func calculate(expression string) float64 {
	parts := strings.Split(expression, " ")
	var result float64
	result, _ = strconv.ParseFloat(parts[0], 64)

	for i := 1; i < len(parts); i++ {
		switch parts[i] {
		case "@":
			result *= 3
		case "%":
			result += 5
		case "#":
			result -= 7
		}
	}

	return result
}
