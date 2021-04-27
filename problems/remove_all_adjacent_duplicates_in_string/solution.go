func removeDuplicates(S string) string {
    fmt.Println(S)
    s := []rune{}
    for i,j := range S{
        if i ==0{
            s = append(s,rune(j))
            continue
        }
        if len(s) > 0 && rune(j) == s[len(s)-1] {
            s = s[:len(s)-1]
        }else{
            s = append(s,rune(j))
        } 
    }
    return string(s)
}