func topKFrequent(words []string, k int) []string {
    m := make(map[string]int)
    type F struct{
        K string
        V int
    }
    for i:=0;i<len(words);i++{
        m[words[i]]++
    }
    l := []F{}
    for key,v := range m{
        l = append(l,F{key,v})
    }
    sort.Slice(l,func(i,j int)bool{
        if l[i].V == l[j].V{
            return l[i].K < l[j].K
        }
        return l[i].V > l[j].V
    })
    result := []string{}
    for i:=0;i<k;i++{
        result = append(result,l[i].K)
    }
    return result
}