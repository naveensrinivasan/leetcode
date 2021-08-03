func removeOccurrences(s string, part string) string {
    pos := strings.Index(s,part)
    if pos == -1{
        return s
    }
    result := s[:pos] + s[pos+len(part):]
    
    if result != s{
        result = removeOccurrences(result,part)  
    }
    
    return result
}