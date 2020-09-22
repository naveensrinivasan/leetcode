func isPalindrome(s string) bool {

	left, right := 0, len(s)-1

	var IsLetter = regexp.MustCompile(`^[a-zA-Z0-9]*$`).MatchString
	for left <= right {
		for left < right && !IsLetter(string(s[left])) {
			left++
		}
		for left < right && !IsLetter(string(s[right])) {
			right--
		}
		if strings.ToLower(string(s[left])) != strings.ToLower(string(s[right])) {
			return false
		}
		left++
		right--
	}
	return true
}
