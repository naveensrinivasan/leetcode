func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i, v := range nums {
		diff := target - v
		if pos, ok := m[diff]; ok {
			return []int{pos, i}
		}
		m[v] = i
	}
	return nil
}
