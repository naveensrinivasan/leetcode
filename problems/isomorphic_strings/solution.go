func isIsomorphic(s string, t string) bool {
    if len(s) != len(t){
        return false
    }
    m := make(map[byte]byte)
    m2 := make(map[byte]byte)
    for i:=0;i<len(s);i++{
        if v,ok:=m[s[i]];!ok{
            m[s[i]] = t[i]
        }else{
            if v != t[i]{
                return false
            }
        }
        if v,ok:=m2[t[i]];!ok{
            m2[t[i]] = s[i]
        }else{
            if v != s[i]{
                return false
            }
        }
    }
      
    return true
}