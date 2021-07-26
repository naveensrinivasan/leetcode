func areOccurrencesEqual(s string) bool {
    m := make(map[rune]int)  
    for _, i := range s{
        m[i]++
    }
    x:= m[rune(s[0])]
    for _,v := range m{
        if v != x{
            return false
        }
    }
    return true
}