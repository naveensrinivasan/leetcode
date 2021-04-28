func makeGood(s string) string {
    if len(s) == 1{
        return s
    }
    left,right := 0,1
    for left < right && right < len(s)   {
        if s[left] != s[right] &&  strings.EqualFold(string(s[left]),string(s[right])){
            s= s[:left]+s[right+1:]
            left=0
            right=1
        }else{
            left++
            right++
        }
    }
    return s
}