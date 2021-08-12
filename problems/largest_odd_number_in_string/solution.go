func largestOddNumber(num string) string {
    result := ""
    for len(num) > 0{
        y := num[len(num)-1]
        x, _ := strconv.Atoi(string(y))
        if x%2 ==1 {
            return num
        }
        num = num[:len(num)-1]
    }
    return result
}