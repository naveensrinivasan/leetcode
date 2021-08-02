func canBeTypedWords(s string, x string) int {
    m := make(map[string]bool)
    
    for _, item := range strings.Split(x,""){
        m[item] = true
    }
    result := 0
    for _ , word := range strings.Fields(s){
        foundBroken := false
        for  item := range m{
            if strings.Contains(word,item){
                foundBroken = true
                break
            }
        }
        if !foundBroken{
            result++
        }
    }
    return result
}