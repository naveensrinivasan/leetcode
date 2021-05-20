func freqAlphabets(s string) string {
    result := []rune{}
    for i:=len(s)-1;i>=0;i--{
        if s[i] == '#'{
            num,_ :=  strconv.Atoi(s[i-2:i])
            result = append([]rune{rune(96+num)}, result...  )    
            i= i-2
        }else{
            num,_ :=  strconv.Atoi(string(s[i]))
            result = append([]rune{rune(96+num)},result... )
            
        }
    }
    return  string(result)
}