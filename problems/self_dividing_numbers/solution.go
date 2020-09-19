func selfDividingNumbers(left int, right int) []int {
	result := []int{}
	for i := left; i <= right; i++ {
		wasDivisble := true
		x := separateDigit(i)
		if len(x) == 0 {
			continue
		}
		for _, d := range x {
			if i%d != 0 {
				wasDivisble = false
				break
			}
		}
		if wasDivisble {
			result = append(result, i)
		}

	}
	return result
}
func separateDigit(n int) []int {
	result := []int{}
	if n == 1 {
		return []int{1}
	}
	for n > 0 {
		if n%10 == 0 {
			return []int{}
		}
		result = append(result, n%10)
		n /= 10
	}
	return result
}
