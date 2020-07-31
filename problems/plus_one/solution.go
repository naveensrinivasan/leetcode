func plusOne(digits []int) []int {
	carry := 0
	first := false
	for i := len(digits) - 1; i >= 0; i-- {
		sum := 0
		if !first {
			sum = digits[i] + 1
			first = true
		} else {
			sum += digits[i] + carry
		}
		if sum > 9 {
			carry = 1
			digits[i] = sum % 10
		} else {
			carry = 0
			digits[i] = sum
		}
	}
	if carry > 0 {
		digits = append([]int{1}, digits...)
	}
	return digits
}
