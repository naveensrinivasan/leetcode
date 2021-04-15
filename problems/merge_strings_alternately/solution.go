func mergeAlternately(word1 string, word2 string) string {
    word1bigger,word2bigger,loop := false,false , len(word1)
    if len(word1) > len(word2){
        word1bigger = true
        loop = len(word2)
    }else if len(word1) < len(word2){
        word2bigger = true
        
    }
    result := []byte{}
    
    for i:=0;i<loop;i++{
        result = append(result,word1[i])
        result = append(result,word2[i])
    }
    s := string(result)
    if word1bigger{
        s += word1[loop:]
    }else if word2bigger{
        s+= word2[loop:]
    }
    return s
}