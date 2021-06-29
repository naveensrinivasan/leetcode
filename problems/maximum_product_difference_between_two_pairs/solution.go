func maxProductDifference(nums []int) int {
    sort.Ints(nums)
    x := nums[0] * nums[1]
    y := nums[len(nums)-2] * nums[len(nums)-1]
    return y- x
}