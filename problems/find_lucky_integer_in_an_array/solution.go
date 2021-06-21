func findLucky(arr []int) int {
    m := make(map[int]int)
    result := -1
    for _, i:= range arr{
        m[i]++
    }
    for k,v := range m{
        if k == v {
            if k > result{
                result = k
            }
        }
    }
    return result
    
}