func findWords(words []string) []string {
    m:= make(map[string]int)
    f := "qwertyuiop"
    s := "asdfghjkl"
    t := "zxcvbnm"
    
    for i:=0;i<len(f);i++{
        m[string(f[i])]=1
    }
    for i:=0;i<len(s);i++{
        m[string(s[i])]=2
    }
    for i:=0;i<len(t);i++{
        m[string(t[i])]=3
    }
    
    result := []string{}
    
    for _, word := range words{
        w := strings.ToLower(word)
        row :=  m[string(w[0])]
        diffRow := false
        for i:=1;i<len(w);i++{
            if m[string(w[i])] != row{
                diffRow = true
                break
            }
        }
        if !diffRow{
            result = append(result,word) 
        }
    }
    return result
}