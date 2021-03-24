func lengthOfLongestSubstring(s string) int {
   a,b,max,mapsize := 0,0,0,0
    m := make(map[byte]byte)
    for (b<len(s)){
        if _,ok := m[s[b]];!ok{
            
            m[s[b]]=s[b]
            b++
            mapsize++
            max = Max(max,mapsize)
        }else{
            delete(m,s[a])
            a++
            mapsize--
        }
    }
    
    return max
}

// Max returns the larger of x or y.
func Max(x, y int) int {
    if x < y {
        return y
    }
    return x
}