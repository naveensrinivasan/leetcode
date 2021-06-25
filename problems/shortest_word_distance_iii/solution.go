func shortestWordDistance(wordsDict []string, word1 string, word2 string) int {
    w1,w2,min,foundw1,foundw2 := 0,0,3000000,false,false
    same:= false
    m := make(map[int]int)
    if word1==word2{
        same = true
    }
    for i:=0;i<len(wordsDict);i++{
        if wordsDict[i] == word1{
            if same{
                m[i] =i
            }else{
                foundw1=true
                w1 = i
            }
        }
        if wordsDict[i] == word2{
            if same{
                m[i] =i
            }else{
                foundw2 = true
                w2 = i
            }
        }
        x :=abs(w1,w2)
        if foundw1 && foundw2 && x < min{
            min =x
        }
    }
    if same{
        l := []int{}
        for k := range m{
            l = append(l,k)
        }
        sort.Ints(l)
        for i:=0;i<len(l)-1;i++{
            r := abs(l[i],l[i+1])
            if r < min{
                min = r
            }
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