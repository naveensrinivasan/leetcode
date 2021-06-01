func frequencySort(s string) string {
    m := make(map[string]int)
    type F struct{
        K string
        V int
    }
    l := []F{}
    for i:=0;i<len(s);i++{
        m[string(s[i])]++
    }
    for k,v := range m{
        l = append(l,F{k,v})
    }
    sort.Slice(l,func(i,j int) bool{
        return l[i].V > l[j].V
    })
    result := ""
    for _,item := range l{
        result += strings.Repeat(item.K,item.V)
    }
    return result
}