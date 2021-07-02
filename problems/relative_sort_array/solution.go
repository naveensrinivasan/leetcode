func relativeSortArray(arr1 []int, arr2 []int) []int {
    type F struct{
        K,V int
    }
    result := make([]int,len(arr1))
    l := make([]F,len(arr1))
    m := make(map[int]int)
    sort.Ints(arr1)
    counter := len(arr1)+1
    
    for i:=0;i<len(arr2);i++{
        m[arr2[i]]=i
    }
    
    for j, i := range arr1{
        if v,ok := m[i];ok{
            l[j] = F{v,i}
        }else{
            l[j] = F{counter,i}
            counter++
        }
    }
    
    sort.Slice(l,func(i,j int)bool{
        return l[i].K < l[j].K
    })
    
    for i:=0;i<len(l);i++{
        result[i]=l[i].V
    }
    
    return result
}