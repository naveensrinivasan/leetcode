func reverse(x int) int {
	pop := 0
	reversed := 0
	for x != 0 {
		pop = x % 10
		// division result
		x /= 10
		reversed = (reversed * 10) + pop
	}
	if math.MaxInt32 < reversed || reversed < math.MinInt32 {
		return 0
	}
	return reversed

}
