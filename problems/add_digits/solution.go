func addDigits(num int) int {
	x := num
	sum := 0
	for x != 0 {
		sum = x%10 + sum
		x /= 10
	}
	if sum > 9 {
		sum = addDigits(sum)
	}

	return sum
}
