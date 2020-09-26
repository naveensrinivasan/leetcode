func isArmstrong(N int) bool {
	d := splitDigits(N)
	sum := 0
	for _, v := range d {
		sum += int(math.Pow(float64(v), float64(len(d))))
	}
	return sum == N
}

func splitDigits(N int) []int {
	result := []int{}
	for N > 0 {
		x := N % 10
		result = append(result, x)
		N /= 10
	}
	return result
}