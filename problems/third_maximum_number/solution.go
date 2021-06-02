func thirdMax(nums []int) int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] > nums[j]
	})
	if len(nums) < 3 {
		return nums[0]
	}

	y := nums[:1]

	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[i-1] {
			y = append(y, nums[i])
		}
	}

	if len(y) > 2 {
		return y[2]
	} else {
		return y[0]
	}
}
