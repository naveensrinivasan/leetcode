func stringMatching(words []string) []string {
    sort.Slice(words,func(i,j int)bool{
        return len(words[i]) < len(words[j])
    })
    m := make(map[string]string)
    result := []string{}
    for i:=0;i<len(words);i++{
        for j:= i+1;j<len(words);j++{
            if _,ok := m[words[i]];!ok && len(words[i]) <= len(words[j]) && strings.Contains(words[j],words[i]) {
                    m[words[i]] = words[i]
            }
        }
    }
    for k,_ := range m{
        result = append(result,k)
    }
    return result
}