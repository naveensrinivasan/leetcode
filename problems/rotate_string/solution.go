func rotateString(s string, goal string) bool {
    if s == goal{
        return true
    }
    
    
    for i:=0;i<len(s);i++{
        s = s[1:]+string(s[0])
        if s == goal{
            return true
        }
    }
    
    return false
}