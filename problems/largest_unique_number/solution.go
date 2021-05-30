func largestUniqueNumber(A []int) int {
    l :=[]int{} 
    m := make(map[int]int)
    for _,i:= range A{
        m[i]++
    }
    for k,v := range m{
        if v == 1{
            l = append(l,k)
        }
    }
    sort.Ints(l)
    if len(l) == 0{
        return -1
    }
    return l[len(l)-1]
}