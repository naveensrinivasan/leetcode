func twoSum(nums []int, target int) []int {
    m := make(map[int]int)
	for i, i2 := range nums {
		key := target - i2
		if v ,ok := m[key];ok{
			return []int{v,i}
		}
		m[i2]=i
	}
	return nil
}