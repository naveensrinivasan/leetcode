func stringMatching(words []string) []string {
    result := []string{}
    
    sort.Slice(words,func(i,j int)bool{
        return len(words[i]) < len(words[j])
    })
    
    for i:=0;i<len(words);i++{
        for j:= i+1;j<len(words);j++{
            if strings.Contains(words[j],words[i]) {
                    result = append(result,words[i])
                    break
            }
        }
    }
    return result
}