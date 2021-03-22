

// @lc code=start
func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	//[2,7,11,15]
	for i, item := range nums {
		t := target - item
		if indice, ok := m[t]; ok {
			return []int{indice, i}
		}
		m[item] = i
	}
	return []int{0, 0}
}