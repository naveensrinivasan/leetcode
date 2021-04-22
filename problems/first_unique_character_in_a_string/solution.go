func firstUniqChar(s string) int {
    m := make([]int,26)
    
    for _,j:= range s{
      m[j-'a']++
    }
    
    for i,j := range s{
        if m[j-'a'] ==1{
            return i
        }
    }
    return -1
}