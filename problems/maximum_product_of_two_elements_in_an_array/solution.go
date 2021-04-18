func maxProduct(nums []int) int {
    if len(nums) ==2{
        return (nums[0]-1) * (nums[1] -1)
    }
    sort.Ints(nums)
    return (((nums[len(nums)-1])-1) * ((nums[len(nums)-2])-1))
}