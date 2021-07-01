func countGoodSubstrings(s string) int {
    sum := 0
    for i:=0;i<len(s)-2;i++{
        if isunique(s[i:i+3]){
            sum++
        }
    }
    return sum
}

func isunique(s string)bool{
    m := make(map[byte]int)
    m[s[0]]++
    m[s[1]]++
    m[s[2]]++
    
    for _,v := range m{
        if v!=1{
            return false
        }
    }
    return true
}