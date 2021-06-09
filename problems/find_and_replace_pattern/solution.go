func findAndReplacePattern(words []string, pattern string) []string {
    result := []string{}
    for _, word := range words{
        m := make(map[string]string)
        w := make(map[string]string)
        
        for i:=0;i<26;i++{
            m[string(rune(i+97))]= string(rune(i+97))
        }
        
        for i:=0;i<len(word);i++{
            if _,ok := w[string(word[i])];!ok{
                if _,exists := m[string(pattern[i])];exists{
                    w[string(word[i])] = m[string(pattern[i])]
                    delete(m,string(pattern[i]))
                }else{
                    break
                }
            }
        }
        
        r := ""
        for i:=0;i<len(word);i++{
            if v,ok := w[string(word[i])];ok{
                r += v
            }else{
                break
            }
        }
        if r == pattern{
            result = append(result,word)
        }
    }
    return result
}