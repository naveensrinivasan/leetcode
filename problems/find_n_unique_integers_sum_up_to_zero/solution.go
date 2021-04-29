func sumZero(n int) []int {
    if n == 1{
        return []int{0}
    }
    
    r := n/2
    start := r * -1
    
    result := []int{}
    for i:= start;i<=r;i++{
        if n%2== 0 && i == 0{
            continue
        }
        result = append(result,i)
    }
    return result
}