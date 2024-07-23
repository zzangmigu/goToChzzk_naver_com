package main

import (
    "fmt"
)

func main() {
    var minGukInfo, minGukMath, minGukSci, minGukEng int
    var manSeInfo, manSeMath, manSeSci, manSeEng int
    
    fmt.Scanf("%d %d %d %d", &minGukInfo, &minGukMath, &minGukSci, &minGukEng)
    fmt.Scanf("%d %d %d %d", &manSeInfo, &manSeMath, &manSeSci, &manSeEng)
    
    sumMinGuk := minGukInfo + minGukMath + minGukSci + minGukEng
    sumManSe := manSeInfo + manSeMath + manSeSci + manSeEng
    
    if sumMinGuk >= sumManSe {
        fmt.Println(sumMinGuk)
    } else {
        fmt.Println(sumManSe)
    }
}
