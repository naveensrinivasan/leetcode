func findNumbers(nums []int) int {
	evenNumbers := 0
	if len(nums) == 0 {
		return evenNumbers
	}
	for _, v := range nums {
		if areevendigits(v) {
			evenNumbers++
		}
	}

	return evenNumbers
}

func areevendigits(n int) bool {
	result := 0
	for n != 0 {
		n /= 10
		result++
	}
	return result%2 == 0
}
