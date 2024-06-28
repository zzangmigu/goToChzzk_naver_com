func solution(my_string string, overwrite_string string, s int) string {
    var temp, converted string
    temp = my_string[:s]
    converted = temp + overwrite_string
    if len(converted) < len(my_string) {
        temp2 := my_string[len(converted):]
        converted += temp2
    }
    return converted
}