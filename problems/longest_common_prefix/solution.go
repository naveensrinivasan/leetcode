func longestCommonPrefix(strs []string) string {
    result := ""
    
    sort.Slice(strs,func(i,j int)bool{
        return len(strs[i]) < len(strs[j])
    })
    
    for j:=0;j<len(strs[0]);j++{   
        for i:= 1;i<len(strs);i++{
            if strs[i][j] != strs[0][j]{
             return result
            }
         } 
        result+= string(strs[0][j])
    }
    return result
}