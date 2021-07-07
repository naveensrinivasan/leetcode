func nextGreaterElements(nums []int) []int {
    type F struct{
        K,P,V int
    }
    ms:= []*F{}
    l := []*F{}
    x := append(nums,nums...)
    for i:=0;i<len(x);i++{
        for len(ms) > 0 && x[i]> ms[len(ms)-1].K{
            ms[len(ms)-1].V = x[i]
            ms = ms[:len(ms)-1]
        }
        v := &F{x[i],i,-1}
        ms = append(ms,v)
        l = append(l,v)
    }
    sort.Slice(l,func(i,j int)bool{
        return l[i].P < l[j].P
    })
    r := make([]int,len(nums))
    for i , x := range l{
        if i < len(nums){
         r[i] = x.V
        }else{
            break
        }
        i++
    }
    return r
}