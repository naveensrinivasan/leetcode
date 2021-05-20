func uncommonFromSentences(s1 string, s2 string) []string {
    m := make(map[string]int)
    result := []string{}
    for _, word := range strings.Split(s1, " "){
        m[word]++
    }
    for _, word := range strings.Split(s2, " "){
        m[word]++
    }
    for k,v := range m{
        if v == 1{
            result = append(result,k)
        }
    }
    return result
}