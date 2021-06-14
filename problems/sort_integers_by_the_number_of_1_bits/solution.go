func sortByBits(arr []int) []int {
    type F struct{
        K, V int
    }
    l := []F{}
    for i:=0;i<len(arr);i++{
        sum:= 0
        x := arr[i]
        for x > 0{
            if x & 1  == 1{
                sum++
            } 
            x >>= 1
        }
        l = append(l, F{arr[i],sum})
    }
    
    sort.Slice(l,func(i,j int)bool{
        if l[i].V == l[j].V{
            return l[i].K < l[j].K
        }
            return l[i].V < l[j].V
    })
    result := []int{}
    for _,i := range l{
        result = append(result, i.K )
    }
    return result
}