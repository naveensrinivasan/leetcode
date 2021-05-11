func checkInclusion(s1 string, s2 string) bool {
    for i:=0;i<len(s2)-len(s1)+1;i++{
        if issubstring(s1,s2[i:i+len(s1)]){
            return true
        }
    }
    return false
}

func issubstring(s1,s2 string)bool{
    
    m := make(map[byte]int)
    for i:=0;i<len(s1);i++{
        m[s1[i]]++
        m[s2[i]]--
    }
    
    for _,v := range m{
        if v != 0 {
            return false
        }
    }
    return true
}