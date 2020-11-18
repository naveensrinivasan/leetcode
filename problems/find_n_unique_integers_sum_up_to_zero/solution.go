func sumZero(n int) []int {
	result := []int{}
	if n%2 == 1 {
		result = append(result, 0)
	}
	for i := 1; i <= n/2; i++ {
		result = append(result, i, -i)
	}
	return result
}
