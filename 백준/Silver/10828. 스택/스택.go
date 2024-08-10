package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Stack structure
type Stack struct {
	items []int
}

// Push adds an item to the stack
func (s *Stack) Push(item int) {
	s.items = append(s.items, item)
}

// Pop removes and returns the top item of the stack, or returns -1 if the stack is empty
func (s *Stack) Pop() int {
	if len(s.items) == 0 {
		return -1
	}
	item := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return item
}

// Size returns the number of items in the stack
func (s *Stack) Size() int {
	return len(s.items)
}

// Empty returns 1 if the stack is empty, 0 otherwise
func (s *Stack) Empty() int {
	if len(s.items) == 0 {
		return 1
	}
	return 0
}

// Top returns the top item of the stack without removing it, or returns -1 if the stack is empty
func (s *Stack) Top() int {
	if len(s.items) == 0 {
		return -1
	}
	return s.items[len(s.items)-1]
}

func main() {
	var n int
	fmt.Scan(&n)

	stack := Stack{}
	scanner := bufio.NewScanner(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	for i := 0; i < n; i++ {
		scanner.Scan()
		command := scanner.Text()
		if strings.HasPrefix(command, "push") {
			parts := strings.Split(command, " ")
			value, _ := strconv.Atoi(parts[1])
			stack.Push(value)
		} else if command == "pop" {
			fmt.Fprintln(writer, stack.Pop())
		} else if command == "size" {
			fmt.Fprintln(writer, stack.Size())
		} else if command == "empty" {
			fmt.Fprintln(writer, stack.Empty())
		} else if command == "top" {
			fmt.Fprintln(writer, stack.Top())
		}
	}
}
