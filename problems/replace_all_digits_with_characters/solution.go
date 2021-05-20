func replaceDigits(s string) string {
    result := make([]rune,len(s))
    for i:=0;i<len(s);i++{
        if (i+1)%2 ==0{
            num,_ := strconv.Atoi(string(s[i]))
            result[i] = rune(rune(s[i-1])+rune(num))
        }else{
            result[i] = rune(s[i])
        }  
    }
    return string(result)
}