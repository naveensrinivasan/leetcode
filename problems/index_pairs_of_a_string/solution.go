func indexPairs(text string, words []string) [][]int {
    result := [][]int{}
    for _, word := range words{
        x:= strings.Index(text, word)
        if x < 0{
            continue
        }
        result = append(result,[]int{x,(x+len(word)-1)})
        for {
            x = indexAt(text,word,x+1)
            if x == -1{
                break
            }
            result = append(result,[]int{x,(x+len(word)-1)})
        }
        
    }
    sort.Slice(result,func(i,j int)bool{
        if result[i][0] == result[j][0]{
          return  result[i][1]< result[j][1]
        }
        return result[i][0] < result[j][0]
    })
    return result
}
func indexAt(s, sep string, n int) int {
    idx := strings.Index(s[n:], sep)
    if idx > -1 {
        idx += n
    }
    return idx
}