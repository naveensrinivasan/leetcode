func mostCommonWord(paragraph string, banned []string) string {
    m := make(map[string]int)
    b := make(map[string]string)
    
    for _,word := range banned{
        b[word] = word
    }
    
    for _, word := range strings.FieldsFunc(paragraph,Split){
        word = strings.ToLower(word)
        reg, _ := regexp.Compile("[^a-z]+")
        word := reg.ReplaceAllString(word, "")
        if _,ok:= b[word];!ok{
            m[word]++
        }
    }
    
    
    key,val := "",0
    for k,v := range m{
        if v > val{
            val = v
            key = k
        }
    }
    
    return key
    
}

func Split(r rune) bool {
    return r == ' ' || r == ','
}