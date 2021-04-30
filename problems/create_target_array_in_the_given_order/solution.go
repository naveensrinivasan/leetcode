func createTargetArray(nums []int, index []int) []int {
    a := []int{}
    for i := range nums{
        a = append(a[:index[i]],append([]int{nums[i]}, a[index[i]:]...)...)
    }
    return a
}