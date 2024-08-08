package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
)

func main() {
    reader := bufio.NewReader(os.Stdin)
    
    // 깊이 측정값을 저장할 슬라이스
    depths := make([]int, 4)
    
    // 깊이 측정값 입력 받기
    for i := 0; i < 4; i++ {
        input, _ := reader.ReadString('\n')
        input = input[:len(input)-1]
        depth, _ := strconv.Atoi(input)
        depths[i] = depth
    }
    
    // 상태 판단하기
    if depths[0] < depths[1] && depths[1] < depths[2] && depths[2] < depths[3] {
        fmt.Println("Fish Rising")
    } else if depths[0] > depths[1] && depths[1] > depths[2] && depths[2] > depths[3] {
        fmt.Println("Fish Diving")
    } else if depths[0] == depths[1] && depths[1] == depths[2] && depths[2] == depths[3] {
        fmt.Println("Fish At Constant Depth")
    } else {
        fmt.Println("No Fish")
    }
}
