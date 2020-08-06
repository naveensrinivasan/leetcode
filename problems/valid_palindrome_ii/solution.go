func validPalindrome(s string) bool {
	result, left, right := isPaalindrome(s, 0, len(s)-1)
	if result == true {
		return true
	}
	result, _, _ = isPaalindrome(s, left+1, right)
	if result {
		return true
	}
	result, _, _ = isPaalindrome(s, left, right-1)
	if result {
		return true
	}

	return false
}
func isPaalindrome(s string, left, right int) (bool, int, int) {
	for left < right {
		if s[left] != s[right] {
			return false, left, right
		}
		left++
		right--
	}
	return true, 0, 0
}
