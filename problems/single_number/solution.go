func singleNumber(nums []int) int {
	m := make(map[int]bool)
	sum := 0
	hsum := 0
	for _, item := range nums {
		if _, exists := m[item]; !exists {
			m[item] = true
		}
		sum += item
	}
	for k, _ := range m {
		hsum += k
	}
	return (2*(hsum) - sum)

}
