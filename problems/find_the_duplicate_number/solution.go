func findDuplicate(nums []int) int {
	m := make(map[int]int)
	for _, v := range nums {
		if _, ok := m[v]; ok {
			return v
		}
		m[v]++
	}
	return -1
}
