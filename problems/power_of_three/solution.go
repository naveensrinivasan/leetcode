func isPowerOfThree(n int) bool {
	if n == 0 {
		return false
	}
	if n == 1 {
		return true
	}
	sum := 1
	for sum <= n {
		sum *= 3
		if sum == n {
			return true
		}
	}
	// if sum != n {
	// 	return false
	// }
	return false
}
