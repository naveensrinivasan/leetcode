func countConsistentStrings(allowed string, words []string) int {
    m:=  make(map[rune]rune)
    result := 0
    for _, i:= range allowed{
        m[i]=i
    }
    for _ , word := range words {
        isallowed := true
        for _, c := range word{
            if _,ok := m[c];!ok{
                isallowed = false
                break
            }
        }
        if isallowed{
            result++
        }
    }
    return result
}