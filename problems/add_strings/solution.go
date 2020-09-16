func addStrings(num1 string, num2 string) string {
	numbers := map[byte]int{
		'0': 0,
		'1': 1,
		'2': 2,
		'3': 3,
		'4': 4,
		'5': 5,
		'6': 6,
		'7': 7,
		'8': 8,
		'9': 9,
	}
	x, y := len(num1)-1, len(num2)-1
	carry := 0
	result := ""
	for x >= 0 || y >= 0 {
		sum := 0 + carry
		if x >= 0 {
			sum += numbers[num1[x]]
			x--
		}
		if y >= 0 {
			sum += numbers[num2[y]]
			y--
		}
		if sum > 9 {
			digit := sum % 10
			result += strconv.Itoa(digit)
			carry = sum / 10
			continue
		}

		result += strconv.Itoa(sum)
		carry = 0

	}
	if carry > 0 {
		result += strconv.Itoa(carry)
	}
	endResult := ""
	for _, x := range result {
		endResult = string(x) + endResult
	}
	return endResult
}
