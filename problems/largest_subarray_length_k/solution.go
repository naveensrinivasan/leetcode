func largestSubarray(nums []int, k int) []int {
    result := []int{}
    max := 0
    for i:=0;i<=len(nums)-k;i++{
        if nums[i] > max{
            max = nums[i]
            result = nums[i:i+k]
        }
    } 
    return result
}