type WordDistance struct {
    wordsDict []string
}


func Constructor(wordsDict []string) WordDistance {
    return WordDistance{wordsDict}
}


func (this *WordDistance) Shortest(word1 string, word2 string) int {
    w1,w2,min,foundw1,foundw2 := 0,0,3000000,false,false
    for i:=0;i<len(this.wordsDict);i++{
        if this.wordsDict[i] == word1{
            foundw1=true
            w1 = i
        }
        if this.wordsDict[i] == word2{
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

/**
 * Your WordDistance object will be instantiated and called as such:
 * obj := Constructor(wordsDict);
 * param_1 := obj.Shortest(word1,word2);
 */