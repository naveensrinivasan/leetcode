func backspaceCompare(s string, t string) bool {
    left,right := 0,1
    
    for left <= right && right <= len(s)   {
        if s[left] == '#'  && left != 0{
            s = s[:left-1]+s[right:]
            left=0
            right=1
        }else{
            left++
            right++
        }
    }
        
    left,right = 0,1
    for left <= right && right <= len(t)   {
        if t[left] == '#' && left != 0 {
            t = t[:left-1]+t[right:]
            left=0
            right=1
        }else{
            left++
            right++
        }
    }
    
    return strings.Trim(s,"#")==strings.Trim(t,"#")
}