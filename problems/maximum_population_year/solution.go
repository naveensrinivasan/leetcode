func maximumPopulation(logs [][]int) int {
    type F struct{
        Y,C int
    }
    m := make(map[int]*F)
    l := []*F{}
    for _,item := range logs{
        for i:=item[0];i<item[1];i++{
            if v,ok := m[i];!ok{
                f := &F{Y:i,C:1}
                l = append(l,f)
                m[i] = f
            }else{
                v.C++
            }
        }
    }
    sort.Slice(l,func(i,j int)bool{
        if l[i].C == l[j].C{
            return l[i].Y < l[j].Y
        }
        return l[i].C > l[j].C
    })
    return l[0].Y
}