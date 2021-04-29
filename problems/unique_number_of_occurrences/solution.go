func uniqueOccurrences(arr []int) bool {
    data := make([]int,2000)
    m := make(map[int]int)
    for _, i := range arr{
        data[i+1000]++
    }
    for _, i:= range data{
        if i==0 {
            continue
        }
        if _, ok:= m[i];!ok{
            m[i]=i
        }else{
            return false
        }
    }
    
    return true
}