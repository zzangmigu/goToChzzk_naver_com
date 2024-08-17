package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var N int
	fmt.Fscan(reader, &N)

	matrix := make([][]int, N)
	for i := 0; i < N; i++ {
		matrix[i] = make([]int, N)
		for j := 0; j < N; j++ {
			fmt.Fscan(reader, &matrix[i][j])
		}
	}

	// Calculate the magic sum
	magicSum := N * (N*N + 1) / 2

	// Check rows
	for i := 0; i < N; i++ {
		sum := 0
		for j := 0; j < N; j++ {
			sum += matrix[i][j]
		}
		if sum != magicSum {
			fmt.Println("FALSE")
			return
		}
	}

	// Check columns
	for j := 0; j < N; j++ {
		sum := 0
		for i := 0; i < N; i++ {
			sum += matrix[i][j]
		}
		if sum != magicSum {
			fmt.Println("FALSE")
			return
		}
	}

	// Check main diagonal
	sum := 0
	for i := 0; i < N; i++ {
		sum += matrix[i][i]
	}
	if sum != magicSum {
		fmt.Println("FALSE")
		return
	}

	// Check anti-diagonal
	sum = 0
	for i := 0; i < N; i++ {
		sum += matrix[i][N-i-1]
	}
	if sum != magicSum {
		fmt.Println("FALSE")
		return
	}

	// Check if all numbers from 1 to N*N are present
	expectedNumbers := make(map[int]bool)
	for i := 1; i <= N*N; i++ {
		expectedNumbers[i] = true
	}
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			if _, exists := expectedNumbers[matrix[i][j]]; !exists {
				fmt.Println("FALSE")
				return
			}
			delete(expectedNumbers, matrix[i][j])
		}
	}

	if len(expectedNumbers) > 0 {
		fmt.Println("FALSE")
		return
	}

	fmt.Println("TRUE")
}
