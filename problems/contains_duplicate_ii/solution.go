func containsNearbyDuplicate(nums []int, k int) bool {
	h := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		if _, ok := h[nums[i]]; ok {
			return true
		}
		h[nums[i]] = 1
		if len(h) > k {
			delete(h, nums[int(math.Abs(float64(k-i)))])
		}

	}
	return false
}
