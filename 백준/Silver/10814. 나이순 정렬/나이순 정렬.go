package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)



type member struct {
	name     string
	age      int
	index int
}

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()
	
	var N int
	fmt.Fscanf(r, "%d\n", &N)
	m := []member{}
	for i := 0; i < N; i++ {
		var tm member
		s, _ := r.ReadString('\n')
		s = strings.TrimSuffix(s, "\n")
		split := strings.Split(s, " ")
		tm.age, _ = strconv.Atoi(split[0])
		tm.name = split[1]
		tm.index = i
		m = append(m, tm)
	}

	sort.SliceStable(m, func(i, j int) bool {
		return m[i].index < m[j].index
	})

	sort.SliceStable(m, func(i, j int) bool {
		return m[i].age < m[j].age
	})

	for _, v := range m {
		fmt.Fprintln(w, v.age, v.name)
	}
}