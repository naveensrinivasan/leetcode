func isHappy(n int) bool {
	result := n
	m := make(map[int]int)
	for result != 1 {
		m[result] = result
		result = separateDigits(result)
		if _, ok := m[result]; ok {
			return false
		}

	}
	return true
}
func separateDigits(n int) int {
	result := 0
	for n != 0 {
		x := n % 10
		result += x * x
		n /= 10
	}
	return result
}
