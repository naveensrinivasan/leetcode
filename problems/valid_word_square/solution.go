func validWordSquare(words []string) bool {
    for i:=0;i<len(words);i++{
        word := words[i]
        word2 := ""
        fmt.Println(word)
        for j:=0;j<len(word);j++{
            fmt.Println(i,j)
            if j< len(words) &&  i<len(words[j]){
                word2+= string(words[j][i])    
            }
            
        }  
        if word != word2{
            return false
        }
    }
    
    return true
}