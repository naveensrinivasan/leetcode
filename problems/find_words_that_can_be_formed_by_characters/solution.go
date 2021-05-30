func countCharacters(words []string, chars string) int {
    m := make(map[string]int)
    for i:=0;i<len(chars);i++{
        m[string(chars[i])]++
    }
    
    sum := 0
    
    for _,word := range words{
        if len(m) == 0{
            return sum
        }
        wordMap := make(map[string]int)
        for i:=0;i<len(word);i++{
            wordMap[string(word[i])]++
        }
        fmt.Println(wordMap)
        foundWord := true
        for k,v := range wordMap{
            if value,ok := m[k];!ok{
                fmt.Println(k,v,value)
                foundWord = false
                break
            }else if v > value{
                
                foundWord = false
                break 
            }
        }
        
        if foundWord{
            sum+= len(word)
            // for k := range wordMap{
            //     delete(m,k)
            // }
        }
    }
    return sum  
}