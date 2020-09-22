func missingNumber(nums []int) int {
	sort.Ints(nums)
	if nums[0] != 0 {
		return 0
	}
	if nums[len(nums)-1] != len(nums) {
		return len(nums)
	}
	for i := 1; i < len(nums); i++ {
		if nums[i]-nums[(i-1)] != 1 {
			return nums[i] - 1
		}
	}
	return len(nums) + 1
}
