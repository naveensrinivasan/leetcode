func sortedSquares(nums []int) []int {
    
    result := []int{}
    for _,i := range nums{
        result = append(result,i*i)
    }
    sort.Ints(result)
    return result
}