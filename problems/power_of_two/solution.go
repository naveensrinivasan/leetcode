func isPowerOfTwo(n int) bool {
	if n == 0 {
		return false
	}
	return twosPower(n)
}
func twosPower(n int) bool {

	if n != 1 {
		if n%2 != 0 {
			return false
		}
		n /= 2
		return twosPower(n)
	}
	return true
}
