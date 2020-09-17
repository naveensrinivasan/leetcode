func rotate(nums []int, k int) {
	for i := 0; i < k; i++ {
		prev, next, last := nums[0], nums[0], nums[len(nums)-1]
		for j := 0; j < len(nums)-1; j++ {
			prev = next
			next = nums[j+1]
			nums[j+1] = prev
		}
		nums[0] = last
	}
}
