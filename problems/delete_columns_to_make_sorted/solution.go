func minDeletionSize(strs []string) int {
    result := 0
    for i:=0;i<len(strs[0]);i++{
        l := []string{} 
        x := ""
        for j:=0;j<len(strs);j++{
            l = append(l,string(strs[j][i]))
            x += string(strs[j][i])
        }
        sort.Strings(l)
        if  strings.Join(l,"") != x{
            result++
        }
    }
    return result
}