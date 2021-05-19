func canPermutePalindrome(s string) bool {
    if len(s) ==1{
        return true
    }
    m := make(map[string]int)
    odds,evens := 0,0
    
    for _, a := range s {
        m[string(a)]++
    }
    if len(m) == 1{
        return true
    }
    for _,v := range m{
        if v%2 == 0{
            evens++
        }else{
            odds++
        } 
    }
    fmt.Println(m,odds,evens)
    if evens == odds{
        return true
    }
    if odds == 0 {
        return true
    }
    if evens > odds {
        if odds == 1{
            return true
        }
        return (odds +evens)%2 ==0
    }
    
    return false
}