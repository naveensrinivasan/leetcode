func arrayRankTransform(arr []int) []int {
    result := []int{}
    x := make([]int,len(arr))
    copy(x,arr)
    sort.Ints(x)
    m := make(map[int]int)
    counter := 1
    for _,j := range x{
        if _,ok:= m[j];!ok{
            m[j]=counter
            counter++
        }
    }
    for _,i := range arr{
        result = append(result,m[i])
    } 
    return result
}