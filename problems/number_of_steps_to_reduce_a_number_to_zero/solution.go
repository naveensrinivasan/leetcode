func numberOfSteps(num int) int {
	if num <= 1 {
		return num
	}
	result := 0
	for num > 0 {
		if num&1 == 1 {
			num -= 1
		} else {
			num >>= 1
		}
		result++
	}
	return result

}
