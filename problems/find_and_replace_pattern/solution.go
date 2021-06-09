func findAndReplacePattern(words []string, pattern string) []string {
    result := []string{}
    for _, word := range words{
        m := make(map[string]string)
        w := make(map[string]string)
        
        for i:=0;i<26;i++{
            m[string(rune(i+97))]= string(rune(i+97))
        }
        
        r := ""
        for i:=0;i<len(word);i++{
            if _,ok := w[string(word[i])];!ok{
                if _,exists := m[string(pattern[i])];exists{
                    w[string(word[i])] = m[string(pattern[i])]
                    r += string(pattern[i])
                    delete(m,string(pattern[i]))
                }else{
                    break
                }
            }else{
                r+= w[string(word[i])]
            }
        }
        /*
        for i:=0;i<len(word);i++{
            if v,ok := w[string(word[i])];ok{
                r += v
            }else{
                break
            }
        }*/
        if r == pattern{
            result = append(result,word)
        }
    }
    return result
}