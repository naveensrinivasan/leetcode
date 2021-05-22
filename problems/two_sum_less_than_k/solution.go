func twoSumLessThanK(nums []int, k int) int {
    sort.Ints(nums)
    left,right := 0, len(nums)-1
    closest := -1
    for left < right{
        sum := nums[left]+nums[right]
        if sum > k -1{
            right--
            continue
        }
        if sum > closest {
            closest = sum
        }
        left++
    }
    return closest
}