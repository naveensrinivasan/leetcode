func shortestDistance(wordsDict []string, word1 string, word2 string) int {
    w1,w2,min,foundw1,foundw2 := 0,0,3000000,false,false
    for i:=0;i<len(wordsDict);i++{
        if wordsDict[i] == word1{
            foundw1=true
            w1 = i
        }
        if wordsDict[i] == word2{
            foundw2 = true
            w2 = i
        }
        x :=abs(w1,w2)
        if foundw1 && foundw2 && x < min{
            min =x
        }
    }
    return min
}

func abs(i,j int)int{
    x := i-j
    if x< 0{
        return x * -1
    }
    return x
}