func restoreString(s string, indices []int) string {
    result := make([]byte,len(s))
    
    for i,j := range indices{
        result[j] = s[i]
    }
    
    return string(result)
}