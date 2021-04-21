func runningSum(nums []int) []int {
    result := []int{}
    sum := 0
    for _,x := range nums{
        sum+=x
        result = append(result,sum)
    }  
    return result
}