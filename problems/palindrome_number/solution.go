func isPalindrome(x int) bool {
    if x < 0  {
		return false
	}
	if x < 10 {
		return true
	}
	y := x
	pop := 0
	reversed := 0
	for x != 0 {
		pop = x % 10
		// division result
		x /= 10
		reversed = (reversed * 10) + pop
	}
	return reversed == y
}