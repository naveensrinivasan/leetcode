func isPowerOfFour(num int) bool {
	if num == 0 {
		return false
	}
	if num == 1 {
		return true
	}
	sum := 1
	for sum <= num {
		sum *= 4
		if sum == num {
			return true
		}
	}
	return false
}
