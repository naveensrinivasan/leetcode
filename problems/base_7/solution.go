func convertToBase7(num int) string {
    if num == 0{
        return "0"
    }
    
    x := []string{}
    wasLessThanZero := false
    
    if num < 0{
        num = num * -1
        wasLessThanZero = true
    } 
    
    for num > 0{
        x = append([]string{strconv.Itoa(num%7)},x...)
        num /= 7
    }
    
    if wasLessThanZero{
        x = append([]string{"-"},x...)
    }
    return strings.Join(x,"")
}