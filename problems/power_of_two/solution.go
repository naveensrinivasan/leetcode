func isPowerOfTwo(n int) bool {
	result := 0
	if n == 0 {
		return false
	}
	for n != 0 {
		result += (n & 1)

		if result > 1 {
			return false
		}
		fmt.Println(result, n)
		n = n >> 1

	}
	fmt.Println(result)
	return true
}
