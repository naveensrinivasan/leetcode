func minimumAbsDifference(arr []int) [][]int {
    type F struct{
        K1,K2,D int
    }
    sort.Ints(arr)
    l := []F{}
    for i:=0;i<len(arr)-1;i++{
        l = append(l,F{arr[i],arr[i+1],abs(arr[i],arr[i+1])})
    }
    
    sort.Slice(l,func(i,j int)bool{
        if l[i].D == l[j].D{
            return l[i].K1 < l[j].K1 && l[i].K2 < l[j].K2
        }
        
        return l[i].D < l[j].D
    })
    
    d := l[0].D
    result :=[][]int{}
    for i:=0;i<len(l);i++{
        if d != l[i].D{
            return result
        }
        result = append(result,[]int{l[i].K1,l[i].K2})
    }
    return result
}

func abs(x,y int)int{
    z := x-y
    if z <0 {
        return z * -1
    }
    return z
}