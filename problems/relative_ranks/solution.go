func findRelativeRanks(score []int) []string {
    type F struct{
        S,K int
    }
    l := make([]F,len(score))
    for i,item := range score{
        l[i] = F{item,i}
    }
    sort.Slice(l,func(i,j int)bool{
        return l[i].S > l[j].S
    })
    result := make([]string,len(score))
    for i:=0;i<len(l);i++{
        if i == 0 {
            result[l[i].K] = "Gold Medal"
        }else if i == 1{
            result[l[i].K] = "Silver Medal"
        }else if i == 2{
            result[l[i].K] = "Bronze Medal"
        }else{
            result[l[i].K] = strconv.Itoa(i+1)
        }
    }
    return result
}