func largestOddNumber(num string) string {
    for len(num) > 0{
        x, _ := strconv.Atoi(string(num[len(num)-1]))
        if x%2 ==1 {
            return num
        }
        num = num[:len(num)-1]
    }
    return ""
}