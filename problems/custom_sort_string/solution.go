type Custom struct{
    S []rune
    m map[string]int
}

func (s Custom) Len() int {
    return len(s.S)
}

func (s Custom) Swap(i, j int) {
    s.S[i], s.S[j] = s.S[j], s.S[i]
}

func (s Custom) Less(i, j int) bool {
    return s.m[string(s.S[i])] <  s.m[string(s.S[j])]
}


func customSortString(order string, str string) string {
    m := make(map[string]int)
    
    for i,j := range order{
        m[string(j)] =i
    }
    
    for i:=97;i<123;i++{
        if _,ok := m[string(i)];!ok{
            m[string(i)] = i
        }
    }
    s := []rune(str)
    c:= Custom{S:s,m:m}
    sort.Sort(c)
    return string(c.S)
}