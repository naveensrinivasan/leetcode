func arrayRankTransform(arr []int) []int {
  //40 1   10 1 40 4
  //10 2   20 2  
  //20 3   30 3
  //30 4   40 4
    result := []int{}
    x := make([]int,len(arr))
    copy(x,arr)
    sort.Ints(x)
    fmt.Println(x)
    m := make(map[int]int)
    counter := 1
    for _,j := range x{
        if _,ok:= m[j];!ok{
            m[j]=counter
            counter++
        }
    }
    fmt.Println(m)
    for _,i := range arr{
        result = append(result,m[i])
    } 
    return result
    
}
func getPos(x []int,item int)int{
    
    if x[0] == item{
        return 1
    }
    for i:=1;i<len(x);i++{
        if x[i] == item {
            if x[i] == x[i-1]{
                return i
            }
            return i+1
        }
    }
    
    return -1
}