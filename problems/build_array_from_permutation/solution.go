func buildArray(nums []int) []int {
    result := []int{}
    for i :=0;i<len(nums);i++{
        result = append(result,nums[nums[i]])
    }
    return result
}